package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/notneet/NNbot/internal/api/handler"
	"github.com/notneet/NNbot/internal/api/service"
)

// SetupRoutes initializes API routes
func SetupRoutes(app *fiber.App) {
	greetService := service.NewGreetService()

	// Inject the service into handlers
	homeHandler := handler.NewHomeHandler(greetService)
	helloHandler := handler.NewHelloHandler(greetService)

	app.Get("/", homeHandler.Home)
	app.Get("/hello", helloHandler.Hello)
}
