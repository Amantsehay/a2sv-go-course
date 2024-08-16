package Repositories

import (
	"context"
	"testing"
	"errors"
	"task_manager_clean_architecture/Domain"
	"task_manager_clean_architecture/Repositories"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoUserRepository struct {
	collection *mongo.Collection
}


func NewMongoUserRepository() (*MongoUserRepository, error){
	_, usersCollection := Infrastructure.CollectionNames()
	if usersCollection == nil {
        return nil, errors.New("failed to retrieve the users collection")
    }
	return &MongoUserRepository{collection: usersCollection}, nil
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



type MockCollection struct {
	mock.Mock
}

func (m *MockCollection) InsertOne(ctx context.Context, document interface{}) (*mongo.InsertOneResult, error) {
	args := m.Called(ctx, document)
	return args.Get(0).(*mongo.InsertOneResult), args.Error(1)
}

func (m *MockCollection) FindOne(ctx context.Context, filter interface{}) *mongo.SingleResult {
	args := m.Called(ctx, filter)
	return &mongo.SingleResult{}
}

func (m *MockCollection) UpdateOne(ctx context.Context, filter interface{}, update interface{}) (*mongo.UpdateResult, error) {
	args := m.Called(ctx, filter, update)
	return args.Get(0).(*mongo.UpdateResult), args.Error(1)
}

func TestCreateUser_Success(t *testing.T) {
	mockCollection := new(MockCollection)
	repo := &Repositories.MongoUserRepository{collection: mockCollection}

	mockUser := Domain.User{
		ID:       primitive.NewObjectID(),
		Username: "testuser",
		Password: "hashedpassword",
		Role:     "user",
	}
	mockInsertResult := &mongo.InsertOneResult{InsertedID: mockUser.ID}
	mockCollection.On("InsertOne", mock.Anything, mockUser).Return(mockInsertResult, nil)

	user, err := repo.CreateUser("testuser", "password", "user")

	assert.NoError(t, err)
	assert.Equal(t, "testuser", user.Username)
	assert.Equal(t, "user", user.Role)
	assert.Empty(t, user.Password) // Password should be cleared
	mockCollection.AssertExpectations(t)
}

func TestAuthenticateUser_Success(t *testing.T) {
	mockCollection := new(MockCollection)
	repo := &Repositories.MongoUserRepository{collection: mockCollection}

	mockUser := Domain.User{
		ID:       primitive.NewObjectID(),
		Username: "testuser",
		Password: "hashedpassword",
		Role:     "user",
	}
	mockCollection.On("FindOne", mock.Anything, bson.M{"username": "testuser"}).Return(mockUser)

	// Mock the password check function
	originalCheckPasswordHash := Infrastructure.CheckPasswordHash
	Infrastructure.CheckPasswordHash = func(password, hashedPassword string) bool {
		return password == "password" && hashedPassword == "hashedpassword"
	}
	defer func() { Infrastructure.CheckPasswordHash = originalCheckPasswordHash }() // Restore original function

	user, err := repo.AuthenticateUser("testuser", "password")

	assert.NoError(t, err)
	assert.Equal(t, "testuser", user.Username)
	assert.Equal(t, "user", user.Role)
	assert.Empty(t, user.Password) // Password should be cleared
	mockCollection.AssertExpectations(t)
}

func TestPromoteUser_Success(t *testing.T) {
	mockCollection := new(MockCollection)
	repo := &Repositories.MongoUserRepository{collection: mockCollection}

	userID := primitive.NewObjectID().Hex()
	mockUpdateResult := &mongo.UpdateResult{MatchedCount: 1}
	mockCollection.On("UpdateOne", mock.Anything, bson.M{"_id": primitive.ObjectIDFromHex(userID)}, bson.M{"$set": bson.M{"role": "admin"}}).Return(mockUpdateResult, nil)

	err := repo.PromoteUser(userID)

	assert.NoError(t, err)
	mockCollection.AssertExpectations(t)
}

func TestPromoteUser_Failure(t *testing.T) {
	mockCollection := new(MockCollection)
	repo := &Repositories.MongoUserRepository{collection: mockCollection}

	userID := primitive.NewObjectID().Hex()
	mockCollection.On("UpdateOne", mock.Anything, bson.M{"_id": primitive.ObjectIDFromHex(userID)}, bson.M{"$set": bson.M{"role": "admin"}}).Return(&mongo.UpdateResult{MatchedCount: 0}, nil)

	err := repo.PromoteUser(userID)

	assert.Error(t, err)
	mockCollection.AssertExpectations(t)
}