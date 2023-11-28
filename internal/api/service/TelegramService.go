package service

import telegramlocal "github.com/notneet/NNbot/internal/libs/common/telegram-local"

type TelegramService struct {
	tgMessageService *telegramlocal.SendMessageService
}

func NewTelegramService(tgLibs *telegramlocal.SendMessageService) *TelegramService {
	return &TelegramService{
		tgMessageService: tgLibs,
	}
}

func (tg *TelegramService) SendMessage(text string) error {
	err := tg.tgMessageService.SendMessage(text)

	return err
}
