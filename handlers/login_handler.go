package handlers

import (
	"encoding/json"
	"fmt"

	"github.com/herbetyp/crud-product-api/internal/configs/logger"
	"github.com/herbetyp/crud-product-api/internal/interfaces"
	loginModel "github.com/herbetyp/crud-product-api/models/login"
	userModel "github.com/herbetyp/crud-product-api/models/user"
	"github.com/herbetyp/crud-product-api/services"
)

const (
	USER_PREFIX = "user>>>"
)

type LoginHandler struct {
	repository interfaces.ILoginRepository
}

func (h *LoginHandler) NewLogin(l loginModel.LoginDTO) (string, string, error) {
	var user userModel.User

	fing := services.GenerateFingerprint(l.Email)
	cacheKey := USER_PREFIX + fing

	cacheData := services.GetCache(cacheKey)
	if cacheData != "" {
		err := json.Unmarshal([]byte(cacheData), &user)
		if err != nil {
			return "", "", err
		}
	} else {
		u, err := h.repository.GetLogin(l.Email)
		if err != nil {
			logger.Error("error on get user from database: %v", err)
			return "", "", err
		}
		cacheValue, _ := json.Marshal(u)
		services.SetCache(cacheKey, string(cacheValue))
		user = u
	}

	passwordMatch := services.SHA512Encoder(l.Password)
	if user.Password != passwordMatch {
		return "", "", fmt.Errorf("invalid password")
	}

	token, tokenId, err := services.GenerateToken(user.ID, fing)
	if err != nil {
		logger.Error("Error on generate token: %v", err)
		return "", "", err
	}

	return token, tokenId, nil
}
func NewLoginHandler(r interfaces.ILoginRepository) *LoginHandler {
	return &LoginHandler{
		repository: r,
	}
}
