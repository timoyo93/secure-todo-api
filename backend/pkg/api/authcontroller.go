package api

import (
	"github.com/labstack/echo/v4"
	"github.com/timoyo93/auth-backend/pkg/handlers"
	mw "github.com/timoyo93/auth-backend/pkg/middleware"
)

func RegisterEndpoints(e *echo.Echo, prefix string) {
	authGroup := e.Group(prefix)
	authGroup.GET("", mw.CheckCookie(handlers.TriggerAuth))
	authGroup.POST("/register", RegisterUser)
	authGroup.POST("/login", authhandler.LoginUser)
	authGroup.POST("/logout", authhandler.LogOutUser)
}
