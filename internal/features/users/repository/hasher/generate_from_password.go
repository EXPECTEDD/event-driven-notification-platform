package users_hasher_repository

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func (h *UsersHasherRepository) GenerateFromPassword(
	password []byte,
) ([]byte, error) {
	passwordHash, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("generate password hash: %w", err)
	}
	return passwordHash, err
}
