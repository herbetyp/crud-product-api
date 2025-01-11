package migrations

import (
	lModel "github.com/herbetyp/crud-product-api/models/login"
	pModel "github.com/herbetyp/crud-product-api/models/product"
	uModel "github.com/herbetyp/crud-product-api/models/user"
	"gorm.io/gorm"
)

func AutoMigrations(db *gorm.DB) {
	db.AutoMigrate(&pModel.Product{}, &lModel.Login{}, &uModel.User{})
}
