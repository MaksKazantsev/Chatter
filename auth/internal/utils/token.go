package utils

import (
	"github.com/golang-jwt/jwt/v4"
	"os"
	"time"
)

var key = os.Getenv("secret_key")

func NewToken(uuid string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS512)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = uuid
	claims["exp"] = time.Now().Add(time.Hour * 24 * 30)

	tokenString, err := token.SignedString(key)
	if err != nil {
		return "", NewError(err.Error(), ErrInternal)
	}
	return tokenString, nil
}
func ParseToken(token string) (jwt.MapClaims, error) {
	t, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})
	if err != nil {
		return nil, NewError(err.Error(), ErrInternal)
	}
	claims, ok := t.Claims.(jwt.MapClaims)
	if claims.Valid() != nil || !ok {
		return nil, NewError(err.Error(), ErrInternal)
	}
	return claims, nil
}
