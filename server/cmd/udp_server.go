package main

import (
	"log"
	"net"
	"strings"

	"github.com/google/uuid"
)

type UdpServer struct {
	conns map[string]string
	conn  *net.UDPConn
}

func NewUdpServer() (*UdpServer, error) {

	conns := make(map[string]string)

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

		if rlen < 16 || err != nil {
			continue
		}

		token, err := uuid.FromBytes(message[:16])

		if err != nil {
			continue
		}

		data := strings.TrimSpace(string(message[16:rlen]))

		log.Printf("received \"%s\" from %s using token %s", data, remote, token)

		broadcast(s.conns, data)
	}
}
