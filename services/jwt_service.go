package services

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofrs/uuid"
)

type jwtService struct {
	secretKey string
	issuer    string
}

func NewJWTService() *jwtService {
	return &jwtService{
		secretKey: "Bj1ZWf19UA9sGE621nFy9eJTFfJzpiDmxPM_MDKVKCT40ZodEW5TT8mH3ww8Oyd8",
		issuer:    "auth-product-api",
	}
}


type Claims struct {
	Sub string `json:"sub"`
	JTI string `json:"jti"`
	jwt.StandardClaims
}

func (s *jwtService) GenerateToken(id int) (string, error) {
	claims := &Claims{
		Sub: fmt.Sprint(id),
		JTI: uuid.Must(uuid.NewV4()).String(),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 1).Unix(),
			Issuer:    s.issuer,
			IssuedAt:  time.Now().Unix(),
			Audience:  "product-api",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	t, err := token.SignedString([]byte(s.secretKey))
	
	if err != nil {
		return "", err
	}
	return t, nil
}

func (s *jwtService) ValidateToken(token string) bool {
	_, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, isValid := t.Method.(*jwt.SigningMethodHMAC); !isValid {
			return nil, fmt.Errorf("invalid token: %v", token)
		}

		return []byte(s.secretKey), nil
	})

	return err == nil
}