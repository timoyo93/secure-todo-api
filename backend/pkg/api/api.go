package api

import "github.com/labstack/echo/v4"

type API struct {
	srv *echo.Echo
	userService
}
