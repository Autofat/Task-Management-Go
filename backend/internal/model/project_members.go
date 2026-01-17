package model

import "gorm.io/gorm"

type ProjectMember struct {
	gorm.Model
	ProjectID 	uint `json:"project_id" gorm:"not null;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Project   	Project `json:"project" gorm:"foreignKey:ProjectID"`
	UserID    	uint    `json:"user_id" gorm:"not null;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	User    	User    `json:"user" gorm:"foreignKey:UserID"`
	Role 		string  `json:"role" gorm:"not null"`
}