package main

import (
	"task-management/internal/config"
	"task-management/internal/repository"
)

func main() {
	config.ConnectDatabase()

	taskRepo := repository.NewUserRepository(config.DB)
	taskRepo.FindByID(2)


	
}