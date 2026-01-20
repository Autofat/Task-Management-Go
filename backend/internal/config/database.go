package config

import (
	"fmt"
	"os"
	"task-management/internal/model"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	err := godotenv.Load()
	if err != nil {
		panic("Failed to load .env file")
	}
	dsn := os.Getenv("DSN")
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!")
	}else {
		fmt.Println("Database connected successfully.")
	}

	DB = database
	DB.AutoMigrate(
		&model.User{},
		&model.Project{},
		&model.Task{},
	)
	fmt.Println("Success Migrate")
}