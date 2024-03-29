package db

import (
	"errors"
	"github.com/timoyo93/auth-backend/pkg/models"
	"go.mongodb.org/mongo-driver/bson"
)

func (r *Repository) GetUser(username string) (*models.UserDB, error) {
	var user *models.UserDB
	filter := bson.M{Username: username}
	if err := r.users.FindOne(r.ctx, filter).Decode(&user); err != nil {
		return nil, err
	}
	return user, nil
}

func (r *Repository) GetUserByAccessToken(token string) (*models.UserDB, error) {
	var user *models.UserDB
	filter := bson.M{AccessToken: token}
	err := r.users.FindOne(r.ctx, filter).Decode(&user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *Repository) AddUser(user *models.UserDB) error {
	if _, err := r.users.InsertOne(r.ctx, user); err != nil {
		return err
	}
	return nil
}

func (r *Repository) CheckForAccessToken(token string) error {
	var user *models.UserDB
	filter := bson.M{AccessToken: token}
	err := r.users.FindOne(r.ctx, filter).Decode(&user)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) SetAccessTokenForUser(username, token string) error {
	filter := bson.M{Username: username}
	update := bson.M{"$set": bson.M{AccessToken: token}}
	result, err := r.users.UpdateOne(r.ctx, filter, update)
	if err != nil {
		return err
	}
	if result.ModifiedCount == 0 {
		return errors.New("no token was set for given user")
	}
	return nil
}

func (r *Repository) RemoveAccessTokenForUser(token string) error {
	filter := bson.M{AccessToken: token}
	update := bson.M{"$unset": bson.M{AccessToken: ""}}
	result, err := r.users.UpdateOne(r.ctx, filter, update)
	if err != nil {
		return err
	}
	if result.ModifiedCount == 0 {
		return errors.New("no access token to remove")
	}
	return nil
}
