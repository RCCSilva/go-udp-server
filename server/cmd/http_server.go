package main

import (
	"encoding/json"
	"log"
	"net/http"
	"rccsilva/go-udp-server/server/auth"
)

type HttpServer struct {
	http.Handler
}

type authResponse struct {
	Token string `json:"token"`
}

func NewHttpServer() *HttpServer {
	router := http.NewServeMux()

	router.Handle("/authenticate", http.HandlerFunc(authHandler))

	server := &HttpServer{
		Handler: router,
	}

	return server
}

func authHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("received request to authenticate")
	id, _ := auth.Authenticate()
	json.NewEncoder(w).Encode(authResponse{Token: id.String()})
}
