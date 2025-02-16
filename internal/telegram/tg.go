package telegram

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

type Telegram struct {
	TelegramToken string
}

func NewTelegram(telegram_token string) *Telegram {
	return &Telegram{telegram_token}
}

func (tg *Telegram) Send(chatId int64, message string) error {
	bot, err := tgbotapi.NewBotAPI(tg.TelegramToken)
	if err != nil {
		return fmt.Errorf("SendTelegramNot: %v", err)
	}

	msg := tgbotapi.NewMessage(chatId, message)
	_, err = bot.Send(msg)
	if err != nil {
		return fmt.Errorf("SendTelegramNot: %v", err)
	}
	log.Printf("Sending a message to a user %d: %s", chatId, message)
	return nil
}
