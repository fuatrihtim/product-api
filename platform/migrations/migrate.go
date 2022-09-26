package migrations

import (
	"github.com/fuatrihtim/product-api/app/models"
	"github.com/fuatrihtim/product-api/platform/database"
)

func Migrate() {
	database.Connection().AutoMigrate(&models.Product{})
}
