package main

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/notneet/NNbot/cmd/nquiz/controller"
	"github.com/notneet/NNbot/helper"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	err := godotenv.Load()
	helper.FatalIfError(err)

	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_API_TOKEN"))
	helper.PanicIfError(err)
	log.Printf("Authorized on account %s", bot.Self.UserName)

	myChatId, err := strconv.ParseInt(os.Getenv("CHAT_ID"), 10, 64)
	helper.PanicIfError(err)

	if appEnv := os.Getenv("APP_ENV"); appEnv == "development" {
		bot.Debug = true
	}

	nQuizController := controller.NewQuizController(bot)

	nQuizController.Monitor()

	helper.TgSendMessage(bot, myChatId, "Test")

	select {}
}
