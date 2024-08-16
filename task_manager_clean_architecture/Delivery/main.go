package main 


import (
    "fmt"
    "log"
    "task_manager_clean_architecture/Infrastructure"
    "task_manager_clean_architecture/Repositories"
    "task_manager_clean_architecture/Delivery/router"
)


func main(){
    db, err := Infrastructure.InitDB()
    if err != nil {
        log.Fatal(err)
    }

    taskRepo := Repositories.NewMongoTaskRepository(db)
    userRepo := Repositories.NewMongoUserRepository(db)

    r := router.SetupRouter(taskRepo, userRepo)
    fmt.Println("Server running on port 8080")

    if err := r.Run(":8080"); err != nil {
        log.Fatalf("Failed to run server: %v", err)
    }


}