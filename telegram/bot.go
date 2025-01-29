package telegram

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"sync"
)

var Bot *tgbotapi.BotAPI
var WG sync.WaitGroup

func InitBOT(tgToken string) {
	var err error
	Bot, err = tgbotapi.NewBotAPI(tgToken)
	if err != nil {
		log.Fatal(err)
	}
}

func SendTelegramMessage(message string) {
	defer WG.Done()
	if _, err := Bot.Send(tgbotapi.NewMessage(ChatId, message)); err != nil {
		log.Printf("Ошибка при отправке сообщения: %v", err)
	}
}

func isAdmin(userID int) bool {
	for _, id := range AdminIDs {
		if id == userID {
			return true
		}
	}
	return false
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
			if !isAdmin(update.CallbackQuery.From.ID) {
				continue
			}
			handleCallbackQuery(&update)
		}

		if update.Message != nil {
			if !isAdmin(update.Message.From.ID) {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "⛔ Вы не имеете доступа к этому боту")
				if _, err := Bot.Send(msg); err != nil {
					log.Println(err)
				}
				continue
			}
			handleMessage(&update)
		}
	}
}
