package services

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	config "github.com/herbetyp/crud-product-api/internal/configs"
	"github.com/herbetyp/crud-product-api/internal/configs/logger"
)

func GetJwtClaims(tokenString string) (jwt.MapClaims, error) {
	token, _, _ := jwt.NewParser().ParseUnverified(tokenString, jwt.MapClaims{})
	claims, _ := token.Claims.(jwt.MapClaims)

	return claims, nil
}

func GenerateFingerprint(f string) string {
	return SHA512Encoder(f)
}

func GenerateToken(id uint, fingerprintPlainTex string) (string, string, error) {
	JWTConf := config.GetConfig().JWT
	tokenId := uuid.Must(uuid.NewRandom()).String()

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
		"sub":         fmt.Sprint(id),
		"fingerprint": fingerprintPlainTex,
		"iss":         "auth-product-api",
		"aud":         "api://product-api",
		"exp":         time.Now().Add(time.Duration(JWTConf.ExpiresIn) * time.Second).Unix(),
		"iat":         time.Now().Unix(),
		"jti":         tokenId,
	})

	t, err := token.SignedString([]byte(JWTConf.SecretKey))

	if err != nil {
		return "", "", err
	}

	return t, tokenId, nil
}

func ValidateToken(token string) (bool, jwt.MapClaims, error) {
	conf := config.GetConfig()

	// Validate token
	tokenDecoded, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, isValid := t.Method.(*jwt.SigningMethodHMAC); !isValid {
			return false, nil
		}

		return []byte(conf.JWT.SecretKey), nil
	})

	if err != nil {
		logger.Error("invalid token: ", err)
		return false, jwt.MapClaims{}, err
	}

	claims, _ := GetJwtClaims(tokenDecoded.Raw)

	// Validate claims
	if claims["iss"] != "auth-product-api" || claims["aud"] != "api://product-api" {
		logger.Error("invalid claim: ", err)
		return false, jwt.MapClaims{}, err
	}

	return true, claims, nil
}
