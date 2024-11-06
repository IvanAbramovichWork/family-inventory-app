package main

import (
	"github.com/IvanAbramovichWork/family-inventory-app/app/config"
	"github.com/IvanAbramovichWork/family-inventory-app/app/database"
	"github.com/IvanAbramovichWork/family-inventory-app/app/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	// Загружаем конфигурацию
	cfg := config.NewConfig()

	// Инициализация базы данных
	db := database.InitDB(cfg)

	r := gin.Default()
	userHandler := handlers.NewUserHandler(db)

	r.POST("/users/signup", userHandler.RegisterUser)
	r.GET("/users/:id", userHandler.GetUser)

	r.Run(":8080")
}
