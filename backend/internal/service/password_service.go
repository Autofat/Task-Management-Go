package service

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type PasswordService struct{}

func NewPasswordService() *PasswordService {
	return &PasswordService{}
}

func (ps *PasswordService) HashPassword(rawPassword string) ([]byte, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(rawPassword), 14)
	if err != nil {
		return nil, err
	}
	return hashedPassword, nil
}

func (ps *PasswordService) CompareHashAndPassword(hashedPassword []byte, rawPassword string) error {
	err := bcrypt.CompareHashAndPassword(hashedPassword, []byte(rawPassword))
	if err != nil {
		return errors.New("password does not match")
	}
	return nil
}