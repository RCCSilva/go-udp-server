package main

import (
	"encoding/json"
	"log"
	"net"
	"net/http"
	"rccsilva/go-udp-server/server/auth"
)

type HttpServer struct {
	conns *Conns
	http.Handler
}

type authRequest struct {
	Port int `json:"port"`
}

type authResponse struct {
	Token string `json:"token"`
}

func NewHttpServer(conns *Conns) *HttpServer {
	router := http.NewServeMux()

	server := &HttpServer{
		Handler: router,
		conns:   conns,
	}

	router.Handle("/authenticate", http.HandlerFunc(server.authHandler))

	return server
}

func (h *HttpServer) authHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("received request to authenticate")
	token, _ := auth.Authenticate()

	ip, _ := net.ResolveTCPAddr("tcp", r.RemoteAddr)
	var request authRequest

	json.NewDecoder(r.Body).Decode(&request)

	h.conns.addConnection(token, ip.IP.String(), request.Port)

	json.NewEncoder(w).Encode(authResponse{Token: token.String()})
}
