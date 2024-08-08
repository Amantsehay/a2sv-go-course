package main

import "fmt"
import "github.com/gin-gonic/gin"
// import "task_manager_api/models"
import "time"

type Task struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	DueDate     time.Time `json:"due_date"`
	Status      string    `json:"status"`
   }

func main() {

	// task := models.Task 

	var tasks = []Task{
		{ID: "1", Title: "Task 1", Description: "First task", DueDate: time.Now(), Status: "Pending"},
		{ID: "2", Title: "Task 2", Description: "Second task", DueDate: time.Now().AddDate(0, 0, 1), Status: "In Progress"},
		{ID: "3", Title: "Task 3", Description: "Third task", DueDate: time.Now().AddDate(0, 0, 2), Status: "Completed"},
	}

	router := gin.Default()
	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// router get all the tasks 

	router.GET("/tasks", func(ctx *gin.Context){
		ctx.JSON(200, gin.H{"tasks": tasks})
	})

	// Get a specific task by id 
	router.GET("/tasks/:id", func(ctx *gin.Context){
		id := ctx.Param("id")
		for _, task := range tasks{
			if task.ID == id{
				ctx.JSON(200, gin.H{"task": task})
				return 
			}
		}
		ctx.JSON(404, gin.H{"message": "Task not found"})

	})


	// update a specific task by id 

	router.PUT("/tasks/:id", func(ctx *gin.Context){
		id := ctx.Param("id")
		var updatedTask Task
		
		if err := ctx.ShouldBindJSON(&updatedTask); err != nil{
			ctx.JSON(400, gin.H{"message": "Invalid request"})
			return 
		}

		for i, task := range tasks{
			
			if task.ID == id{
				if updatedTask.Title != ""{
					tasks[i].Title = updatedTask.Title
				}
				if updatedTask.Description != ""{
					tasks[i].Description = updatedTask.Description
				}
				if updatedTask.DueDate != (time.Time{}) {
					tasks[i].DueDate = updatedTask.DueDate
				}
				if updatedTask.Status != ""{
					tasks[i].Status = updatedTask.Status
				}
				
				ctx.JSON(200, gin.H{"message": "Task Updated Successfully"})
				return 
			}
		}
		ctx.JSON(404, gin.H{"message": "Task not found"})
	})


	router.DELETE("/tasks/:id", func(ctx *gin.Context){
		id := ctx.Param("id")
		for i, task := range tasks{
			if task.ID == id {
				tasks = append(tasks[:i], tasks[i+1:]...)
				ctx.JSON(200, gin.H{"message": "Task deleted successfully"})
				return 
			}

		}
		ctx.JSON(404, gin.H{"message": "Task failt to delete"})
	})

	// for creating new tasks 
	router.POST("/tasks", func(ctx *gin.Context){
		var newTask Task 

		if err := ctx.ShouldBindJSON(&newTask); err != nil{
			ctx.JSON(400, gin.H{"message": "invalid request"})
			return 
		}
		tasks = append(tasks, newTask)
		ctx.JSON(200, gin.H{"message": "Task created successfully"})

	})



	fmt.Println("Server is running at port 8080")
	router.Run() 
}


