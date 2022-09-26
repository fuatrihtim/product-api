package product

import (
	"github.com/fuatrihtim/product-api/app/models"
	"github.com/fuatrihtim/product-api/platform/database"
	"github.com/gofiber/fiber/v2"
)

func Update(ctx *fiber.Ctx) error {
	request := &models.Product{}

	if err := ctx.BodyParser(&request); err != nil {
		return err
	}

	product := &models.Product{}

	err := database.Connection().First(&product, "id = ?", ctx.Params("id")).Error
	if err != nil {
		return err
	}

	err = database.Connection().Model(&product).Updates(&models.Product{
		Id:            request.Id,
		Code:          request.Code,
		Name:          request.Name,
		StockQuantity: request.StockQuantity,
	}).Error

	if err != nil {
		return err
	}

	return ctx.JSON(product)
}
