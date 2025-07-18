package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secret []byte

func Init(secretKey string) {
	secret = []byte(secretKey)
}

func Generate(userId string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userId,
		"exp":     time.Now().Add(128 * time.Hour).Unix(),
	}
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return t.SignedString(secret)
}

func Validate(tokenStr string) (bool, string) {
	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (any, error) {
		return secret, nil
	})
	if err != nil || !token.Valid {
		return false, ""
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return false, ""
	}
	userId, ok := claims["user_id"].(string)
	return ok, userId
}
