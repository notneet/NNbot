package tgbot

import (
	"fmt"
	"os"
)

func RunTgBot() {
	fmt.Printf("TgBot http timeout is: %s\n", os.Getenv("HTTP_TIMEOUT"))
}
