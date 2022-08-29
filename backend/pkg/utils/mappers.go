package utils

import "github.com/timoyo93/auth-backend/pkg/models"

func MapTodoToTodoDb(userId string, todo models.Todo) models.TodoDb {
	return models.TodoDb{
		UserId:    userId,
		Name:      todo.Name,
		Completed: todo.Completed,
	}
}
