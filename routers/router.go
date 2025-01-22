package routes

import (
	"github.com/gin-gonic/gin"
	"ketra-back/db"
	"ketra-back/models"
	"ketra-back/validation"
	"net/http"
)

func RegisterTicketRoutes(router *gin.Engine) {
	// Регистрация маршрутов для билетов
	router.POST("/tickets", createTicket)

	// Регистрация health check маршрута
	router.GET("/health", healthCheck)
}

func createTicket(c *gin.Context) {
	var ticket models.Ticket
	if err := c.ShouldBindJSON(&ticket); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Валидация данных
	if err := validation.ValidateTicket(ticket); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := db.DB.Create(&ticket)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"ticket": ticket})
}

func healthCheck(c *gin.Context) {
	// Проверка соединения с базой данных
	var test string
	result := db.DB.Raw("SELECT 1").Scan(&test)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "Database connection failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "ok", "message": "Server is running"})
}
