package product

import (
	"github.com/fuatrihtim/product-api/app/models"
	"github.com/fuatrihtim/product-api/platform/database"
	"github.com/gofiber/fiber/v2"
)

func Delete(ctx *fiber.Ctx) error {
	product := &models.Product{}

	err := database.Connection().First(&product, "id = ?", ctx.Params("id")).Error
	if err != nil {
		return err
	}

	err = database.Connection().Delete(&product).Error
	if err != nil {
		return err
	}

	return ctx.JSON(fiber.Map{
		"message": "Product deleted successfully",
	})
}
