package main

import (
	"github.com/timoyo93/auth-backend/pkg/api"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
	"github.com/timoyo93/auth-backend/docs"
	"github.com/timoyo93/auth-backend/pkg/db"
)

var loggerConfig = middleware.LoggerConfig{
	Format: "${status} - ${method} ${uri}\t ->\t took ${latency_human}\n",
}

func registerMiddleware(e *echo.Echo) {
	e.Use(middleware.LoggerWithConfig(loggerConfig))
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowCredentials: true,
	}))
}

func main() {
	docs.SwaggerInfo.Title = "Secure TODO API with Cookie based authentication"
	db.InitDbConnection()
	e := echo.New()
	registerMiddleware(e)
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.GET("/healthcheck", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "healthy")
	})
	api.RegisterEndpoints(e, "/auth")
	api.RegisterEndpoints(e, "/api/v1")
	e.Logger.Fatal(e.Start(":8080"))
}
