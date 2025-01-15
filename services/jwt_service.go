package services

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	config "github.com/herbetyp/crud-product-api/internal/configs"
	"github.com/herbetyp/crud-product-api/internal/configs/logger"
)

func GenerateToken(id uint) (string, string, error) {
	JWTConf := config.GetConfig().JWT
	tokenId := uuid.Must(uuid.NewRandom()).String()

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
		"sub": fmt.Sprint(id),
		"iss": "auth-product-api",
		"aud": "api://product-api",
		"exp": time.Now().Add(time.Duration(JWTConf.ExpiresIn) * time.Second).Unix(),
		"iat": time.Now().Unix(),
		"jti": tokenId,
	})

	t, err := token.SignedString([]byte(JWTConf.SecretKey))

	if err != nil {
		return "", "", err
	}

	return t, tokenId, nil
}
func ValidateToken(token string) (bool, string, string, error) {
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
		return false, "", "", nil
	}

	// Validate claims
	claims, _ := tokenDecoded.Claims.(jwt.MapClaims)

	if claims["iss"] != "auth-product-api" || claims["aud"] != "api://product-api" {
		logger.Error("invalid claim", err)
		return false, "", "", nil
	}

	jwtId := claims["jti"].(string)
	SubUserID := claims["sub"].(string)

	return true, jwtId, SubUserID, nil
}
