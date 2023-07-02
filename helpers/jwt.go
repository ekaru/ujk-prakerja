package helpers

import (
	"errors"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func GenerateToken(id_user uint, username string) (string, error) {
	claims := jwt.MapClaims{
		"id":       id_user,
		"username": username,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("secretKey")))
}

func VerifyToken(ctx echo.Context) (interface{}, error) {
	headerToken := ctx.Request().Header.Get("Authorization")
	bearer := strings.HasPrefix(headerToken, "Bearer")

	if !bearer {
		return nil, errors.New("bearer token not found")
	}

	stringToken := headerToken[7:]

	token, err := jwt.Parse(stringToken, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("failed to get sign token")
		}

		return []byte(os.Getenv("secretKey")), nil
	})

	if err != nil {
		return nil, errors.New("invalid token")
	}

	if _, ok := token.Claims.(jwt.MapClaims); !ok {
		return nil, errors.New("failed to parse claims")
	}

	return token.Claims.(jwt.MapClaims), nil
}
