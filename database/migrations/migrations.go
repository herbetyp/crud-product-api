package migrations

import (
	pModel "github.com/herbetyp/crud-product-api/models/product"
	lModel "github.com/herbetyp/crud-product-api/models/login"
	"gorm.io/gorm"
)

func AutoMigrations(db *gorm.DB) {
	db.AutoMigrate(&pModel.Product{}, &lModel.Login{})
}
