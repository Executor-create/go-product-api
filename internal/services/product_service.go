package services

import (
	"golang-api/internal/database"
	"golang-api/internal/model"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func GetProductById(id string) (model.Product, error) {
	var product model.Product

	db := database.Database
	db.Find(&product, "id = ?", id)

	return product, nil
}

func GetProductByName(name string) (model.Product, error) {
	var product model.Product

	db := database.Database
	db.Find(&product, "name = ?", name)

	return product, nil
}

func GetSimilarProducts(id string) ([]model.Product, model.Product, error) {
	var product model.Product

	db := database.Database
	db.Find(&product, "id = ?", id)

	if product.Categories == "" {
		return nil, product, fiber.NewError(fiber.StatusBadRequest, "Product has no categories to find similar products.")
	}

	categoryIDs := strings.Split(product.Categories, ",")

	var similarProducts []model.Product
	db.Where("categories LIKE ?", "%"+categoryIDs[0]+"%").Not("id", product.ID).Limit(10).Find(&similarProducts)

	return similarProducts, product, nil
}
