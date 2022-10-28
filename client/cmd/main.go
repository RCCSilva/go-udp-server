package main

import (
	"log"
	"os"
)

func main() {
	listener := NewListener()
	err := listener.Setup()

	if err != nil {
		log.Fatalf("failed to setup listener: %v", err)
	}

	go listener.Listen()

	auth := NewAuth()
	token, err := auth.Authenticate(listener.udpAddr.Port)

	if err != nil {
		log.Fatalf("failed to authenticate: %v", err)
	}

	//

	client := NewClient(listener, token)

	if len(os.Args) > 1 && os.Args[1] == "infinite" {
		client.InfiniteLoop()
	}

	client.LoopInput()

}
