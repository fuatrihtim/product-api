package routes

import (
	"github.com/fuatrihtim/product-api/app/controllers/product"
	"github.com/gofiber/fiber/v2"
)

func ProductRoutes(app *fiber.App) {
	productGroup := app.Group("/products")
	productGroup.Get("/", product.List)
	productGroup.Get("/:id", product.Get)
	productGroup.Post("/", product.Create)
	productGroup.Post("/:id", product.Update)
	productGroup.Delete("/:id", product.Delete)
}
