package Repositories

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"task_manager_clean_architecture/Domain"
	"task_manager_clean_architecture/Infrastructure"
)

type MongoUserRepository struct {
	collection *mongo.Collection
}

func NewMongoUserRepository() *MongoUserRepository {
	_, usersCollection := Infrastructure.CollectionNames()
	return &MongoUserRepository{collection: usersCollection}
}

func (r *MongoUserRepository) CreateUser(username, password, role string) (Domain.User, error) {
	hashedPassword := Infrastructure.HashPassword(password)
	user := Domain.User{
		ID:       primitive.NewObjectID(),
		Username: username,
		Password: hashedPassword,
		Role:     role,
	}

	_, err := r.collection.InsertOne(context.Background(), user)
	if err != nil {
		return Domain.User{}, err
	}
	user.Password = "" 
	return user, nil
}

func (r *MongoUserRepository) AuthenticateUser(username, password string) (Domain.User, error) {
	var user Domain.User
	err := r.collection.FindOne(context.Background(), bson.M{"username": username}).Decode(&user)
	if err != nil {
		return Domain.User{}, errors.New("invalid credentials")
	}
	if !Infrastructure.CheckPasswordHash(password, user.Password) {
		return Domain.User{}, errors.New("invalid credentials")
	}
	user.Password = ""
	return user, nil
}

func (r *MongoUserRepository) PromoteUser(userID string) error {
	id, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return err
	}
	_, err = r.collection.UpdateOne(
		context.Background(),
		bson.M{"_id": id},
		bson.M{"$set": bson.M{"role": "admin"}},
	)
	if err != nil {
		return err
	}
	return nil
}
