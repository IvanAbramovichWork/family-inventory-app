package main

import (
	"github.com/IvanAbramovichWork/family-inventory-app/app/config"
	"github.com/IvanAbramovichWork/family-inventory-app/app/database"

	"github.com/gin-gonic/gin"
)

func main() {
	// Загружаем конфигурацию
	cfg := config.NewConfig()

	// Инициализация базы данных
	database.InitDB(cfg)

	// Создаем Gin роутер
	router := gin.Default()

	// Простой эндпоинт для проверки работы сервера
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// Запуск сервера
	router.Run(":8080") // Порт можно настроить
}
