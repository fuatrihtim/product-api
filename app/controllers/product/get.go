package product

import (
	"github.com/fuatrihtim/product-api/app/models"
	"github.com/fuatrihtim/product-api/platform/database"
	"github.com/gofiber/fiber/v2"
)

func Get(ctx *fiber.Ctx) error {
	product := &models.Product{}

	err := database.Connection().First(&product, "id = ?", ctx.Params("id")).Error
	if err != nil {
		return err
	}

	return ctx.JSON(product)
}
