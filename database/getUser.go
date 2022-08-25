package database

import (
	"context"
	"errors"
	"firestore/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetUserFromEmail(email string) (models.User, error) {
	var user models.User
	err := UsersColl.FindOne(context.TODO(), bson.M{"email": email}).Decode(&user)
	if err != nil {
		return user, errors.New("we have problem getting the user")
	}
	return user, nil
}
func GetUserFromId(id string) (models.User, error) {
	var user models.User
	idobj, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return user, errors.New("we have problem getting the user")
	}

	err = UsersColl.FindOne(context.TODO(), bson.M{"_id": idobj}).Decode(&user)
	if err != nil {
		return user, errors.New("we have problem getting the user")
	}
	return user, nil
}
