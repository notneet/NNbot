package service

import telegramlocal "github.com/notneet/NNbot/internal/libs/common/telegram-local"

type TelegramService struct {
	telegramLibs *telegramlocal.SendMessageService
}

func NewTelegramService(tgLibs *telegramlocal.SendMessageService) *TelegramService {
	return &TelegramService{
		telegramLibs: tgLibs,
	}
}

func (tg *TelegramService) SendMessage(text string) error {
	err := tg.telegramLibs.SendMessage(text)

	return err
}
