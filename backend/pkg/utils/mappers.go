package utils

import "github.com/timoyo93/auth-backend/pkg/models"

func (dbTodo models.TodoDb) Map() models.Todo {
	return models.Todo{
		ID:        dbTodo.ID.Hex(),
		Name:      dbTodo.Name,
		Completed: dbTodo.Completed,
	}
}
