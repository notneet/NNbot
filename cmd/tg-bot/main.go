package main

import (
	"log"

	"github.com/joho/godotenv"
	tgbot "github.com/notneet/NNbot/internal/bot/tg-bot"
)

func main() {
	errEnv := godotenv.Load()
	if errEnv != nil {
		log.Fatal("Error loading .env file")
	}

	tgbot.RunTgBot()
}
