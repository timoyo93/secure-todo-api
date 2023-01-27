package db

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
)

const (
	Username      string = "username"
	AccessToken          = "access_token"
	UserId               = "user_id"
	TodoId               = "_id"
	TodoName             = "name"
	TodoCompleted        = "completed"
)

type Repository struct {
	db    *mongo.Database
	ctx   context.Context
	users *mongo.Collection
	todos *mongo.Collection
}

func New() *Repository {
	credentials := options.Credential{
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
	}
	connString := os.Getenv("DB_CONN_STRING")
	clientOpts := options.Client().ApplyURI(connString).SetAuth(credentials)
	client, err := mongo.Connect(context.TODO(), clientOpts)
	if err != nil {
		fmt.Printf("Can not connect to database with connection string: %v\n", connString)
		panic("")
	} else {
		conn := client.Database("app")
		return &Repository{
			db:    conn,
			users: conn.Collection("users"),
			todos: conn.Collection("todos"),
			ctx:   context.TODO(),
		}
	}
}
