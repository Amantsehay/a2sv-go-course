// Delivery/router.go
package Delivery

import (
	"github.com/gin-gonic/gin"
	"task_manager_clean_architecture/Delivery/controllers"
	"task_manager_clean_architecture/Infrastructure"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Task routes
	r.GET("/tasks", controllers.GetTasks)
	r.POST("/tasks", controllers.AuthMiddleware(), controllers.CreateTask)
	r.PUT("/tasks/:id", controllers.AuthMiddleware(), controllers.UpdateTask)
	r.GET("/tasks/:id", controllers.GetTaskByID)
	r.DELETE("/tasks/:id", controllers.AuthMiddleware(), controllers.AdminMiddleware(), controllers.DeleteTaskByID)

	// User routes
	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)
	r.POST("/promote", controllers.AuthMiddleware(), controllers.AdminMiddleware(), controllers.PromoteUser)

	return r
}
