package main

import (
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

	port, _ := strconv.Atoi(os.Getenv("3000"))

	utils.CreateServer(port)
}
