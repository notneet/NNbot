package controller

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type QuizController interface {
	Monitor()
	TriggerQuiz(uChan tgbotapi.Update)
	HandleStartComannd(uChan tgbotapi.Update)
}
