package service

import "errors"

var (
	ErrUserAlreadyExists 	= errors.New("User already exists")
	ErrUserNotFound      	= errors.New("User not found")
	ErrUserDeleted	   		= errors.New("User already deleted")
	ErrInvalidInput			= errors.New("Invalid input")

	ErrProjectNotFound		= errors.New("Project not found")
	ErrProjectDeleted		= errors.New("Project already deleted")
	ErrProjectAlreadyExists	= errors.New("Project already exists")

	ErrTaskNotFound			= errors.New("Task not found")
	ErrTaskDeleted			= errors.New("Task already deleted")
)