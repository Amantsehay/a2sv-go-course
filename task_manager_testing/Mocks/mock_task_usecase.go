package Mocks

import (
	"github.com/stretchr/testify/mock"
	"task_manager_clean_architecture/Domain"
	"github.com/stretchr/testify/assert"
	"testing"
)

type MockTaskUsecase struct{
	mock.Mock
}
func (m *MockTaskUsecase) CreateTask(task *Domain.Task) error{
	args := m.Called(task)
	return args.Error(0)
}

func (m *MockTaskUsecase) GetTaskByID(taskID string) (*Domain.Task, error) {
	args := m.Called(taskID)
	return args.Get(0).(*Domain.Task), args.Error(1)
}


func (m *MockTaskUsecase) UpdateTask(taskID string, updatedTask Domain.Task) error {
	args := m.Called(taskID, updatedTask)
	return args.Error(0)
}

func (m *MockTaskUsecase) DeleteTask(taskID string) error {
	args := m.Called(taskID)
	return args.Error(0)
}




