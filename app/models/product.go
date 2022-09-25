package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Product struct {
	Id            string `gorm:"primary_key"`
	Code          string `gorm:"unique;not null"`
	Name          string
	StockQuantity int
}

func (product *Product) BeforeCreate(tx *gorm.DB) error {
	product.Id = uuid.NewString()
	return nil
}
