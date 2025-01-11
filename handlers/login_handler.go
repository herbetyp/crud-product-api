package handlers

import (
	model "github.com/herbetyp/crud-product-api/models/user"
	"github.com/herbetyp/crud-product-api/repositories"
	"github.com/herbetyp/crud-product-api/services"
)

func LoginHandler(l *model.Login) (string, error) {
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
