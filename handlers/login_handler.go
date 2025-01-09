package handlers

import (
	"github.com/herbetyp/crud-product-api/models"
	"github.com/herbetyp/crud-product-api/repositories"
	"github.com/herbetyp/crud-product-api/services"
)

func LoginHandler(l *models.Login) (string, error) {
	user, err := repositories.UsersLoginRepository(l)
	if err != nil {
		return "", err
	}

	if user.Password != services.SHA512Encoder(l.Password) {
		return "", err
	}

	token, err := services.GenerateToken(user.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}
