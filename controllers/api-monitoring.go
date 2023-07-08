package controllers

import (
	"bytes"
	"os"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/notneet/NNbot/helpers"
	"github.com/notneet/NNbot/helpers/stringlib"
	"github.com/notneet/NNbot/interfaces"
	"github.com/notneet/NNbot/services"
)

func HandleApiMonitoring(botInstance *tgbotapi.BotAPI) {
	var buf bytes.Buffer
	var err error
	apiUrl := os.Getenv("MAIN_API")
	textMsgSlow := stringlib.CreateStringTemplate("ApiSlow",
		"API {{.ApiUrl}} is slow (> 3 sec), with result {{.Elapsed}} secs and status code {{.StatusCode}}")
	textMsgErr := stringlib.CreateStringTemplate("ApiError", "API {{.ApiUrl}} has status code {{.StatusCode}}")

	startTime := time.Now()
	_, statusCode, _ := services.MakeAPIRequest(apiUrl)
	elapsed := time.Since(startTime)
	isThreeSecond := 3 * time.Second

	dataMsg := interfaces.MonitoringAPI{
		ApiUrl:     apiUrl,
		Elapsed:    elapsed.Seconds(),
		StatusCode: statusCode,
	}

	if elapsed > isThreeSecond {
		err = textMsgSlow.Execute(&buf, dataMsg)
		helpers.PanicIfError(err)
	} else if dataMsg.StatusCode != 200 {
		err = textMsgErr.Execute(&buf, dataMsg)
		helpers.PanicIfError(err)
	}

	if buf.Len() > 0 {
		msg := tgbotapi.NewMessage(helpers.GetChatID(), buf.String())
		_, err = botInstance.Send(msg)
		helpers.PanicIfError(err)
	}
}
