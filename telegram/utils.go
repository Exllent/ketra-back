package telegram

import "github.com/go-telegram-bot-api/telegram-bot-api"

type UserState struct {
	Offset int
}

type RemaingCommandHandler struct {
	Pattern []string
	Handler func(update *tgbotapi.Update, ticketID uint)
}

var userStates = make(map[int64]*UserState)
var ChatId int64
var AdminIDs [2]int
