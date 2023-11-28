package tgbotlocal

import (
	"github.com/cbrgm/githubevents/githubevents"
	"github.com/notneet/NNbot/internal/bot/tg-bot/controller"
	"github.com/notneet/NNbot/internal/bot/tg-bot/helper"
	"github.com/notneet/NNbot/internal/bot/tg-bot/service"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func RunTgBot(h *githubevents.EventHandler, botInstance *tgbotapi.BotAPI) {
	tgBot := helper.HandleTgBotInstance(botInstance, configTgBot())
	ghWebhookService := service.NewPulRequestService(tgBot)
	controller.HandlePullRequestEvent(h, ghWebhookService)

	// updates := botInstance.GetUpdatesChan(configTgBot())

	// for update := range updates {
	// 	if update.Message == nil {
	// 		continue
	// 	}
	// }
}

func configTgBot() tgbotapi.UpdateConfig {
	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 30

	return updateConfig
}
