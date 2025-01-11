package repositories

import (
	"github.com/herbetyp/crud-product-api/database"
	model "github.com/herbetyp/crud-product-api/models/user"
)

type LoginRepository struct {
}

func (r *LoginRepository) GetLogin(email string) (model.User, error) {
	db := database.GetDatabase()

	var u model.User

	err := db.First(&u, email).Error

	return u, err
}
