package main

import (
	"os"
	"task-management/internal/config"
	"task-management/internal/handler"
	"task-management/internal/repository"
	"task-management/internal/routes"
	"task-management/internal/service"

	"github.com/gin-gonic/gin"
)



func main() {
	config.ConnectDatabase()
	
	userRepository := repository.NewUserRepository(config.DB)
	projectRepository := repository.NewProjectRepository(config.DB)
	projectMemberRepository := repository.NewProjectMemberRepository(config.DB)
	taskRepository := repository.NewTaskRepository(config.DB)
	
	passwordService := service.NewPasswordService()
	userService := service.NewUserService(userRepository, passwordService)
	projectService := service.NewProjectService(userRepository, projectRepository, projectMemberRepository)
	projectMemberService := service.NewProjectMemberService(userRepository, projectRepository, projectMemberRepository)
	taskService := service.NewTaskService(taskRepository, projectRepository, userRepository)
	
	userHandler := handler.NewUserHandler(userService)
	projectHandler := handler.NewProjectHandler(projectService)
	projectMemberHandler := handler.NewProjectMemberHandler(projectMemberService)
	taskHandleer := handler.NewTaskHandler(taskService)

	router := gin.Default()
	
	routes.SetupCheckRoutes(router)
	routes.SetupUserRoutes(router, userHandler)
	routes.SetupProjectRoutes(router, projectHandler, projectMemberHandler)
	routes.SetupTaskRoutes(router, taskHandleer)

	port := os.Getenv("PORT")
	if port == "" {
		port = "4444"
	}

  router.Run(":" + port)
}
