package handlers

import (
	"fmt"

	"github.com/herbetyp/crud-product-api/internal/interfaces"
	model "github.com/herbetyp/crud-product-api/models/login"
	"github.com/herbetyp/crud-product-api/services"
)

const MySecret string = "abc&1*~#^2^#s0^=)^^7%b34"

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

	fingerprint, err := services.EncryptData(fmt.Sprintf("%s:%s", user.Email, user.UId), MySecret)
	if err != nil {
		return "", err
	}

	token, err := services.GenerateToken(uint(user.ID), fingerprint)
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
