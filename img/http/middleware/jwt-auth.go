package middleware

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/rizkysaputra4/moviwiki/img/env"
)

// JWTAuth ...
func JWTAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cookie, err := c.Cookie("Auth-Token")

		if err != nil {
			return c.JSON(http.StatusInternalServerError, "Error when parsing cookies")
		}

		_, errJWT := jwt.Parse(cookie.Value, func(token *jwt.Token) (interface{}, error) {
			return []byte(env.GetConfiguration().TokenKey), nil
		})

		if errJWT != nil {
			return c.JSON(http.StatusUnauthorized, "Token Invalid")
		}

		return next(c)
	}
}
