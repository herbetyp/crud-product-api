package services

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofrs/uuid"
	"github.com/herbetyp/crud-product-api/configs"
)



type Claims struct {
	Sub string `json:"sub"`
	JTI string `json:"jti"`
	jwt.StandardClaims
}


func GenerateToken(id int) (string, error) {
	conf := configs.GetConfig()
	claims := &Claims{
		Sub: fmt.Sprint(id),
		JTI: uuid.Must(uuid.NewV4()).String(),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 1).Unix(),
			Issuer:    "auth-product-api",
			IssuedAt:  time.Now().Unix(),
			Audience:  "product-api",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	t, err := token.SignedString([]byte(conf.JWT.SecretKey))
	
	if err != nil {
		return "", err
	}
	return t, nil
}

func ValidateToken(token string) bool {
	conf := configs.GetConfig()
	_, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, isValid := t.Method.(*jwt.SigningMethodHMAC); !isValid {
			return nil, fmt.Errorf("invalid token: %v", token)
		}

		return []byte(conf.JWT.SecretKey), nil
	})

	return err == nil
}