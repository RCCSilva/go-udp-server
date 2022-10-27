package main

import (
	"log"
	"net"
	"regexp"
	"strings"
)

type Listener struct {
	port string
	conn *net.UDPConn
}

func NewListener() *Listener {
	listener := &Listener{}

	return listener
}

func (l *Listener) Setup() error {
	addr := &net.UDPAddr{
		IP: net.ParseIP("0.0.0.0"),
	}
	conn, err := net.ListenUDP("udp", addr)

	rex := regexp.MustCompile(":(?P<port>\\d+)")
	finds := rex.FindStringSubmatch(conn.LocalAddr().String())

	l.port = finds[1]
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
