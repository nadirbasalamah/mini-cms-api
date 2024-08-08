package utils

import (
	"mini-cms-api/middlewares"
	"mini-cms-api/models"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(userID int, jwtOptions models.JWTOptions) (string, error) {
	expire := jwt.NewNumericDate(time.Now().Local().Add(time.Hour * time.Duration(int64(jwtOptions.ExpiresDuration))))

	claims := &middlewares.JWTCustomClaims{
		ID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: expire,
		},
	}

	rawToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := rawToken.SignedString([]byte(jwtOptions.SecretKey))

	if err != nil {
		return "", err
	}

	return token, nil
}
