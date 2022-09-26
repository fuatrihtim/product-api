package database

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func initializePostgres() *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s sslmode=disable timezone=Europe/Amsterdam",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_NAME"),
	)

	connection, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil
	}
	return connection
}
