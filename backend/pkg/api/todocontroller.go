package api

import (
	"github.com/labstack/echo/v4"
	"github.com/timoyo93/auth-backend/pkg/handlers"
	mw "github.com/timoyo93/auth-backend/pkg/middleware"
)

func RegisterEndpoints(e *echo.Echo, prefix string) {
	todoGroup := e.Group(prefix)
	todoGroup.Use(mw.CheckCookie)
	todoGroup.POST("/todo", handlers.AddTodo)
	todoGroup.GET("/todos", handlers.GetAllTodos)
	todoGroup.PUT("/todo/:id", handlers.UpdateTodo)
	todoGroup.GET("/todo/:id", handlers.GetTodo)
	todoGroup.DELETE("/todo/:id", handlers.RemoveTodo)
}
