package migrations

import (
	model "github.com/herbetyp/crud-product-api/models/product"
	"gorm.io/gorm"
)

func AutoMigrations(db *gorm.DB) {
	db.AutoMigrate(&model.Product{})
}