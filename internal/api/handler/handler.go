package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/notneet/NNbot/internal/api/service"
)

// HomeHandler handles the home route
type HomeHandler struct {
	greetService *service.GreetService
}

// NewHomeHandler creates a new instance of HomeHandler
func NewHomeHandler(greetService *service.GreetService) *HomeHandler {
	return &HomeHandler{greetService: greetService}
}

// Home is the handler for the home route
func (h *HomeHandler) Home(c *fiber.Ctx) error {
	message := h.greetService.GetWelcomeMessage()
	return c.SendString(message)
}

// HelloHandler handles the /hello route
type HelloHandler struct {
	greetService *service.GreetService
}

// NewHelloHandler creates a new instance of HelloHandler
func NewHelloHandler(greetService *service.GreetService) *HelloHandler {
	return &HelloHandler{greetService: greetService}
}

// Hello is the handler for the /hello route
func (h *HelloHandler) Hello(c *fiber.Ctx) error {
	message := h.greetService.GetHelloMessage()
	return c.SendString(message)
}
