package handlers

import (
	"github.com/herbetyp/crud-product-api/models"
	"github.com/herbetyp/crud-product-api/repositories"
)

func CreateUserHandler(u models.UserModel) (string, error) {
	username, err := repositories.CreateUserRepository(u)
	if err != nil {
		return username, err
	}

	return username, nil
}

func UpdateUserPassHandler(id int, p string) (int, error) {
	userId, err := repositories.UpdateUserPassRepository(id, p)
	if err != nil {
		return userId, err
	}

	return userId, nil
}
