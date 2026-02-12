package main

import (
	"os"
	"task-management/internal/config"
	"task-management/internal/handler"
	"task-management/internal/repository"
	"task-management/internal/routes"
	"task-management/internal/service"

	"github.com/gin-contrib/cors"
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
	authService := service.NewAuthService(userRepository, passwordService)
	projectService := service.NewProjectService(userRepository, projectRepository, projectMemberRepository)
	projectMemberService := service.NewProjectMemberService(userRepository, projectRepository, projectMemberRepository)
	taskService := service.NewTaskService(taskRepository, projectRepository, userRepository)
	
	userHandler := handler.NewUserHandler(userService)
	authHandler := handler.NewAuthHandler(authService)
	projectHandler := handler.NewProjectHandler(projectService, projectMemberService)
	projectMemberHandler := handler.NewProjectMemberHandler(projectMemberService)
	taskHandler := handler.NewTaskHandler(taskService, projectMemberService)

	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:5173"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Authorization"}
	config.AllowCredentials = true

	router.Use(cors.New(config))
	
	routes.SetupCheckRoutes(router)
	routes.SetupUserRoutes(router, userHandler)
	routes.SetupAuthRoutes(router, authHandler) 
	routes.SetupProjectRoutes(router, projectHandler, projectMemberHandler)
	routes.SetupTaskRoutes(router, taskHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "4444"
	}

  router.Run(":" + port)
}
