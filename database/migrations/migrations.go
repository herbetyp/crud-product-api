package migrations

import (
	pModel "github.com/herbetyp/crud-product-api/models/product"
	uModel "github.com/herbetyp/crud-product-api/models/user"
	"gorm.io/gorm"
)

func AutoMigrations(db *gorm.DB) {
	db.AutoMigrate(&pModel.Product{}, &uModel.User{})
}
