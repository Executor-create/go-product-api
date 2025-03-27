package main

import (
	"fmt"
	"golang-api/config"
	"golang-api/internal/database"
	"log"

	"gorm.io/gen"
)

func main() {
	if err := config.LoadEnv(); err != nil {
		log.Fatalf("Error loading environment variables: %v", err)
	}

	if err := database.Connect(); err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	g := gen.NewGenerator(gen.Config{OutPath: "./internal/query", Mode: gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface})

	g.UseDB(database.Database)

	g.ApplyBasic(
		g.GenerateModel("products"),
		g.GenerateModel("categories"),
		g.GenerateModel("reviews"),
		g.GenerateModel("generated_reviews"),
		g.GenerateModel("product_categories"),
	)

	g.ApplyBasic(
		g.GenerateAllTable()...,
	)

	g.Execute()

	fmt.Println("Models generated successfully")
}
