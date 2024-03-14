package main

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/notneet/NNbot/cmd/nquiz/controller"
	"github.com/notneet/NNbot/pkg/common"
	"github.com/notneet/NNbot/pkg/telegram"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	err := godotenv.Load()
	common.FatalIfError(err)

	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_API_TOKEN"))
	common.PanicIfError(err)
	log.Printf("Authorized on account %s", bot.Self.UserName)

	myChatId, err := strconv.ParseInt(os.Getenv("CHAT_ID"), 10, 64)
	common.PanicIfError(err)

	if appEnv := os.Getenv("APP_ENV"); appEnv == "development" {
		bot.Debug = true
	}

	nQuizController := controller.NewQuizController(bot)

	nQuizController.Monitor()

	telegram.TgSendMessage(bot, myChatId, "Test")

	select {}
}
