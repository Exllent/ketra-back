package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"ketra-back/db"
	"ketra-back/models"
	"ketra-back/routers"
	"log"
	// "os"
)

func main() {
	// Загрузка переменных окружения из файла .env
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// // Инициализация базы данных
	db.InitDB()

	// Автомиграция моделей
	db.DB.AutoMigrate(&models.Ticket{})

	// Создание маршрутов
	r := gin.Default()
	routes.RegisterTicketRoutes(r)

	// Запуск сервера
	r.Run(":8080")
}
