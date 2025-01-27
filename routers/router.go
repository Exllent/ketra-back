package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"ketra-back/db"
	"ketra-back/models"
	"ketra-back/telegram"
	"ketra-back/validation"
	"net/http"
)

func RegisterTicketRoutes(router *gin.Engine) {
	// Регистрация маршрутов для билетов
	router.GET("/health", healthCheck)
	v1 := router.Group("/api/v1")
	ticketsGroup := v1.Group("/tickets")
	{
		ticketsGroup.POST("", createTicket)
		ticketsGroup.GET("/:id", receiveTicket)
		ticketsGroup.GET("", receiveTickets)
		ticketsGroup.PUT("/:id", updateTicket)
		ticketsGroup.DELETE("/:id", deleteTicket)
	}
}

func receiveTicket(c *gin.Context) {

}

func deleteTicket(c *gin.Context) {

}

func updateTicket(c *gin.Context) {

}

func receiveTickets(c *gin.Context) {

}

func createTicket(c *gin.Context) {
	var ticket models.Ticket
	if err := c.ShouldBindJSON(&ticket); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(ticket)
	if err := validation.ValidateTicket(&ticket); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := ticket.Create(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	ticketID := ticket.ID
	Wishlist := ticket.Wishlist
	if Wishlist == "" {
		Wishlist = "Без пожелания"
	}
	telegram.WG.Add(1)
	go telegram.SendTelegramMessage(
		fmt.Sprintf(
			"Заявка с ID %d успешно создана!\nФИО: %s\nПочта: %s\nНомер телефона: %s\nПожелание заказчика: %s",
			ticketID,
			ticket.FIO,
			ticket.Email,
			ticket.PhoneNumber,
			Wishlist,
		),
	)
	c.JSON(http.StatusCreated, gin.H{"ticket": ticket})
}

func healthCheck(c *gin.Context) {
	var test string
	result := db.DB.Raw("SELECT 1").Scan(&test)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "Database connection failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "ok", "message": "Server is running"})
}
