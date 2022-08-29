package db

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
)

var Db *mongo.Database
var UsersCollection *mongo.Collection
var Ctx context.Context
var TodoCollection *mongo.Collection

const (
	Username      string = "username"
	AccessToken          = "access_token"
	UserId               = "user_id"
	TodoId               = "_id"
	TodoName             = "name"
	TodoCompleted        = "completed"
)

func InitDbConnection() {
	credentials := options.Credential{
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
	}
	connString := os.Getenv("DB_CONN_STRING")
	clientOpts := options.Client().ApplyURI(connString).SetAuth(credentials)
	client, err := mongo.Connect(context.TODO(), clientOpts)
	if err != nil {
		fmt.Printf("Can not connect to database with connection string: %v\n", connString)
	} else {
		Db = client.Database("app")
		fmt.Println("Connected to database!")
		UsersCollection = Db.Collection("users")
		TodoCollection = Db.Collection("todos")
		Ctx = context.TODO()
	}
}
