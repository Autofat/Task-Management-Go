package model

import "gorm.io/gorm"

type Project struct {
	gorm.Model
	Title 	string `json:"title" gorm:"not null"`
	OwnerID uint   `json:"owner_id" gorm:"not null;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Owner   User   `json:"owner" gorm:"foreignKey:OwnerID"`
}