package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type AuthRequest struct {
	Username string
	Password string
}

type UserDb struct {
	ID       *primitive.ObjectID `bson:"_id,omitempty"`
	Token    string              `bson:"access_token"`
	Username string              `bson:"username"`
	Hash     string              `bson:"hash"`
}
