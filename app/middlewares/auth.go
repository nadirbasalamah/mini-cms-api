package middlewares

import (
	"errors"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

type JWTCustomClaims struct {
	ID int `json:"id"`
	jwt.RegisteredClaims
}

type JWTConfig struct {
	SecretKey       string
	ExpiresDuration int
}

func (jwtCfg *JWTConfig) Init() echojwt.Config {
	return echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(JWTCustomClaims)
		},
		SigningKey: []byte(jwtCfg.SecretKey),
	}
}

func (jwtCfg *JWTConfig) GenerateToken(userID int) (string, error) {
	expire := jwt.NewNumericDate(time.Now().Local().Add(time.Hour * time.Duration(int64(jwtCfg.ExpiresDuration))))

	claims := &JWTCustomClaims{
		ID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: expire,
		},
	}

	rawToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := rawToken.SignedString([]byte(jwtCfg.SecretKey))

	if err != nil {
		return "", err
	}

	return token, nil
}

func GetUser(c echo.Context) (*JWTCustomClaims, error) {
	user := c.Get("user").(*jwt.Token)

	if user == nil {
		return nil, errors.New("invalid token")
	}

	claims := user.Claims.(*JWTCustomClaims)

	return claims, nil
}

func VerifyToken(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		userData, err := GetUser(c)

		isInvalid := userData == nil || err != nil

		if isInvalid {
			return c.JSON(http.StatusUnauthorized, map[string]string{
				"message": "invalid token",
			})
		}

		return next(c)
	}
}
