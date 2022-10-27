package auth

import (
	"github.com/google/uuid"
)

func Authenticate() (uuid.UUID, error) {
	id := uuid.New()

	return id, nil
}
