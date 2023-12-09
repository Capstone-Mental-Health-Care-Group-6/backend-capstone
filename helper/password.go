package helper

import (
	"golang.org/x/crypto/bcrypt"
)

type Password interface {
	HashPassword(string) (string, error)
}

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}
