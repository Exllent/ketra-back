package telegram

import (
	"fmt"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"ketra-back/models"
	"log"
	"strconv"
	"strings"
)

func handleCommand(update *tgbotapi.Update) {
	chatID := update.Message.Chat.ID
	userStates[chatID] = &UserState{Offset: 0}
	limit := 5
	tickets, err := models.GetTickets(limit, 0)
	if err != nil {
		log.Printf("Ошибка при получении заявок: %v", err)
		return
	}
	totalTickets, err := models.GetTotalTicketCount()
	if err != nil {
		log.Printf("Ошибка при получении общего количества заявок: %v", err)
		return
	}
	ticketText := "Заявки:\n\n"
	for _, ticket := range tickets {
		ticketText += fmt.Sprintf(
			"ID: %v\nФИО: %v\nНомер: %v\nПочта: %v\nНаличие пожелания: %v\n\n",
			ticket.ID, ticket.FIO, ticket.PhoneNumber, ticket.Email, checkWishlist(ticket.Wishlist),
		)
	}
	keyboard := GetKeyboard(totalTickets, limit, 0)
	msg := tgbotapi.NewMessage(chatID, ticketText)
	msg.ReplyMarkup = keyboard
	Bot.Send(msg)
}

func handleCallbackQuery(update *tgbotapi.Update) {
	chatID := update.CallbackQuery.Message.Chat.ID
	data := update.CallbackQuery.Data
	state, exists := userStates[chatID]
	if !exists {
		state = &UserState{Offset: 0}
		userStates[chatID] = state
	}
	limit := 5
	var newOffset int
	switch data {
	case "next":
		newOffset = state.Offset + limit
	case "prev":
		newOffset = state.Offset - limit
		if newOffset < 0 {
			newOffset = 0
		}
	case "close":
		Bot.DeleteMessage(tgbotapi.DeleteMessageConfig{ChatID: chatID, MessageID: update.CallbackQuery.Message.MessageID})
		return
	default:
		newOffset = 0
	}
	tickets, err := models.GetTickets(limit, newOffset)
	if err != nil {
		log.Printf("Ошибка при получении заявок: %v", err)
		return
	}
	totalTickets, err := models.GetTotalTicketCount()
	if err != nil {
		log.Printf("Ошибка при получении общего количества заявок: %v", err)
		return
	}
	ticketText := "Заявки:\n\n"
	for _, ticket := range tickets {
		ticketText += fmt.Sprintf(
			"ID: %v\nФИО: %v\nНомер: %v\nПочта: %v\nНаличие пожелания: %v\n\n",
			ticket.ID, ticket.FIO, ticket.PhoneNumber, ticket.Email, checkWishlist(ticket.Wishlist),
		)
	}
	keyboard := GetKeyboard(totalTickets, limit, newOffset)
	editMsg := tgbotapi.NewEditMessageText(chatID, update.CallbackQuery.Message.MessageID, ticketText)
	editMsg.ReplyMarkup = &keyboard
	Bot.Send(editMsg)
	userStates[chatID].Offset = newOffset
}

func checkWishlist(s string) bool {
	return len(s) != 0
}

func detailTicketHandler(update *tgbotapi.Update, ticketID uint) {
	ticket, err := models.GetTicketByID(ticketID)
	var text string
	if err != nil {
		text = fmt.Sprintf("Не удалось найти ticket с id: %v", ticketID)
	} else {
		text = fmt.Sprintf(
			"ID: %v\nФИО: %v\nНомер: %v\nПочта: %v\nПожелание: %s",
			ticket.ID, ticket.FIO, ticket.PhoneNumber, ticket.Email, ticket.Wishlist,
		)
	}
	msg := tgbotapi.NewMessage(ChatId, text)
	Bot.Send(msg)

}

func deleteTicketHandler(update *tgbotapi.Update, ticketID uint) {
	err := models.DeleteTicketByID(ticketID)
	var text string
	if err != nil {
		text = fmt.Sprintf("Не удалось удалить ticket с id: %v", ticketID)
	} else {
		text = fmt.Sprintf("Удалось удалить ticket с id: %v", ticketID)
	}
	msg := tgbotapi.NewMessage(ChatId, text)
	Bot.Send(msg)
}

func handleMessage(update *tgbotapi.Update) {
	if update.Message.Command() == "tickets" {
		handleCommand(update)
	} else if update.Message.Command() == "help" {
		handleHelpCommand(update.Message)
	} else {
		handlers := []RemaingCommandHandler{
			{Pattern: []string{"/view", "ticket"}, Handler: detailTicketHandler},
			{Pattern: []string{"/delete", "ticket"}, Handler: deleteTicketHandler},
		}
		if !remaingCommands(update, handlers) {
			handleUnknownCommand(update.Message)
		}
	}
}

func remaingCommands(update *tgbotapi.Update, handlers []RemaingCommandHandler) bool {
	cmd := update.Message.Text
	for _, handler := range handlers {
		row_command := strings.Split(cmd, "-")
		if len(row_command) == 3 && row_command[0] == handler.Pattern[0] && row_command[1] == handler.Pattern[1] {
			ticketID, err := strconv.Atoi(row_command[2])
			if err != nil {
				log.Println("Error converting ticket number to integer:", err)
				return false
			}
			handler.Handler(update, uint(ticketID))
			return true
		}
	}
	return false
}

// func handleStartCommand(message *tgbotapi.Message) {
// 	msg := tgbotapi.NewMessage(message.Chat.ID, "Добро пожаловать! Используйте команды /status для получения статуса или /help для помощи.")
// 	bot.Send(msg)
// }

func handleHelpCommand(message *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(message.Chat.ID, "Доступные команды:\n/start - Начать\n/status - Статус заявки\n/help - Справка")
	Bot.Send(msg)
}

func handleUnknownCommand(message *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(message.Chat.ID, "Неизвестная команда. Используйте /help для получения списка команд.")
	Bot.Send(msg)
}
