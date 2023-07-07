package services

import (
	"io"
	"net/http"
	"os"
	"strconv"
	"time"

	helpers "github.com/notneet/NNbot/helpers"
)

func MakeAPIRequest(url string) (string, int, error) {
	timeoutStr := os.Getenv("HTTP_TIMEOUT")
	timeout, err := strconv.Atoi(timeoutStr)
	helpers.PanicIfError(err)

	client := &http.Client{
		Timeout: time.Duration(timeout) * time.Second,
	}

	resp, err := client.Get(url)
	if err != nil {
		return "", 404, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", 404, err
	}

	return string(body), resp.StatusCode, nil
}
