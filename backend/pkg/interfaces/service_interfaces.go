package interfaces

import "github.com/timoyo93/auth-backend/pkg/models"

type TodoService interface {
	GetTodos(userId string) (*[]models.Todo, error)
	GetTodoById(userId, todoId string) (*models.Todo, error)
	UpdateTodo(todo *models.Todo, userId string) (bool, error)
	CreateTodo(todo *models.Todo, userId string) (bool, error, string)
	DeleteTodo(todoId, userId string) (bool, error)
}

type AuthService interface {
	GetUser(username string) (*models.UserDb, error)
	GetUserByAccessToken(token string) (*models.UserDb, error)
	CreateUser(auth *models.AuthRequest) error
	CheckAccessToken(token string) bool
	SetAccessToken(username string) (string, error)
	RemoveAccessToken(token string) (bool, error)
}
