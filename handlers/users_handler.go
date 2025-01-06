package handlers

import (
	"github.com/herbetyp/crud-product-api/models"
	"github.com/herbetyp/crud-product-api/repositories"
)


func CreateUserHandler(user models.User) (string, error) {
	username, err := repositories.CreateUserRepository(user)
	if err != nil {
		return username, err
	}

	return username, nil
}