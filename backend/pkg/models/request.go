package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type AuthRequest struct {
	Username string
	Password string
}

type UserDB struct {
	ID       *primitive.ObjectID `bson:"_id,omitempty"`
	Token    string              `bson:"access_token"`
	Username string              `bson:"username"`
	Hash     []byte              `bson:"hash"`
}

func NewDBUser(username string, hash []byte) *UserDB {
	return &UserDB{
		Username: username,
		Hash:     hash,
	}
}
