package controllers

import (
    "net/http"
    "time"
    "github.com/golang-jwt/jwt/v4"
    "github.com/gin-gonic/gin"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "task_manager_clean_architecture/Domain"
    "task_manager_clean_architecture/Usecases"
)

type Controller struct {
    userUsecases *Usecases.UserUsecases
    taskUsecases *Usecases.TaskUsecase
}

// Register creates a new user
func (c *Controller) Register(ctx *gin.Context) {
    var user Domain.User
    if err := ctx.ShouldBindJSON(&user); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    createdUser, err := c.userUsecases.CreateUser(user.Username, user.Password, user.Role)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusCreated, createdUser)
}

// Login authenticates a user and returns a JWT token
func (c *Controller) Login(ctx *gin.Context) {
    var loginData struct {
        Username string `json:"username"`
        Password string `json:"password"`
    }

    if err := ctx.ShouldBindJSON(&loginData); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    user, err := c.userUsecases.AuthenticateUser(loginData.Username, loginData.Password)
    if err != nil {
        ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
        return
    }

    claims := jwt.MapClaims{
        "userId":   user.ID.Hex(),
        "userName": user.Username,
        "role":     user.Role,
        "exp":      time.Now().Add(time.Hour * 72).Unix(),
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    tokenString, err := token.SignedString([]byte("jwt_secret_key0"))
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create token"})
        return
    }

    ctx.JSON(http.StatusOK, gin.H{"token": tokenString})
}

// PromoteUser promotes a user to admin
func (c *Controller) PromoteUser(ctx *gin.Context) {
    var promotionData struct {
        UserID string `json:"user_id"`
    }

    if err := ctx.ShouldBindJSON(&promotionData); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    userID, err := primitive.ObjectIDFromHex(promotionData.UserID)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
        return
    }

    err = c.userUsecases.PromoteUser(userID.Hex())
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, gin.H{"message": "User promoted to admin"})
}

// GetTasks retrieves all tasks
func (c *Controller) GetTasks(ctx *gin.Context) {
    tasks, err := c.taskUsecases.GetTasks()
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, tasks)
}

// GetTaskByID retrieves a task by ID
func (c *Controller) GetTaskByID(ctx *gin.Context) {
    id := ctx.Param("id")
    taskID, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
        return
    }

    task, err := c.taskUsecases.GetTaskByID(taskID.Hex())
    if err != nil {
        ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, task)
}

// CreateTask creates a new task
func (c *Controller) CreateTask(ctx *gin.Context) {
    var newTask Domain.Task

    if err := ctx.ShouldBindJSON(&newTask); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
        return
    }

    newTask.ID = primitive.NewObjectID().Hex()

    err := c.taskUsecases.CreateTask(&newTask)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusCreated, gin.H{"message": "Task created successfully"})
}

// UpdateTask updates an existing task
func (c *Controller) UpdateTask(ctx *gin.Context) {
    id := ctx.Param("id")
    taskID, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
        return
    }

    var updatedTask Domain.Task
    if err := ctx.ShouldBindJSON(&updatedTask); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
        return
    }

    err = c.taskUsecases.UpdateTask(taskID.Hex(), updatedTask)
    if err != nil {
        ctx.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
        return
    }

    ctx.JSON(http.StatusOK, gin.H{"message": "Task updated successfully"})
}

// DeleteTaskByID deletes a task by ID
func (c *Controller) DeleteTaskByID(ctx *gin.Context) {
    id := ctx.Param("id")
    taskID, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
        return
    }

    err = c.taskUsecases.DeleteTask(taskID.Hex())
    if err != nil {
        ctx.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
        return
    }

    ctx.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})
}
