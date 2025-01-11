package handlers

import (
	model "github.com/herbetyp/crud-product-api/models/user"
	"github.com/herbetyp/crud-product-api/repositories"
)

func CreateUserHandler(u model.User) (string, error) {
	username, err := repositories.CreateUserRepository(u)
	if err != nil {
		return username, err
	}

	return username, nil
}

func UpdateUserPassHandler(id int, p string) (int, error) {
	userId, err := repositories.UpdateUserPassRepository(id, p)
	if err != nil {
		return 0, err
	}

	return userId, nil
}

func DeleteUserHandler(id int) (int, error) {
	userId, err := repositories.DeleteUserRepository(id)
	if err != nil {
		return 0, err
	}

	return userId, nil
}
