package services

import (
	"github.com/timoyo93/auth-backend/pkg/db"
	"github.com/timoyo93/auth-backend/pkg/models"
)

type TodoRepository interface {
	GetAllTodos(string) ([]models.Todo, error)
	GetTodoById(string, string) (*models.TodoDb, error)
	UpdateTodo(*models.Todo, string, string) (bool, error)
	AddTodo(models.TodoDb) (bool, error, string)
	DeleteTodo(string, string) (bool, error)
}

type TodoService struct {
	repo TodoRepository
}

func RegisterTodo(repo db.Repository) TodoService {
	return TodoService{repo: &repo}
}

func (s TodoService) GetTodos(userId string) ([]models.Todo, error) {
	todos, err := s.repo.GetAllTodos(userId)
	if err != nil {
		return nil, err
	}
	return todos, nil
}

func (s TodoService) GetTodoById(userId, todoId string) (models.Todo, error) {
	dbTodo, err := s.repo.GetTodoById(userId, todoId)
	if err != nil {
		return models.Todo{}, err
	}

}