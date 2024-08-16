package controllers_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"task_manager_clean_architecture/Domain"
	"task_manager_clean_architecture/Delivery/controllers"
	"task_manager_clean_architecture/mocks"
)

func TestRegister(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockUserUsecase := new(mocks.MockUserUsecases)

	controller := &controllers.Controller{
		UserUsecases: mockUserUsecase,
	}

	mockUser := &Domain.User{Username: "testuser", Role: "user"}
	mockUserUsecase.On("CreateUser", "testuser", "password", "user").Return(mockUser, nil)

	reqBody := `{"username":"testuser", "password":"password", "role":"user"}`
	req, _ := http.NewRequest(http.MethodPost, "/register", strings.NewReader(reqBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = req
	controller.Register(ctx)

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Contains(t, w.Body.String(), "testuser")

	mockUserUsecase.AssertExpectations(t)
}

func TestLogin(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockUserUsecase := new(mocks.MockUserUsecases)

	controller := &controllers.Controller{
		UserUsecases: mockUserUsecase,
	}

	mockUser := &Domain.User{ID: "1", Username: "testuser", Role: "user"}
	mockUserUsecase.On("AuthenticateUser", "testuser", "password").Return(mockUser, nil)

	reqBody := `{"username":"testuser", "password":"password"}`
	req, _ := http.NewRequest(http.MethodPost, "/login", strings.NewReader(reqBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = req
	controller.Login(ctx)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "token")

	mockUserUsecase.AssertExpectations(t)
}

func TestCreateTask(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockTaskUsecase := new(mocks.MockTaskUsecase)

	controller := &controllers.Controller{
		TaskUsecases: mockTaskUsecase,
	}

	mockTask := &Domain.Task{ID: "1", Title: "Test Task"}
	mockTaskUsecase.On("CreateTask", mock.AnythingOfType("*Domain.Task")).Return(nil)

	reqBody := `{"title":"Test Task"}`
	req, _ := http.NewRequest(http.MethodPost, "/tasks", strings.NewReader(reqBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = req
	controller.CreateTask(ctx)

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Contains(t, w.Body.String(), "Task created successfully")

	mockTaskUsecase.AssertExpectations(t)
}
