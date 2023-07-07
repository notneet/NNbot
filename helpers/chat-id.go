package helpers

import (
	"os"
	"strconv"
)

func GetChatID() int64 {
	chatId := os.Getenv("CHAT_ID")
	numChatId, err := strconv.Atoi(chatId)
	PanicIfError(err)

	return int64(numChatId)
}
