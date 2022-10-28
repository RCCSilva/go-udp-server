package main

import (
	"log"
	"net"
	"strings"

	"github.com/google/uuid"
)

const tokenLength = 16

type UdpServer struct {
	conn  *net.UDPConn
	conns *Conns
}

func NewUdpServer(conns *Conns) (*UdpServer, error) {

	conn, err := net.ListenUDP("udp", &net.UDPAddr{
		Port: 3000,
		IP:   net.ParseIP("0.0.0.0"),
	})

	if err != nil {
		return nil, err
	}

	return &UdpServer{conn: conn, conns: conns}, nil

}

func (s *UdpServer) Listen() {
	defer s.conn.Close()

	log.Printf("udp server listening %s", s.conn.LocalAddr().String())
	message := make([]byte, 508)

	for {
		rlen, remote, err := s.conn.ReadFromUDP(message[:])

		if rlen < tokenLength || err != nil {
			continue
		}

		token, err := uuid.FromBytes(message[:tokenLength])

		if err != nil {
			continue
		}

		isValid := s.conns.verifyToken(token)

		if !isValid {
			log.Printf("%s token is invalid\n", token)
			continue
		}

		data := strings.TrimSpace(string(message[tokenLength:rlen]))

		log.Printf("received \"%s\" from %s using token %s", data, remote, token)

		s.conns.broadcast(data)
	}
}
