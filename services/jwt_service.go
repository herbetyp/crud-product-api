package services

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/herbetyp/crud-product-api/config"
	"github.com/herbetyp/crud-product-api/config/logger"
)

func GenerateToken(id uint) (string, error) {
	JWTConf := config.GetConfig().JWT

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
		"sub": fmt.Sprint(id),
		"iss": "auth-product-api",
		"aud": "api://product-api",
		"exp": time.Now().Add(time.Duration(JWTConf.ExpiresIn) * time.Second).Unix(),
		"iat": time.Now().Unix(),
		"jti": uuid.Must(uuid.NewRandom()).String(),
	})

	t, err := token.SignedString([]byte(JWTConf.SecretKey))

	if err != nil {
		return "", err
	}

	return t, nil
}
func ValidateToken(token string, uid string) (bool, error) {
	conf := config.GetConfig()

	// Validate token
	tokenDecoded, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, isValid := t.Method.(*jwt.SigningMethodHMAC); !isValid {
			return false, nil
		}

		return []byte(conf.JWT.SecretKey), nil
	})

	if err != nil {
		logger.Error("invalid token", err)
		return false, nil
	}

	// Validate claims
	claims, _ := tokenDecoded.Claims.(jwt.MapClaims)

	if claims["iss"] != "auth-product-ap" || claims["aud"] != "api://product-api" {
		logger.Error("invalid claim", err)
		return false, nil
	}

	return true, nil
}
