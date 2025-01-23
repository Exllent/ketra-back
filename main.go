package main

import (
	"github.com/joho/godotenv"
	"ketra-back/db"
	"ketra-back/models"
	"ketra-back/routers"
	"ketra-back/telegram"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")
	dbSslMode := os.Getenv("DB_SSLMODE")
	tgToken := os.Getenv("TG_TOKEN")


	db.InitDB(dbHost, dbUser, dbPassword, dbName, dbPort, dbSslMode)
	telegram.ChatId = -4740762266
	telegram.InitBOT(tgToken)
	telegram.WG.Add(1)
	go telegram.HandleTelegramUpdates()

	db.DB.AutoMigrate(&models.Ticket{})
	r := gin.Default()
	routes.RegisterTicketRoutes(r)

	r.Run(":8080")
}
