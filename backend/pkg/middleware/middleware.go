package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/timoyo93/auth-backend/pkg/db"
	"net/http"
	"strings"
)

func CheckCookie(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cookie, err := c.Cookie("JSESSIONID")
		if err != nil {
			if strings.Contains(err.Error(), "named cookie not present") {
				return c.JSON(http.StatusUnauthorized, "No cookie present")
			}
			return err
		}
		if isTokenPresent := db.CheckForAccessToken(cookie.Value); !isTokenPresent {
			return c.JSON(http.StatusUnauthorized, "Please register/login first")
		}
		return next(c)
	}
}
