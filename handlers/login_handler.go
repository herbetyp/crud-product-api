package handlers

import (
	"encoding/json"
	"fmt"

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
	var u userModel.User
	cacheKey := USER_PREFIX + l.Email

	if cachedData := services.GetCache(cacheKey); cachedData != "" {
		err := json.Unmarshal([]byte(cachedData), &u)
		if err != nil {
			u, err = h.repository.GetLogin(l.Email)
			if err != nil {
				return "", "", err
			}
			cacheValue, _ := json.Marshal(u)
			services.SetCache(cacheKey, string(cacheValue))
		}
	}

	if u.Password != services.SHA512Encoder(l.Password) {
		return "", "", fmt.Errorf("invalid password")
	}

	token, tokenId, err := services.GenerateToken(uint(u.ID))
	if err != nil {
		return "", "", err
	}

	return token, tokenId, nil
}
func NewLoginHandler(r interfaces.ILoginRepository) *LoginHandler {
	return &LoginHandler{
		repository: r,
	}
}
