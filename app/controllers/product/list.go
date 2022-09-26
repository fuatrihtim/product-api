package product

import (
	"github.com/fuatrihtim/product-api/app/models"
	"github.com/fuatrihtim/product-api/platform/database"
	"github.com/gofiber/fiber/v2"
)

func List(ctx *fiber.Ctx) error {
	products := []models.Product{}
	database.Connection().Find(&products)

	return ctx.JSON(products)
}
