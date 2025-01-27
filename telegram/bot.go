package telegram

import (
	"log"
	"sync"
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

var Bot *tgbotapi.BotAPI
var WG sync.WaitGroup
var ChatId int64

func InitBOT(tgToken string) {
	var err error
	Bot, err = tgbotapi.NewBotAPI(tgToken)
	if err != nil {
		log.Fatal(err)
	}
}

func SendTelegramMessage(message string) {
	defer WG.Done()
	msg := tgbotapi.NewMessage(ChatId, message)
	_, err := Bot.Send(msg)
	if err != nil {
		log.Printf("Ошибка при отправке сообщения: %v", err)
	}
}

func HandleTelegramUpdates() {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates, err := Bot.GetUpdatesChan(u)
	if err != nil {
		log.Fatal(err)
	}
	for update := range updates {
		if update.CallbackQuery != nil {
			handleCallbackQuery(&update)
		}
		if update.Message != nil {
			handleMessage(&update)
		}
	}
}
