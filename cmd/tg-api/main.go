package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/notneet/NNbot/internal/api"
	telegramlocal "github.com/notneet/NNbot/internal/libs/common/telegram-local"
)

func main() {
	errEnv := godotenv.Load()
	if errEnv != nil {
		log.Fatal("Error loading .env file")
	}

	port := getPort()
	tgToken, chatId, err := getTelegramConfig()
	if err != nil {
		fmt.Println("Error:", err.Error())
		return
	}

	tgMessageService, err := telegramlocal.Setup(tgToken, chatId)
	if err != nil {
		fmt.Println("Error:", err.Error())
		return
	}

	app := fiber.New()
	api.SetupRoutes(app, tgMessageService)

	log.Fatal(app.Listen(":" + port))
}

func getPort() string {
	port := os.Getenv("API_PORT")

	if port == "" {
		// Default to 8080 if PORT is not set in .env
		port = "8080"
	}

	return port
}

func getTelegramConfig() (string, int64, error) {
	telegramTokenAPI := os.Getenv("TELEGRAM_APITOKEN")
	chatIdStr := os.Getenv("CHAT_ID")

	if telegramTokenAPI == "" {
		return "", 0, errors.New("TELEGRAM_APITOKEN is missing")
	}

	if chatIdStr == "" {
		return "", 0, errors.New("TELEGRAM_CHAT_ID is missing")
	}

	chatID, err := strconv.ParseInt(chatIdStr, 10, 64)
	if err != nil {
		return "", 0, fmt.Errorf("error parsing TELEGRAM_CHAT_ID: %s", err.Error())
	}

	return telegramTokenAPI, chatID, nil
}
