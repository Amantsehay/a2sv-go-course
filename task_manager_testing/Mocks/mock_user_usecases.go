package Mocks

import (
	"github.com/stretchr/testify/mock"
	"task_manager_clean_architecture/Domain"
)

type MockUserUsecases struct {
	mock.Mock
}

func (m *MockUserUsecases) CreateUser(username, password, role string) (*Domain.User, error) {
	args := m.Called(username, password, role)
	return args.Get(0).(*Domain.User), args.Error(1)
}

func (m *MockUserUsecases) AuthenticateUser(username, password string) (*Domain.User, error) {
	args := m.Called(username, password)
	return args.Get(0).(*Domain.User), args.Error(1)
}

func (m *MockUserUsecases) PromoteUser(userID string) error {
	args := m.Called(userID)
	return args.Error(0)
}
