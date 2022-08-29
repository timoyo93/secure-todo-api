package db

import (
	"github.com/timoyo93/auth-backend/pkg/models"
	"go.mongodb.org/mongo-driver/bson"
)

func GetUser(username string) (*models.UserDb, error) {
	var user *models.UserDb
	filter := bson.M{Username: username}
	if err := UsersCollection.FindOne(Ctx, filter).Decode(&user); err != nil {
		return nil, err
	}
	return user, nil
}

func GetUserByAccessToken(token string) (*models.UserDb, error) {
	var user *models.UserDb
	filter := bson.M{AccessToken: token}
	err := UsersCollection.FindOne(Ctx, filter).Decode(&user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func AddUser(user models.UserDb) error {
	if _, err := UsersCollection.InsertOne(Ctx, user); err != nil {
		return err
	}
	return nil
}

func CheckForAccessToken(token string) bool {
	var user *models.UserDb
	filter := bson.M{AccessToken: token}
	err := UsersCollection.FindOne(Ctx, filter).Decode(&user)
	if err != nil || user == nil {
		return false
	}
	return true
}

func SetAccessTokenForUser(user models.UserDb) (error, bool) {
	filter := bson.M{Username: user.Username}
	update := bson.M{"$set": bson.M{AccessToken: user.Token}}
	result, err := UsersCollection.UpdateOne(Ctx, filter, update)
	if err != nil {
		return err, false
	}
	if result.ModifiedCount == 0 {
		return nil, false
	}
	return nil, true
}

func RemoveAccessTokenForUser(token string) (error, bool) {
	filter := bson.M{AccessToken: token}
	update := bson.M{"$set": bson.M{AccessToken: ""}}
	result, err := UsersCollection.UpdateOne(Ctx, filter, update)
	if err != nil {
		return err, false
	}
	if result.ModifiedCount == 0 {
		return nil, false
	}
	return nil, true
}
