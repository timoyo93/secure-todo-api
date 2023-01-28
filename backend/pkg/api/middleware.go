package api

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"strings"
)

func (a API) CheckCookie(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cookie, err := c.Cookie(CookieName)
		if err != nil {
			if strings.Contains(err.Error(), "named cookie not present") {
				return c.JSON(http.StatusUnauthorized, "No cookie present")
			}
			return err
		}
		if err := a.authService.CheckAccessToken(cookie.Value); err != nil {
			return c.JSON(http.StatusUnauthorized, "Please register/login first")
		}
		return next(c)
	}
}

func (a API) RegisterMiddleware() API {
	loggerConfig := middleware.LoggerConfig{
		Format: "${status} - ${method} ${uri}\t ->\t took ${latency_human}\n",
	}
	a.srv.Use(middleware.LoggerWithConfig(loggerConfig))
	a.srv.Use(middleware.Recover())
	a.srv.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowCredentials: true,
	}))
	return a
}
