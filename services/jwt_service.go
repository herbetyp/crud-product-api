package services

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/herbetyp/crud-product-api/configs"
)

func GenerateToken(id int) (string, error) {
	conf := configs.GetConfig()

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
		"sub": fmt.Sprint(id),          
		"iss": "auth-product-api",          
		"aud": "api://product-api",    
		"exp": time.Now().Add(time.Hour).Unix(),
		"iat": time.Now().Unix(),
		"jti": uuid.Must(uuid.NewRandom()).String(),
	})

	t, err := token.SignedString([]byte(conf.JWT.SecretKey))
	if err != nil {
		return "", err
	}

	return t, nil
}

func ValidateToken(token string, uid string) (bool, string) {
	conf := configs.GetConfig()

	// Validate token
	tokenDecoded, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, isValid := t.Method.(*jwt.SigningMethodHMAC); !isValid {
			return false, nil
		}

		return []byte(conf.JWT.SecretKey), nil
	})

	if err != nil {
		fmt.Printf("invalid token: %s\n", err)
		return false, ""
	}

	// Validate token claims
	claims, _ := tokenDecoded.Claims.(jwt.MapClaims)
	if claims["iss"] != "auth-product-api" || claims["aud"] != "api://product-api" {
		fmt.Printf("invalid claims\n")
		return false, ""
	}

	return true, claims["sub"].(string)
}
