package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/notneet/NNbot/internal/api"
)

func main() {
	errEnv := godotenv.Load()
	if errEnv != nil {
		log.Fatal("Error loading .env file")
	}

	port := getPort()

	app := fiber.New()
	api.SetupRoutes(app)

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
