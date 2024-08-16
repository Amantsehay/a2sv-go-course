package Usecases_test

import (
	"context"
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"task_manager_clean_architecture/Domain"
	// "task_manager_clean_architecture/Repositories"
	"task_manager_clean_architecture/Usecases"
)

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) CreateUser(username, password, role string) (Domain.User, error) {
	args := m.Called(username, password, role)
	return args.Get(0).(Domain.User), args.Error(1)
}

func (m *MockUserRepository) AuthenticateUser(username, password string) (Domain.User, error) {
	args := m.Called(username, password)
	return args.Get(0).(Domain.User), args.Error(1)
}

func (m *MockUserRepository) PromoteUser(userID string) error {
	args := m.Called(userID)
	return args.Error(0)
}

func TestCreateUser_Success(t *testing.T) {
	mockRepo := new(MockUserRepository)
	userUsecases := Usecases.NewUserUsecases(mockRepo)

	mockUser := Domain.User{Username: "testuser", Role: "user"}
	mockRepo.On("CreateUser", "testuser", "testpassword", "user").Return(mockUser, nil)

	createdUser, err := userUsecases.CreateUser("testuser", "testpassword", "user")

	assert.NoError(t, err)
	assert.Equal(t, "testuser", createdUser.Username)
	assert.Equal(t, "user", createdUser.Role)
	mockRepo.AssertExpectations(t)
}

func TestAuthenticateUser_Success(t *testing.T) {
	mockRepo := new(MockUserRepository)
	userUsecases := Usecases.NewUserUsecases(mockRepo)

	mockUser := Domain.User{Username: "testuser", Password: "hashedpassword"}
	mockRepo.On("AuthenticateUser", "testuser", "testpassword").Return(mockUser, nil)

	authUser, err := userUsecases.AuthenticateUser("testuser", "testpassword")

	assert.NoError(t, err)
	assert.Equal(t, "testuser", authUser.Username)
	mockRepo.AssertExpectations(t)
}

func TestPromoteUser_Success(t *testing.T) {
	mockRepo := new(MockUserRepository)
	userUsecases := Usecases.NewUserUsecases(mockRepo)

	mockRepo.On("PromoteUser", "userID").Return(nil)

	err := userUsecases.PromoteUser("userID")

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}


type MockTaskRepository struct {
	mock.Mock
}

func (m *MockTaskRepository) GetTasks(ctx context.Context) ([]*Domain.Task, error) {
	args := m.Called(ctx)
	return args.Get(0).([]*Domain.Task), args.Error(1)
}

func (m *MockTaskRepository) GetTaskByID(ctx context.Context, id string) (*Domain.Task, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*Domain.Task), args.Error(1)
}

func (m *MockTaskRepository) CreateTask(ctx context.Context, task *Domain.Task) error {
	args := m.Called(ctx, task)
	return args.Error(0)
}

func (m *MockTaskRepository) UpdateTask(ctx context.Context, id string, updatedTask Domain.Task) error {
	args := m.Called(ctx, id, updatedTask)
	return args.Error(0)
}

func (m *MockTaskRepository) DeleteTask(ctx context.Context, id string) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

// Test for CreateTask
func TestCreateTask_Success(t *testing.T) {
	mockRepo := new(MockTaskRepository)
	taskUsecase := Usecases.NewTaskUsecase(mockRepo)

	task := &Domain.Task{
		ID:    "1",
		Title: "Test Task",
	}

	mockRepo.On("GetTaskByID", context.TODO(), task.ID).Return(nil, nil)
	mockRepo.On("CreateTask", context.TODO(), task).Return(nil)

	err := taskUsecase.CreateTask(task)
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestCreateTask_AlreadyExists(t *testing.T) {
	mockRepo := new(MockTaskRepository)
	taskUsecase := Usecases.NewTaskUsecase(mockRepo)

	existingTask := &Domain.Task{
		ID:    "1",
		Title: "Existing Task",
	}

	mockRepo.On("GetTaskByID", context.TODO(), existingTask.ID).Return(existingTask, nil)

	err := taskUsecase.CreateTask(existingTask)
	assert.Error(t, err)
	assert.Equal(t, "task already exists", err.Error())
	mockRepo.AssertExpectations(t)
}


func TestGetTaskByID_Success(t *testing.T) {
	mockRepo := new(MockTaskRepository)
	taskUsecase := Usecases.NewTaskUsecase(mockRepo)

	task := &Domain.Task{
		ID:    "1",
		Title: "Test Task",
	}

	mockRepo.On("GetTaskByID", context.TODO(), task.ID).Return(task, nil)

	result, err := taskUsecase.GetTaskByID(task.ID)
	assert.NoError(t, err)
	assert.Equal(t, task, result)
	mockRepo.AssertExpectations(t)
}

// Test for UpdateTask
func TestUpdateTask_Success(t *testing.T) {
	mockRepo := new(MockTaskRepository)
	taskUsecase := Usecases.NewTaskUsecase(mockRepo)

	task := &Domain.Task{
		ID:    "1",
		Title: "Test Task",
	}

	updatedTask := Domain.Task{
		Title: "Updated Task",
	}

	mockRepo.On("GetTaskByID", context.TODO(), task.ID).Return(task, nil)
	mockRepo.On("UpdateTask", context.TODO(), task.ID, updatedTask).Return(nil)

	err := taskUsecase.UpdateTask(task.ID, updatedTask)
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestDeleteTask_Success(t *testing.T) {
	mockRepo := new(MockTaskRepository)
	taskUsecase := Usecases.NewTaskUsecase(mockRepo)

	taskID := "1"

	mockRepo.On("GetTaskByID", context.TODO(), taskID).Return(&Domain.Task{}, nil)
	mockRepo.On("DeleteTask", context.TODO(), taskID).Return(nil)

	err := taskUsecase.DeleteTask(taskID)
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}