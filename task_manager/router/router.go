package router

import (
    "github.com/gin-gonic/gin"
    "task_manager/controllers"
)

func SetupRouter() *gin.Engine {
    r := gin.Default()

    r.GET("/tasks", controllers.GetTasks)
    r.GET("/tasks/:id", controllers.GetTasksById)
    r.POST("/tasks", controllers.CreateTask)
    r.PUT("/tasks/:id", controllers.UpdateTask)
    r.DELETE("/tasks/:id", controllers.DeleteTaskById)

    return r
}
