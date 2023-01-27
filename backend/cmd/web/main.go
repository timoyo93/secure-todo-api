package main

import (
	"github.com/timoyo93/auth-backend/docs"
	"github.com/timoyo93/auth-backend/pkg/api"
	"github.com/timoyo93/auth-backend/pkg/db"
	"github.com/timoyo93/auth-backend/pkg/services"
)

var (
	database    = db.New()
	authService = services.InitAuth(database)
	todoService = services.InitTodo(database)
)

func main() {
	docs.SwaggerInfo.Title = "Secure TODO API with Cookie based authentication"
	restService := api.New(authService, todoService)
	restService.RegisterMiddleware()
	restService.EnableSwaggerAndHealthcheck()
	restService.RegisterAuthEndpoints("/auth")
	restService.RegisterTodoEndpoints("/api/v1")
	restService.StartServer(8080)
}
