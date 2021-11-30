package middleware

import (
	"net/http"
	"strings"
	"time"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

const key = "legal"

func CreateToken(userId int) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["userId"] = userId
	claims["exp"] = time.Now().Add(time.Hour).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(key))

}

func AuthJWTMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authorizationFromHeader := c.Request().Header.Get("authorization")
		if authorizationFromHeader == "" {
			return c.String(http.StatusForbidden, "Forbidden")
		}

		tokenString := strings.ReplaceAll(authorizationFromHeader, "Bearer ", "")

		claims := jwt.MapClaims{}
		token, err := jwt.ParseWithClaims(tokenString, &claims, func(t *jwt.Token) (interface{}, error) {
			return []byte(key), nil
		})
		if err != nil && !token.Valid {
			return c.String(http.StatusForbidden, "Token Salah")
		}

		c.Set("email", claims["userId"])
		return next(c)
	}
}