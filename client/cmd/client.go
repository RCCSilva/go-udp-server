package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"net"
	"os"
	"time"

	"github.com/google/uuid"
)

type Client struct {
	conn   net.Conn
	server *Listener
	token  *uuid.UUID
}

func NewClient(server *Listener, token *uuid.UUID) *Client {
	conn, err := net.Dial("udp", "0.0.0.0:3000")
	if err != nil {
		log.Panicf("failed to send connection: %v", err)
	}

	return &Client{conn: conn, server: server, token: token}
}

func (s *Client) sendString(msg string) error {
	writer := bufio.NewWriterSize(s.conn, 408)

	tokenBin, err := s.token.MarshalBinary()

	if err != nil {
		return err
	}

	_, err = writer.Write(tokenBin)

	if err != nil {
		return err
	}

	_, err = writer.WriteString(msg)

	if err != nil {
		return err
	}

	return writer.Flush()
}

func (s *Client) InfiniteLoop() {
	defer s.conn.Close()
	for {
		err := s.sendString(fmt.Sprint(rand.Int()))

		if err != nil {
			log.Printf("failed to send message")
		}

		time.Sleep(30 * time.Millisecond)
	}
}

func (s *Client) LoopInput() {
	defer s.conn.Close()
	reader := bufio.NewReader(os.Stdin)

	for {
		text, _ := reader.ReadString('\n')

		err := s.sendString(text)

		if err != nil {
			log.Printf("failed to write message: %v", err)
		}
	}
}
