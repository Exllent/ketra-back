package telegram

import (
	// "fmt"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"sync"
)

var bot *tgbotapi.BotAPI
var WG sync.WaitGroup
var ChatId int64

func InitBOT(tgToken string) {
	var err error
	bot, err = tgbotapi.NewBotAPI(tgToken)
	if err != nil {
		log.Fatal(err)
	}
}

func SendTelegramMessage(message string) {
	defer WG.Done()

	// Отправка сообщения в Telegram
	msg := tgbotapi.NewMessage(ChatId, message)
	_, err := bot.Send(msg)
	if err != nil {
		log.Printf("Ошибка при отправке сообщения: %v", err)
	}
}

func HandleTelegramUpdates() {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		log.Fatal(err)
	}

	for update := range updates {
		if update.Message == nil {
			continue
		}

		switch update.Message.Command() {
		case "help":
			handleHelpCommand(update.Message)
		default:
			handleUnknownCommand(update.Message)
		}
	}
}

// func handleStartCommand(message *tgbotapi.Message) {
// 	msg := tgbotapi.NewMessage(message.Chat.ID, "Добро пожаловать! Используйте команды /status для получения статуса или /help для помощи.")
// 	bot.Send(msg)
// }

// func handleStatusCommand(message *tgbotapi.Message) {
// 	// Логика для получения статуса заявки (например, из базы данных)
// 	ticketID := 123 // Пример ID заявки
// 	msg := tgbotapi.NewMessage(message.Chat.ID, fmt.Sprintf("Статус заявки с ID %d: В обработке", ticketID))
// 	bot.Send(msg)
// }

func handlerListTickets(message *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(message.Chat.ID, "")
	bot.Send(msg)
}

func handleHelpCommand(message *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(message.Chat.ID, "Доступные команды:\n/start - Начать\n/status - Статус заявки\n/help - Справка")
	bot.Send(msg)
}

func handleUnknownCommand(message *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(message.Chat.ID, "Неизвестная команда. Используйте /help для получения списка команд.")
	bot.Send(msg)
}
