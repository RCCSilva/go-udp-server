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
	Port int `json:"port"`
}

type authResponse struct {
	Token string `json:"token"`
}

func NewAuth() *Auth {
	auth := &Auth{}

	return auth
}

const serverHost = "http://localhost:5000"
const jsonContentType = "application/json"

func (a *Auth) Authenticate(port int) (*uuid.UUID, error) {
	jsonData, err := json.Marshal(authRequest{Port: port})
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(
		fmt.Sprintf("%s/authenticate", serverHost),
		jsonContentType,
		bytes.NewBuffer(jsonData),
	)

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
