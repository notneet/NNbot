package telegramlocal

// Setup initializes the Telegram library with the provided bot token and chat ID
func Setup(botToken string, chatID int64) (*SendMessageService, error) {
	telegramService, err := NewSendMessageService(botToken, chatID)
	if err != nil {
		panic(err)
	}

	// Additional initialization logic if needed

	return telegramService, nil
}
