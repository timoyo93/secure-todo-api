package todocontroller

import (
	"github.com/labstack/echo/v4"
	"github.com/timoyo93/auth-backend/pkg/handlers/todohandler"
	mw "github.com/timoyo93/auth-backend/pkg/middleware"
)

func RegisterEndpoints(e *echo.Echo, prefix string) {
	todoGroup := e.Group(prefix)
	todoGroup.Use(mw.CheckCookie)
	todoGroup.POST("/todo", todohandler.AddTodo)
	todoGroup.GET("/todos", todohandler.GetAllTodos)
	todoGroup.PUT("/todo/:id", todohandler.UpdateTodo)
	todoGroup.GET("/todo/:id", todohandler.GetTodo)
	todoGroup.DELETE("/todo/:id", todohandler.RemoveTodo)
}
