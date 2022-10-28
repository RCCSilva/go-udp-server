package auth

import (
	"log"

	"github.com/google/uuid"
)

func Authenticate() (uuid.UUID, error) {
	id := uuid.New()
	log.Printf("created token: %s", id)

	return id, nil
}
