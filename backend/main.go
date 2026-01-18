package main

import (
	"task-management/internal/config"
	"task-management/internal/model"
	"task-management/internal/repository"
)

func main() {
	config.ConnectDatabase()

	value := model.Project{
		OwnerID: 4,
		Title:   "Test Project 2",
	}
	
	projectRepo := repository.NewProjectRepository(config.DB)
	err := projectRepo.Update(4, &value)
	if err != nil {
		panic(err)
	}
}