package service

import (
	"log"

	"github.com/google/go-github/v50/github"
	"github.com/notneet/NNbot/internal/bot/tg-bot/helper"
)

type PullRequestService struct {
	botInstace *helper.TelegramBot
}

func NewPulRequestService(botInstace *helper.TelegramBot) *PullRequestService {
	return &PullRequestService{
		botInstace: botInstace,
	}
}

func (s *PullRequestService) ProcessOnPullRequestEventEdited(deliveryID, eventName string, event *github.PullRequestEvent) error {

	log.Println(deliveryID)
	s.botInstace.SendMessage(deliveryID)
	// log.Printf("[%s]", event.PullRequest.GetTitle())
	// log.Printf("%s has edited pull #%d\n", event.Sender.GetLogin(), event.PullRequest.GetNumber())

	return nil
}
