package services

import (
	"github.com/timoyo93/auth-backend/pkg/db"
	"github.com/timoyo93/auth-backend/pkg/models"
)

type TodoRepository interface {
	GetAllTodos(userID string) (*[]models.Todo, error)
	GetTodoById(userID, todoID string) (*models.TodoDB, error)
	UpdateTodo(todo *models.Todo, userID string) (bool, error)
	AddTodo(todo *models.TodoDB) (bool, error, string)
	DeleteTodo(todoID, userID string) (bool, error)
}

type TodoService struct {
	repo TodoRepository
}

func InitTodo(repo *db.Repository) TodoService {
	return TodoService{repo: repo}
}

func (s TodoService) GetTodos(userId string) (*[]models.Todo, error) {
	todos, err := s.repo.GetAllTodos(userId)
	if err != nil {
		return nil, err
	}
	return todos, nil
}

func (s TodoService) GetTodoById(userId, todoId string) (*models.Todo, error) {
	dbTodo, err := s.repo.GetTodoById(userId, todoId)
	if err != nil {
		return nil, err
	}
	return dbTodo.Map(), nil
}

func (s TodoService) UpdateTodo(todo *models.Todo, userId string) (bool, error) {
	ok, err := s.repo.UpdateTodo(todo, userId)
	return ok, err
}

func (s TodoService) CreateTodo(todo *models.Todo, userId string) (bool, error, string) {
	ok, err, insertedId := s.repo.AddTodo(todo.Map(userId))
	return ok, err, insertedId
}

func (s TodoService) DeleteTodo(todoId, userId string) (bool, error) {
	ok, err := s.repo.DeleteTodo(todoId, userId)
	return ok, err
}
