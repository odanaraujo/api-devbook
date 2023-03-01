package security

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
)

func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func VerifyPassword(passwordHash string, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password))

	if err != nil {
		return errors.New("incorrect password")
	}

	return nil
}
