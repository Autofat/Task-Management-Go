package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string `json:"email" type:varchar(255) gorm:"uniqueIndex:idx_email_deleted_at;not null"`
	Password []byte `json:"password" gorm:"not null"`
	Fullname string `json:"fullname" type:varchar(255) gorm:"not null"`
	Role     string `json:"role" type:varchar(50) gorm:"not null"`
}

func (User *User) SetPassword(password string, passwordService interface {
	HashPassword(rawPassword string) (string, error)
}) error {
	hashedPassword, err := passwordService.HashPassword(password)
	if err != nil {
		return err
	}
	User.Password = []byte(hashedPassword)
	return nil
	
}
