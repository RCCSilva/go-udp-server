package main

import (
	"fmt"
	"log"
	"net"
)

func broadcast(conns map[string]string, message string) {
	for _, addr := range conns {
		go send(addr, message)
	}

}

func send(addr, message string) {
	conn, err := net.Dial("udp", addr)
	if err != nil {
		log.Panicf("failed to create connection: %v", err)
	}
	defer conn.Close()

	_, err = fmt.Fprint(conn, message)

	if err != nil {
		log.Printf("failed to write message: %v", err)
	}
}
