package telegram

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

func GetKeyboard(totalTickets int64, limit, offset int) tgbotapi.InlineKeyboardMarkup {
	var keyboard tgbotapi.InlineKeyboardMarkup
	var nextButton, prevButton tgbotapi.InlineKeyboardButton

	if offset+limit < int(totalTickets) {
		nextButton = tgbotapi.NewInlineKeyboardButtonData("➡️", "next")
	} else {
		nextButton = tgbotapi.NewInlineKeyboardButtonData("🚫", "not_next")
	}

	if offset > 0 {
		prevButton = tgbotapi.NewInlineKeyboardButtonData("⬅️", "prev")
	} else {
		prevButton = tgbotapi.NewInlineKeyboardButtonData("🚫", "not_prev")
	}

	keyboard.InlineKeyboard = append(
		keyboard.InlineKeyboard,
		[]tgbotapi.InlineKeyboardButton{
			prevButton,
			tgbotapi.NewInlineKeyboardButtonData("❌", "close"),
			nextButton,
		},
	)
	return keyboard
}
