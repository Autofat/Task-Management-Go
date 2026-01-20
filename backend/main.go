package main

import (
	"net/http"
	"os"
	"task-management/internal/config"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDatabase()

	router := gin.Default()
	
	router.GET("/health", func (c *gin.Context)  {
		sqlDB, err := config.DB.DB()
			if err != nil{
				c.JSON(http.StatusInternalServerError, gin.H{
					"status": "Unhealthy",
					"database": "disconnected",
				})
			}

			if err := sqlDB.Ping(); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
				"status": "Unhealthy",
				"database": "error",
			})
			}
			
			c.JSON(http.StatusOK, gin.H{
				"status": "Healthy",
				"database": "connected",
			})
	})
	router.GET("/ping", func (c *gin.Context)  {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "4444"
	}

  router.Run(":" + port)
}
