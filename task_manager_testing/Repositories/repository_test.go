package Repositories_test

import (
	"context"
	"testing"
	"task_manager_clean_architecture/Domain"
	"task_manager_clean_architecture/Repositories"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)


type MockCollection struct {
	mock.Mock
}

func (m *MockCollection) Find(ctx context.Context, filter interface{}) (*mongo.Cursor, error) {
	args := m.Called(ctx, filter)
	return args.Get(0).(*mongo.Cursor), args.Error(1)
}

func (m *MockCollection) FindOne(ctx context.Context, filter interface{}) *mongo.SingleResult {
	args := m.Called(ctx, filter)
	return &mongo.SingleResult{}
}

func (m *MockCollection) InsertOne(ctx context.Context, document interface{}) (*mongo.InsertOneResult, error) {
	args := m.Called(ctx, document)
	return args.Get(0).(*mongo.InsertOneResult), args.Error(1)
}

func (m *MockCollection) UpdateOne(ctx context.Context, filter interface{}, update interface{}) (*mongo.UpdateResult, error) {
	args := m.Called(ctx, filter, update)
	return args.Get(0).(*mongo.UpdateResult), args.Error(1)
}

func (m *MockCollection) DeleteOne(ctx context.Context, filter interface{}) (*mongo.DeleteResult, error) {
	args := m.Called(ctx, filter)
	return args.Get(0).(*mongo.DeleteResult), args.Error(1)
}

func TestGetTasks_Success(t *testing.T) {
	mockCollection := new(MockCollection)
	repo := &Repositories.MongoTaskRepository{collection: mockCollection}

	mockCursor := &mongo.Cursor{} // Placeholder, needs proper handling in real tests
	mockCollection.On("Find", mock.Anything, mock.Anything).Return(mockCursor, nil)

	tasks, err := repo.GetTasks(context.Background())
	assert.NoError(t, err)
	assert.NotNil(t, tasks)
	mockCollection.AssertExpectations(t)
}

func TestCreateTask_Success(t *testing.T) {
	mockCollection := new(MockCollection)
	repo := &Repositories.MongoTaskRepository{collection: mockCollection}

	mockTask := &Domain.Task{
		ID:    primitive.NewObjectID().Hex(),
		Title: "Test Task",
	}

	mockInsertResult := &mongo.InsertOneResult{InsertedID: primitive.NewObjectID()}
	mockCollection.On("InsertOne", mock.Anything, mock.Anything).Return(mockInsertResult, nil)

	err := repo.CreateTask(context.Background(), mockTask)
	assert.NoError(t, err)
	mockCollection.AssertExpectations(t)
}

func TestUpdateTask_Success(t *testing.T) {
	mockCollection := new(MockCollection)
	repo := &Repositories.MongoTaskRepository{collection: mockCollection}

	mockUpdateResult := &mongo.UpdateResult{MatchedCount: 1}
	mockCollection.On("UpdateOne", mock.Anything, mock.Anything, mock.Anything).Return(mockUpdateResult, nil)

	err := repo.UpdateTask(context.Background(), "some-id", Domain.Task{Title: "Updated Task"})
	assert.NoError(t, err)
	mockCollection.AssertExpectations(t)
}

func TestDeleteTask_Success(t *testing.T) {
	mockCollection := new(MockCollection)
	repo := &Repositories.MongoTaskRepository{collection: mockCollection}

	mockDeleteResult := &mongo.DeleteResult{DeletedCount: 1}
	mockCollection.On("DeleteOne", mock.Anything, mock.Anything).Return(mockDeleteResult, nil)

	err := repo.DeleteTask(context.Background(), "some-id")
	assert.NoError(t, err)
	mockCollection.AssertExpectations(t)
}
