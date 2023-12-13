package enkrip

import (
	"golang.org/x/crypto/bcrypt"
)

type HashInterface interface {
	Compare(hashed string, input string) error
	HashPassword(string) (string, error)
}

type hash struct{}

func New() HashInterface {
	return &hash{}
}

func (h *hash) Compare(hashed string, input string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashed), []byte(input))
}

func (h *hash) HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}
