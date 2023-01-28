package db

import (
	"github.com/timoyo93/auth-backend/pkg/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r *Repository) GetAllTodos(userID string) (*[]models.Todo, error) {
	filter := bson.M{UserID: userID}
	var todos []models.Todo
	c, err := r.todos.Find(r.ctx, filter)
	if err != nil {
		return nil, err
	}
	for c.Next(r.ctx) {
		var todo *models.TodoDB
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

	return &todos, nil
}

func (r *Repository) GetTodoById(userID, todoID string) (*models.TodoDB, error) {
	objectID, err := primitive.ObjectIDFromHex(todoID)
	if err != nil {
		return nil, err
	}
	filter := bson.M{TodoID: objectID, UserID: userID}
	var todo *models.TodoDB
	if err := r.todos.FindOne(r.ctx, filter).Decode(&todo); err != nil {
		return nil, err
	}
	return todo, nil
}

func (r *Repository) UpdateTodo(todo *models.Todo, userID string) (bool, error) {
	objectID, err := primitive.ObjectIDFromHex(todo.ID)
	if err != nil {
		return false, err
	}
	filter := bson.M{TodoID: objectID, UserID: userID}
	update := bson.M{"$set": bson.M{TodoName: &todo.Name, TodoCompleted: &todo.Completed}}
	result, err := r.todos.UpdateOne(r.ctx, filter, update)
	if err != nil {
		return false, err
	}
	if result.ModifiedCount == 0 {
		return false, err
	}
	return true, nil
}

func (r *Repository) AddTodo(todo *models.TodoDB) (bool, error, string) {
	result, err := r.todos.InsertOne(r.ctx, &todo)
	if err != nil {
		return true, err, ""
	}
	return true, nil, result.InsertedID.(primitive.ObjectID).Hex()
}

func (r *Repository) DeleteTodo(todoID, userID string) (bool, error) {
	objectID, err := primitive.ObjectIDFromHex(todoID)
	if err != nil {
		return false, err
	}
	filter := bson.M{TodoID: objectID, UserID: userID}
	result, err := r.todos.DeleteOne(r.ctx, filter)
	if err != nil {
		return false, err
	}
	return result.DeletedCount != 0, nil
}
