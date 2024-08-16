package usecase_test

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"task_manager_clean_architecture/Domain"
    "task_manager_clean_architecture/Mocks"
)

func TestCreateTask_Success(t *testing.T){
	mockUsecase := new(Mocks.MockTaskUsecase)
	task := &Domain.Task{
		ID:          "1",
		Title:       "Sample Task",
		Description: "This is a sample task",
		DueDate:     time.Now().Add(24 * time.Hour),
		UserID:      primitive.NewObjectID(),
	}

	mockUsecase.On("CreateTask", task).Return(nil)
	err := mockUsecase.CreateTask(task)
	assert.NoError(t, err)
	mockUsecase.AssertExpectations(t)
}


