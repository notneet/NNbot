package helper

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func TgSendMessage(bot *tgbotapi.BotAPI, chatID int64, text string) {
	msg := tgbotapi.NewMessage(chatID, text)
	// msg.ReplyToMessageID = int(chatID)

	_, err := bot.Send(msg)
	PanicIfError(err)
}
