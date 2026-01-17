package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string `json:"email" type:varchar(255) gorm:"uniqueIndex;not null"`
	Password []byte `json:"password" gorm:"not null"`
	Fullname string `json:"fullname" type:varchar(255) gorm:"not null"`
	Role     string `json:"role" type:varchar(50) gorm:"not null"`
}

// func (User *User) setPassword(password []byte) {
// 	hashedPassword := bcrypt.GenerateFromPassword([]byte(password), 14)
// 	User.Password = hashedPassword
// }
