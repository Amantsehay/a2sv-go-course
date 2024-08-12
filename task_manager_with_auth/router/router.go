package router

import (
	"github.com/gin-gonic/gin"
	"task_manager_with_auth/controllers"
	"task_manager_with_auth/middleware"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Task routes
	r.GET("/tasks", controllers.GetTasks)
	r.POST("/tasks", middleware.AuthMiddleware(), controllers.CreateTask) 
	r.PUT("/tasks/:id", middleware.AuthMiddleware(), controllers.UpdateTask)
	r.GET("/tasks/:id", controllers.GetTaskById)
	r.DELETE("/tasks/:id", middleware.AuthMiddleware(), middleware.AdminMiddleware(), controllers.DeleteTaskById)


	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)
	r.POST("/promote", middleware.AuthMiddleware(), middleware.AdminMiddleware(), controllers.PromoteUser)

	return r
}
