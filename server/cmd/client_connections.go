package main

import (
	"fmt"
	"log"
	"net"

	"github.com/google/uuid"
)

type ClientConnection struct {
	ip   string
	port int
}

func (c *ClientConnection) toAddr() string {
	return fmt.Sprintf("%s:%d", c.ip, c.port)
}

type Conns struct {
	connections map[string]ClientConnection
}

func NewConns() *Conns {
	return &Conns{connections: make(map[string]ClientConnection)}
}

func (c *Conns) addConnection(token uuid.UUID, ip string, port int) {
	log.Printf("adding connection (token: %s; ip: %s; port: %d)", token, ip, port)
	c.connections[token.String()] = ClientConnection{ip: ip, port: port}
}

func (c *Conns) verifyToken(token uuid.UUID) bool {
	_, has := c.connections[token.String()]

	return has
}

func (c *Conns) broadcast(message string) {
	for _, addr := range c.connections {
		go c.send(addr, message)
	}

}

func (*Conns) send(addr ClientConnection, message string) {
	conn, err := net.Dial("udp", addr.toAddr())
	if err != nil {
		log.Panicf("failed to create connection: %v", err)
	}
	defer conn.Close()

	_, err = fmt.Fprint(conn, message)

	if err != nil {
		log.Printf("failed to write message: %v", err)
	}
}
