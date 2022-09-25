package product

import (
	"github.com/fuatrihtim/product-api/app/models"
	"github.com/fuatrihtim/product-api/platform/database"
)

func Create(ctx *fiber.Ctx) error {
	product := models.Product{}
	if err := ctx.BodyParser(&product); err != nil {
		return err
	}

	if len(product.Code) < 1 {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "Code is required",
		})
	}

	if err := database.Connection().Create(&product).Error; err != nil {
		return err
	}

	return ctx.JSON(product)
}
