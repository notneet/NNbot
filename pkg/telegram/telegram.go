package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/notneet/NNbot/pkg/common"
)

func TgSendMessage(bot *tgbotapi.BotAPI, chatID int64, text string) {
	msg := tgbotapi.NewMessage(chatID, text)
	// msg.ReplyToMessageID = int(chatID)

	_, err := bot.Send(msg)
	common.PanicIfError(err)
}
