package main

import (
	"golang-api/config"
	"golang-api/internal/database"
	"golang-api/internal/routes"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	if err := config.LoadEnv(); err != nil {
		log.Fatalf("Error loading environment variables: %v", err)
	}

	if err := database.Connect(); err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	app := fiber.New()
	app.Use(cors.New())

	routes.SetupRoutes(app)

	app.Listen(":3000")
}
