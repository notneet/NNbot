package main

import (
	"log"
	"net/http"
	"os"

	"github.com/cbrgm/githubevents/githubevents"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	tgbotlocal "github.com/notneet/NNbot/internal/bot/tg-bot"
)

func main() {
	errEnv := godotenv.Load()
	if errEnv != nil {
		log.Fatal("Error loading .env file")
	}

	ghWebhookHandler := githubevents.New("Tonghilapnya_17!!")
	bot, err := tgbotapi.NewBotAPI(getBotToken())
	if err != nil {
		log.Fatal(err)
	}

	tgbotlocal.RunTgBot(ghWebhookHandler, bot)

	// add a http handleFunc
	http.HandleFunc("/webhook", func(w http.ResponseWriter, r *http.Request) {
		err := ghWebhookHandler.HandleEventRequest(r)
		if err != nil {
			log.Fatal("error")
		}
	})

	go func() {
		if err := http.ListenAndServe(getWebhookPort(), nil); err != nil {
			panic(err)
		}
	}()

	log.Printf("Bot is running and listening at '%s'!", getWebhookPort())

	// Keep the program running (e.g., use a blocking select{})
	select {}
}

func getWebhookPort() string {
	port := os.Getenv("WEBHOOK_PORT")

	if port == "" {
		// Default to 8080 if PORT is not set in .env
		port = ":8080"
	}

	return ":" + port
}

func getBotToken() string {
	token := os.Getenv("TELEGRAM_APITOKEN")

	if token == "" {
		panic("please set telegram token first!")
	}

	return token
}
