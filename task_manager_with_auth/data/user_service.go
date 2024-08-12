package data

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"task_manager_with_auth/models"
	"golang.org/x/crypto/bcrypt"
	"log"
)

var userCollection *mongo.Collection
func hashPassword(password string) string {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}
	return string(hashedPassword)
}
func checkPasswordHash(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
func init() {
	userCollection = GetDB().Collection("users")
}

func CreateUser(username, password, role string) (models.User, error) {
	hashedPassword := hashPassword(password)
	user := models.User{
		ID:       primitive.NewObjectID(),
		Username: username,
		Password: hashedPassword,
		Role:     role,
	}

	_, err := userCollection.InsertOne(context.TODO(), user)
	if err != nil {
		return models.User{}, err
	}
	user.Password = ""
	return user, nil
}

func AuthenticateUser(username, password string) (models.User, error) {
	var user models.User
	err := userCollection.FindOne(context.TODO(), bson.M{"username": username}).Decode(&user)
	if err != nil {
		return models.User{}, errors.New("invalid credentials")
	}
	if !checkPasswordHash(password, user.Password) {
		return models.User{}, errors.New("invalid credentials")
	}
	user.Password = ""
	return user, nil
}

func PromoteUser(userID string) error {
	id, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return err
	}
	_, err = userCollection.UpdateOne(
		context.TODO(),
		bson.M{"_id": id},
		bson.M{"$set": bson.M{"role": "admin"}},
	)
	if err != nil {
		return err
	}

	return nil
}
