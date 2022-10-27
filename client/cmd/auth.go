package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"
)

type Auth struct{}

type authRequest struct {
	Port string `json:"port"`
}

type authResponse struct {
	Token string `json:"token"`
}

func NewAuth() *Auth {
	auth := &Auth{}

	return auth
}

func (a *Auth) Authenticate(port string) (*uuid.UUID, error) {
	jsonData, err := json.Marshal(authRequest{Port: port})
	if err != nil {
		return nil, err
	}

	resp, err := http.Post("http://localhost:5000/authenticate", "application/json", bytes.NewBuffer(jsonData))

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to reach server with status 200: %v", resp.Status)
	}

	defer resp.Body.Close()
	var response authResponse

	err = json.NewDecoder(resp.Body).Decode(&response)

	if err != nil {
		return nil, err
	}

	log.Printf("authenticated with server using token \"%s\"", response.Token)

	token, err := uuid.Parse(response.Token)

	if err != nil {
		return nil, err
	}

	return &token, nil
}
