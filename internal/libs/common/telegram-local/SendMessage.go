package telegramlocal

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// SendMessageService is a service for sending messages via Telegram
type SendMessageService struct {
	bot    *tgbotapi.BotAPI
	chatID int64
}

// NewSendMessageService creates a new instance of SendMessageService
func NewSendMessageService(botToken string, chatID int64) (*SendMessageService, error) {
	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		return nil, err
	}

	log.Println("Telegram API are set!")

	return &SendMessageService{
		bot:    bot,
		chatID: chatID,
	}, nil
}

// SendMessage sends a text message to the specified chat ID
func (s *SendMessageService) SendMessage(text string) error {
	msg := tgbotapi.NewMessage(s.chatID, text)
	_, err := s.bot.Send(msg)

	return err
}
