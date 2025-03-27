package database

import (
	"fmt"
	"golang-api/config"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Database *gorm.DB

func Connect() error {
	var DATABASE_URI = config.GetEnvVariable("DATABASE_URI")

	var err error
	Database, err = gorm.Open(postgres.Open(DATABASE_URI), &gorm.Config{})

	if err != nil {
		return fmt.Errorf("failed to connect to the database: %w", err)
	}

	log.Println("Connected to the database")

	return nil
}
