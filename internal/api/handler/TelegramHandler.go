package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
	"github.com/notneet/NNbot/internal/api/service"
)

type TelegramHandler struct {
	telegramService *service.TelegramService
}

func NewTelegramHandler(telegramService *service.TelegramService) *TelegramHandler {
	return &TelegramHandler{telegramService: telegramService}
}

func (h *TelegramHandler) SendMessage(c *fiber.Ctx) error {
	paramMsg := utils.CopyString(c.Params("msg"))

	err := h.telegramService.SendMessage(paramMsg)

	return err
}
