package main

import (
	"log"
	"net/http"
)

func main() {
	conns := NewConns()
	udpServer, err := NewUdpServer(conns)
	if err != nil {
		log.Panicf("failed to start server: %v", err)
	}

	go udpServer.Listen()

	httpServer := NewHttpServer(conns)

	err = http.ListenAndServe(":5000", httpServer)

	log.Panicf("failed to create HTTP server: %v", err)
}
