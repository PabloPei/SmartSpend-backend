package auth

import (
	"time"

	"github.com/PabloPei/SmartSpend-backend/conf"
	"github.com/golang-jwt/jwt/v5"
)

type UserJWT struct {
	UserId   string
	Email    string
	UserName string
}

func CreateJWT(secret []byte, user UserJWT) (string, error) {

	expiration := time.Duration(conf.ServerConfig.JWTExpirationInSeconds) * time.Second
	expirationTime := time.Now().UTC().Add(expiration).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID":    user.UserId,
		"email":     user.Email,
		"userName":  user.UserName,
		"expiresAt": expirationTime,
	})

	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
