package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type TodoDb struct {
	ID        *primitive.ObjectID `bson:"_id,omitempty"`
	Name      string              `bson:"name"`
	Completed bool                `bson:"completed"`
	UserId    string              `bson:"user_id"`
}

type Todo struct {
	ID        string `json:"id,omitempty"`
	Name      string `json:"name"`
	Completed bool   `json:"completed"`
}
