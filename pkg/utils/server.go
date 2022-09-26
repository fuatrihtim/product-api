package utils

import (
	"fmt"
	"github.com/fuatrihtim/product-api/pkg/middleware"
	"github.com/fuatrihtim/product-api/pkg/routes"
	"github.com/fuatrihtim/product-api/platform/server"
	"github.com/gofiber/fiber/v2"
	"log"
)

func CreateServer(port int) {
	app := fiber.New(fiber.Config{
		Prefork:      true,
		ErrorHandler: server.ErrorHandler,
	})

	file := middleware.Logger(app)
	defer file.Close()

	routes.ProductRoutes(app)

	log.Fatal(app.Listen(fmt.Sprintf(":%d", port)))
}
