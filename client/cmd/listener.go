package main

import (
	"log"
	"net"
	"strings"

	"github.com/google/uuid"
)

type Listener struct {
	udpAddr *net.UDPAddr
	conn    *net.UDPConn
	token   *uuid.UUID
}

func NewListener() *Listener {
	listener := &Listener{}

	return listener
}

func (l *Listener) Setup() error {
	conn, err := net.ListenUDP("udp", &net.UDPAddr{
		IP: net.ParseIP("0.0.0.0"),
	})

	addr, _ := net.ResolveUDPAddr("udp", conn.LocalAddr().String())

	l.udpAddr = addr
	l.conn = conn

	return err
}

func (l *Listener) Listen() {
	defer l.conn.Close()
	log.Printf("server listening %s", l.conn.LocalAddr().String())

	message := make([]byte, 20)

	for {
		rlen, remote, err := l.conn.ReadFromUDP(message[:])

		if err != nil {
			log.Printf("failed to receive message: %v", err)
			continue
		}

		data := strings.TrimSpace(string(message[:rlen]))
		log.Printf("received: \"%s\" from %s\n", data, remote)
	}
}
