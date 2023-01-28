package interfaces

import "github.com/timoyo93/auth-backend/pkg/models"

type TodoService interface {
	GetTodos(userID string) (*[]models.Todo, error)
	GetTodoById(userID, todoId string) (*models.Todo, error)
	UpdateTodo(todo *models.Todo, userID string) (bool, error)
	CreateTodo(todo *models.Todo, userID string) (bool, error, string)
	DeleteTodo(todoID, userID string) (bool, error)
}

type AuthService interface {
	GetUser(username string) (*models.UserDB, error)
	GetUserByAccessToken(token string) (*models.UserDB, error)
	CreateUser(auth *models.AuthRequest) error
	CheckAccessToken(token string) error
	SetAccessToken(username string) (string, error)
	RemoveAccessToken(token string) error
}
