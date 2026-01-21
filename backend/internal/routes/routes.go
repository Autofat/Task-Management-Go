package routes

import (
	"net/http"
	"task-management/internal/config"
	"task-management/internal/handler"

	"github.com/gin-gonic/gin"
)

func SetupCheckRoutes(router *gin.Engine) {
	router.GET("/health", func (c *gin.Context)  {
		sqlDB, err := config.DB.DB()
			if err != nil{
				c.JSON(http.StatusInternalServerError, gin.H{
					"status": "Unhealthy",
					"database": "disconnected",
				})
			return

			}

			if err := sqlDB.Ping(); err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{
					"status": "Unhealthy",
					"database": "error",
				})
			return
			}
			
			c.JSON(http.StatusOK, gin.H{
				"status": "Healthy",
				"database": "connected",
			})
	})
	router.GET("/ping", func (c *gin.Context)  {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
}

func SetupUserRoutes(router *gin.Engine, h *handler.UserHandler) {
	users := router.Group("/users")
	{
		users.POST("", h.RegisterUser)
		users.GET("/:id", h.GetUser)
		users.PUT("/:id", h.UpdateUser)
		users.DELETE("/:id", h.DeleteUser)
	}

}

func SetupProjectRoutes(router *gin.Engine, h *handler.ProjectHandler) {
	projects := router.Group("/projects")
	{
		projects.POST("", h.CreateProject)
		projects.GET("/:id", h.GetProjectByID)
		projects.GET("", h.GetProjectsByOwnerID)
		projects.PUT("/:id", h.UpdateProject)
		projects.DELETE("/:id", h.DeleteProject)
	}
}

func SetupTaskRoutes(router *gin.Engine, h *handler.TaskHandler) {
	tasks := router.Group("/tasks")
	{
		tasks.POST("", h.CreateTask)
		tasks.GET("/:id", h.GetTaskByID)
		tasks.GET("", h.GetTasksByProjectID)
		tasks.PUT("/:id", h.UpdateTask)
		tasks.DELETE("/:id", h.DeleteTask)
	}
}