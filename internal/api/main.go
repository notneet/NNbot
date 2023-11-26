package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/notneet/NNbot/internal/api/handler"
	"github.com/notneet/NNbot/internal/api/service"
	telegramlocal "github.com/notneet/NNbot/internal/libs/common/telegram-local"
)

// SetupRoutes initializes API routes
func SetupRoutes(app *fiber.App, tgMessageService *telegramlocal.SendMessageService) {
	setupRootRoute(app)
	setupTelegramRoute(app, tgMessageService)
}

func setupRootRoute(app *fiber.App) {
	greetService := service.NewGreetService()

	// Inject the service into handlers
	homeHandler := handler.NewHomeHandler(greetService)
	helloHandler := handler.NewHelloHandler(greetService)

	app.Get("/", homeHandler.Home)
	app.Get("/hello", helloHandler.Hello)
}

func setupTelegramRoute(app *fiber.App, tgMessageService *telegramlocal.SendMessageService) {
	telegramService := service.NewTelegramService(tgMessageService)

	telegramHandler := handler.NewTelegramHandler(telegramService)

	app.Get("/telegram/send/:msg", telegramHandler.SendMessage)
}
