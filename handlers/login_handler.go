package handlers

import (
	"github.com/herbetyp/crud-product-api/internal/interfaces"
	model "github.com/herbetyp/crud-product-api/models/login"
	"github.com/herbetyp/crud-product-api/services"
)

type LoginHandler struct {
	repository interfaces.ILoginRepository
}

func (h *LoginHandler) NewLogin(l model.LoginDTO) (string, error) {
	user, err := h.repository.GetLogin(l.Email)
	if err != nil {
		return "", err
	}

	if user.Password != services.SHA512Encoder(l.Password) {
		return "", err
	}

	token, err := services.GenerateToken(uint(user.ID))
	if err != nil {
		return "", err
	}

	return token, nil
}
func NewLoginHandler(r interfaces.ILoginRepository) *LoginHandler {
	return &LoginHandler{
		repository: r,
	}
}
