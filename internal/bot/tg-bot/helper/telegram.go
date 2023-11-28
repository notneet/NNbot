package helper

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type TelegramBot struct {
	botInstance *tgbotapi.BotAPI
}

func HandleTgBotInstance(botInstance *tgbotapi.BotAPI, botConfig tgbotapi.UpdateConfig) *TelegramBot {
	botInstance.GetUpdatesChan(botConfig)

	return &TelegramBot{
		botInstance: botInstance,
	}
}

func (t *TelegramBot) SendMessage(text string) {
	msg := tgbotapi.NewMessage(727321498, text)

	_, err := t.botInstance.Send(msg)
	PanicIfError(err)
}
