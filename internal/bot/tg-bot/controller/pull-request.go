package controller

import (
	"github.com/cbrgm/githubevents/githubevents"
	"github.com/notneet/NNbot/internal/bot/tg-bot/service"
)

func HandlePullRequestEvent(handler *githubevents.EventHandler, service *service.PullRequestService) {
	handler.OnPullRequestEventEdited(service.ProcessOnPullRequestEventEdited)
}
