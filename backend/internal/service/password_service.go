package service

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type PasswordService struct{}

func NewPasswordService() *PasswordService {
	return &PasswordService{}
}

func (ps *PasswordService) HashPassword(rawPassword string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(rawPassword), 14)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func (ps *PasswordService) CompareHashAndPassword(hashedPassword string, rawPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(rawPassword))
	if err != nil {
		return errors.New("password does not match")
	}
	return nil
}