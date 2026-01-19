package main

import (
	"task-management/internal/config"
)

func main() {
	config.ConnectDatabase()

}

// TASK REPO TESTING ===

// 	value := model.Task{
// 	Title:   "Task 23",
// 	Description: "Tassdasdas",
// 	ProjectID: 6,
// 	Priority: "High",
// 	Status:   "Done",
// 	AssignedID: 65,
// 	DueDate: "2024-12-31",
// }

// taskRepo := repository.NewTaskRepository(config.DB)
// err := taskRepo.Create( &value)
// if err != nil {
// 	panic(err)
// }

// taskRepo := repository.NewTaskRepository(config.DB)
// err := taskRepo.Update(3, &value)
// if err != nil {
// 	panic(err)
// }

// TaskRepo := repository.NewTaskRepository(config.DB)
// res, err := TaskRepo.FindByID(6, 3)
// if err != nil {
// 	panic(err)
// }
// println(res.Title)

// TaskRepo := repository.NewTaskRepository(config.DB)
// res, err := TaskRepo.FindByProjectID(3)
// if err != nil {
// 	panic(err)
// }
// for _, T := range res {
// 	println(T.Title)
// }

// taskRepo := repository.NewTaskRepository(config.DB)
// err := taskRepo.DeleteById(6)
// if err != nil {
// 	panic(err)
// }


// PROJECT REPO TESTING ===

// value := model.Project{
// 	Title:   "Project Deleted",
// 	OwnerID: 4,
// }

// projectRepo := repository.NewProjectRepository(config.DB)
// err := projectRepo.Create(&value)
// if err != nil {
// 	panic(err)
// }

// projectRepo := repository.NewProjectRepository(config.DB)
// err := projectRepo.DeleteById(7)
// if err != nil {
// 	panic(err)
// }