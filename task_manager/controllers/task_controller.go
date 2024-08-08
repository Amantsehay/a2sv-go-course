package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"task_manager/data"
	"task_manager/models"
)

func GetTasks(ctx *gin.Context) {
	tasks := data.GetTasks()
	ctx.JSON(http.StatusOK, tasks)
}

func GetTasksById(ctx *gin.Context) {
	id := ctx.Param("id")
	task, err := data.GetTasksById(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, task)
}

func CreateTask(ctx *gin.Context) {
	var newTask models.Task

	if err := ctx.ShouldBindJSON(&newTask); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	if err := data.CreateTask(&newTask); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"message": "Task created successfully"})
}


func UpdateTask(ctx *gin.Context) {
	id := ctx.Param("id")
	var updatedTask models.Task

	if err := ctx.ShouldBindJSON(&updatedTask); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	err := data.UpdateTask(id, updatedTask)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Task updated successfully"})
}


func DeleteTaskById(ctx *gin.Context) {
	id := ctx.Param("id")
	err := data.DeleteTask(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})
}
