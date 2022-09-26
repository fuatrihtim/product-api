package main

import (
	"github.com/gofiber/fiber/v2"
	"os"
	"strconv"

	"github.com/fuatrihtim/product-api/pkg/utils"
	"github.com/fuatrihtim/product-api/platform/migrations"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	if !fiber.IsChild() {
		migrations.Migrate()
	}

	port, _ := strconv.Atoi(os.Getenv("APP_PORT"))

	utils.CreateServer(port)
}
