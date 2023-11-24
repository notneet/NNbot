package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/notneet/NNbot/runner"
)

func main() {
	errEnv := godotenv.Load()
	if errEnv != nil {
		panic(errEnv)
	}

	botName := os.Getenv("BOT_NAME")

	switch botName {
	case "tg-bot":
		runner.TgBot()
	case "tg-api":
		runner.TgAPI()
	default:
		log.Fatalf("Bot name not recognized: '%s'", botName)
	}
}
