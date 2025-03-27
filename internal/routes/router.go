package routes

import (
	"golang-api/internal/routes/product"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	product.SetupProductRoutes(app)
}
