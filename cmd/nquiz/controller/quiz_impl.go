package controller

import (
	"fmt"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/notneet/NNbot/pkg/telegram"
)

type QuizControllerImpl struct {
	BotInstance *tgbotapi.BotAPI
}

func NewQuizController(bot *tgbotapi.BotAPI) QuizController {
	return &QuizControllerImpl{
		BotInstance: bot,
	}
}

func (q QuizControllerImpl) Monitor() {
	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 30
	updates := q.BotInstance.GetUpdatesChan(updateConfig)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		if strings.HasPrefix(update.Message.Text, "/") {
			switch update.Message.Text {
			case "/start":
				q.HandleStartComannd(update)
				continue
			default:
				telegram.TgSendMessage(q.BotInstance, update.Message.Chat.ID, "What do you mean?")
			}
		}

		fmt.Println(update.Message.Chat.UserName+":", update.Message.Text)

	}
}

func (q QuizControllerImpl) TriggerQuiz(uChan tgbotapi.Update) {

}

func (q QuizControllerImpl) HandleStartComannd(uChan tgbotapi.Update) {
	telegram.TgSendMessage(q.BotInstance, uChan.FromChat().ChatConfig().ChatID, "??")
}
