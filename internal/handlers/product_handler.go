package handlers

import (
	"golang-api/internal/helpers/responses"
	"golang-api/internal/services"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func GetProductById(c *fiber.Ctx) error {
	id := c.Params("id")
	product, err := services.GetProductById(id)

	if err != nil || product.ID == "" {
		return c.Status(fiber.StatusNotFound).JSON(responses.ErrorResponse{
			Status:  http.StatusNotFound,
			Message: "Product not found",
			Data:    nil,
		})
	}

	return c.JSON(product)
}

func GetProductByName(c *fiber.Ctx) error {
	name := c.Query("name")
	product, err := services.GetProductByName(name)

	if err != nil || product.Name == "" {
		return c.Status(fiber.StatusNotFound).JSON(responses.ErrorResponse{
			Status:  http.StatusNotFound,
			Message: "Product not found",
			Data:    nil,
		})
	}

	return c.JSON(product)
}

func GetSimilarProducts(c *fiber.Ctx) error {
	id := c.Params("id")
	similarProducts, product, err := services.GetSimilarProducts(id)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return c.JSON(fiber.Map{
		"product":          product,
		"similar_products": similarProducts,
	})
}
