package migrations

import (
	prodModel "github.com/herbetyp/crud-product-api/models/product"
	userModel "github.com/herbetyp/crud-product-api/models/user"
	"gorm.io/gorm"
)

func AutoMigrations(db *gorm.DB) {
	err := db.AutoMigrate(&prodModel.Product{}, &userModel.User{})

	if err != nil {
		panic(err)
	}
}
