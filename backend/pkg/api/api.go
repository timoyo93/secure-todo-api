package api

import (
	"fmt"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	"github.com/timoyo93/auth-backend/pkg/interfaces"
	"net/http"
)

const (
	CookieName = "JSESSIONID"
	CookiePath = "/"
)

type API struct {
	srv         *echo.Echo
	authService interfaces.AuthService
	todoService interfaces.TodoService
}

func New[A interfaces.AuthService, T interfaces.TodoService](authSvc A, todoSvc T) API {
	e := echo.New()
	return API{
		srv:         e,
		authService: authSvc,
		todoService: todoSvc,
	}.RegisterMiddleware()
}

func (a API) StartServer(port int) {
	portString := fmt.Sprintf(":%d", port)
	a.srv.Logger.Fatal(a.srv.Start(portString))
}

func (a API) RegisterTodoEndpoints(prefix string) API {
	apiGroup := a.srv.Group(prefix)
	apiGroup.Use(a.CheckCookie)
	apiGroup.POST("/todo", a.AddTodo)
	apiGroup.GET("/todos", a.GetAllTodos)
	apiGroup.PUT("/todo", a.UpdateTodo)
	apiGroup.GET("/todo/:id", a.GetTodo)
	apiGroup.DELETE("/todo/:id", a.RemoveTodo)

	return a
}

func (a API) RegisterAuthEndpoints(prefix string) API {
	apiGroup := a.srv.Group(prefix)
	apiGroup.GET("", a.CheckCookie(a.TriggerAuth))
	apiGroup.POST("/register", a.RegisterUser)
	apiGroup.POST("/login", a.LoginUser)
	apiGroup.POST("/logout", a.LogoutUser)

	return a
}

func (a API) EnableSwaggerAndHealthcheck() {
	a.srv.GET("/swagger/*", echoSwagger.WrapHandler)
	a.srv.GET("/healthcheck", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "healthy")
	})
}
