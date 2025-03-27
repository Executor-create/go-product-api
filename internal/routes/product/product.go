package product

import (
	"golang-api/internal/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupProductRoutes(router fiber.Router) {
	router.Get("/products/:id", handlers.GetProductById)
	router.Get("/products", handlers.GetProductByName)
	router.Get("/products/:id/similar", handlers.GetSimilarProducts)
}
