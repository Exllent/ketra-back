package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	// Регистрация маршрута /ping
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// Логирование всех запросов (для отладки)
	r.Use(gin.Logger())

	// Запуск сервера на порту 8080
	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}
