package db

import (
	"github.com/timoyo93/auth-backend/pkg/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r *Repository) GetAllTodos(userId string) ([]models.Todo, error) {
	filter := bson.M{UserId: userId}
	var todos []models.Todo
	c, err := r.todos.Find(r.ctx, filter)
	if err != nil {
		return nil, err
	}
	for c.Next(r.ctx) {
		var todo *models.TodoDb
		if err := c.Decode(&todo); err != nil {
			return nil, err
		}
		t := models.Todo{
			ID:        todo.ID.Hex(),
			Name:      todo.Name,
			Completed: todo.Completed,
		}
		todos = append(todos, t)
	}

	return todos, nil
}

func (r *Repository) GetTodoById(userId, todoId string) (*models.TodoDb, error) {
	objectId, err := primitive.ObjectIDFromHex(todoId)
	if err != nil {
		return nil, err
	}
	filter := bson.M{TodoId: objectId, UserId: userId}
	var todo *models.TodoDb
	if err := r.todos.FindOne(r.ctx, filter).Decode(&todo); err != nil {
		return nil, err
	}
	return todo, nil
}

func (r *Repository) UpdateTodo(todo *models.Todo, userId, todoId string) (bool, error) {
	objectId, err := primitive.ObjectIDFromHex(todoId)
	if err != nil {
		return false, err
	}
	filter := bson.M{TodoId: objectId, UserId: userId}
	update := bson.M{"$set": bson.M{TodoName: todo.Name, TodoCompleted: todo.Completed}}
	result, err := r.todos.UpdateOne(r.ctx, filter, update)
	if err != nil {
		return false, err
	}
	if result.ModifiedCount == 0 {
		return false, err
	}
	return true, nil
}

func (r *Repository) AddTodo(todo models.TodoDb) (bool, error, string) {
	result, err := r.todos.InsertOne(r.ctx, todo)
	if err != nil {
		return false, err, ""
	}
	return true, nil, result.InsertedID.(primitive.ObjectID).Hex()
}

func (r *Repository) DeleteTodo(todoId, userId string) (bool, error) {
	objectId, err := primitive.ObjectIDFromHex(todoId)
	if err != nil {
		return false, err
	}
	filter := bson.M{TodoId: objectId, UserId: userId}
	result, err := r.todos.DeleteOne(r.ctx, filter)
	if err != nil {
		return false, err
	}
	return result.DeletedCount != 0, nil
}
