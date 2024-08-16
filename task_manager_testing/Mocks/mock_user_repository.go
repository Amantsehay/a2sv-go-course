package Mocks

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"task_manager_clean_architecture/Domain"
	"task_manager_clean_architecture/Infrastructure"
	"task_manager_clean_architecture/Repositories"
)


type MockCollection struct {
	mock.Mock
}

func (m *MockCollection) InsertOne(ctx context.Context, document interface{}) (*mongo.InsertOneResult, error) {
	args := m.Called(ctx, document)
	return args.Get(0).(*mongo.InsertOneResult), args.Error(1)
}

func (m *MockCollection) FindOne(ctx context.Context, filter interface{}) *mongo.SingleResult {
	args := m.Called(ctx, filter)
	result := &mongo.SingleResult{}
	return result
}

func (m *MockCollection) UpdateOne(ctx context.Context, filter interface{}, update interface{}) (*mongo.UpdateResult, error) {
	args := m.Called(ctx, filter, update)
	return args.Get(0).(*mongo.UpdateResult), args.Error(1)
}

func TestCreateUser_Success(t *testing.T) {
	mockCollection := new(MockCollection)
	mockInsertResult := &mongo.InsertOneResult{InsertedID: primitive.NewObjectID()}
	repo := &Repositories.MongoUserRepository{collection: mockCollection}

	hashedPassword := Infrastructure.HashPassword("testpassword")

	mockCollection.On("InsertOne", mock.Anything, mock.Anything).Return(mockInsertResult, nil)

	user, err := repo.CreateUser("testuser", "testpassword", "user")

	assert.NoError(t, err)
	assert.Equal(t, "testuser", user.Username)
	assert.Equal(t, "user", user.Role)
	mockCollection.AssertExpectations(t)
}

func TestAuthenticateUser_Success(t *testing.T) {
	mockCollection := new(MockCollection)
	repo := &Repositories.MongoUserRepository{collection: mockCollection}

	mockUser := Domain.User{
		ID:       primitive.NewObjectID(),
		Username: "testuser",
		Password: Infrastructure.HashPassword("testpassword"),
		Role:     "user",
	}

	mockCollection.On("FindOne", mock.Anything, mock.Anything).Return(mockUser)

	user, err := repo.AuthenticateUser("testuser", "testpassword")

	assert.NoError(t, err)
	assert.Equal(t, "testuser", user.Username)
	mockCollection.AssertExpectations(t)
}

func TestPromoteUser_Success(t *testing.T) {
	mockCollection := new(MockCollection)
	mockUpdateResult := &mongo.UpdateResult{MatchedCount: 1}
	repo := &Repositories.MongoUserRepository{collection: mockCollection}

	userID := primitive.NewObjectID().Hex()

	mockCollection.On("UpdateOne", mock.Anything, mock.Anything, mock.Anything).Return(mockUpdateResult, nil)

	err := repo.PromoteUser(userID)

	assert.NoError(t, err)
	mockCollection.AssertExpectations(t)
}
