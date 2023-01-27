package api

import (
	"github.com/labstack/echo/v4"
	"github.com/timoyo93/auth-backend/pkg/models"
	"net/http"
)

// AddTodo godoc
// @Summary Add a Todo
// @Tags Todo
// @Accept json
// @Produce json
// @Param todo body models.Todo true "Add Todo"
// @Success 201 {object} models.Todo
// @Failure 400 {string} string
// @Router /api/v1/todo [post]
func (a API) AddTodo(c echo.Context) error {
	todo := new(models.Todo)
	cookie, _ := c.Cookie("JSESSIONID")
	user, err := a.authService.GetUserByAccessToken(cookie.Value)
	if err != nil {
		return err
	}
	if err := c.Bind(&todo); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	if ok, err, id := a.todoService.CreateTodo(todo, user.ID.Hex()); err != nil || !ok {
		return err
	} else {
		todo.ID = id
	}
	return c.JSON(http.StatusCreated, todo)
}

// UpdateTodo godoc
// @Summary Update a Todo
// @Tags Todo
// @Accept json
// @Produce json
// @Param todo body models.Todo true "Update Todo"
// @Success 200 {object} models.Todo
// @Failure 400 {string} string
// @Failure 404 {string} string
// @Router /api/v1/todo [put]
func (a API) UpdateTodo(c echo.Context) error {
	var todo *models.Todo
	if err := c.Bind(&todo); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	cookie, _ := c.Cookie("JSESSIONID")
	user, err := a.authService.GetUserByAccessToken(cookie.Value)
	if err != nil {
		return err
	}
	ok, err := a.todoService.UpdateTodo(todo, user.ID.Hex())
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	if !ok {
		return c.JSON(http.StatusNotFound, "No todo found for given id")
	}
	return c.JSON(http.StatusOK, todo)
}

// RemoveTodo godoc
// @Summary Remove a Todo
// @Tags Todo
// @Accept json
// @Produce json
// @Param id path string true "Todo ID"
// @Success 204
// @Failure 400 {string} string
// @Failure 404 {string} string
// @Router /api/v1/todo/{id} [delete]
func (a API) RemoveTodo(c echo.Context) error {
	id := c.Param("id")
	if len(id) == 0 {
		return c.JSON(http.StatusBadRequest, "No ID for todo provided")
	}
	cookie, _ := c.Cookie("JSESSIONID")
	user, err := a.authService.GetUserByAccessToken(cookie.Value)
	if err != nil {
		return err
	}
	ok, err := a.todoService.DeleteTodo(id, user.ID.Hex())
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	if !ok {
		return c.JSON(http.StatusNotFound, "No Todo found for given id")
	}
	return c.JSON(http.StatusNoContent, nil)
}

// GetTodo godoc
// @Summary Gets Todo by ID
// @Tags Todo
// @Accept json
// @Produce json
// @Param id path string true "Todo ID"
// @Success 200 {array} models.Todo
// @Failure 400 {string} string
// @Failure 401 {string} string
// @Failure 404 {string} string
// @Router /api/v1/todo/{id} [get]
func (a API) GetTodo(c echo.Context) error {
	cookie, _ := c.Cookie("JSESSIONID")
	user, err := a.authService.GetUserByAccessToken(cookie.Value)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, "")
	}
	id := c.Param("id")
	if len(id) == 0 {
		return c.JSON(http.StatusBadRequest, "No ID for todo provided")
	}
	todo, err := a.todoService.GetTodoById(user.ID.Hex(), id)
	if err != nil {
		return c.JSON(http.StatusNotFound, "No Todo found for given ID")
	}
	return c.JSON(http.StatusOK, &todo)
}

// GetAllTodos godoc
// @Summary Get all Todos
// @Tags Todo
// @Accept json
// @Produce json
// @Success 200 {array} models.Todo
// @Failure 400 {string} string
// @Router /api/v1/todos [get]
func (a API) GetAllTodos(c echo.Context) error {
	cookie, _ := c.Cookie("JSESSIONID")
	user, err := a.authService.GetUserByAccessToken(cookie.Value)
	if err != nil {
		return err
	}
	todos, err := a.todoService.GetTodos(user.ID.Hex())
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, &todos)
}
