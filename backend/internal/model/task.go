package model

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	Title       	string `json:"title" gorm:"not null"`
	Description 	string `json:"description"`
	Priority    	string `json:"priority" gorm:"not null"`
	Status      	string `json:"status" gorm:"not null"`
	ProjectID   	uint   `json:"project_id" gorm:"not null;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Project     	Project `json:"project" gorm:"foreignKey:ProjectID"`
	AssignedID  	uint   `json:"assigned_id" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Assigned_to    	User   `json:"assigned_to" gorm:"foreignKey:AssignedID"`
	DueDate	 		string `json:"due_date"`

}