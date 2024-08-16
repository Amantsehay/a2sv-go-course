package router

import (
	"github.com/gin-gonic/gin"

	"task_manager_clean_architecture/Infrastructure"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/tasks",GetTasks)
	r.POST("/tasks", AuthMiddleware(), CreateTask) 
	r.PUT("/tasks/:id", AuthMiddleware(), UpdateTask)
	r.GET("/tasks/:id", GetTaskById)
	r.DELETE("/tasks/:id", AuthMiddleware(), AdminMiddleware(), DeleteTaskById)

	r.POST("/register", Register)
	r.POST("/login", Login)
	r.POST("/promote", AuthMiddleware(), AdminMiddleware(), PromoteUser)

	return r

}
