package main

import (
    "fmt"
    "log"
    "task_manager_clean_architecture/Infrastructure"
    "task_manager_clean_architecture/Repositories"

    "task_manager_clean_architecture/Delivery" // Import the router package
)

func main() {
    // Initialize the database
    _, err := Infrastructure.InitDB()
    if err != nil {
        log.Fatal(err)
    }

    _, err = Repositories.NewMongoTaskRepository() 
    if err != nil {
        log.Fatal(err)
    }

    _, err = Repositories.NewMongoUserRepository() // Assign to existing err
    if err != nil {
        log.Fatal(err)
    }

    // Setup the router
    r := SetupRouter() // Make sure the router is properly imported and used
    fmt.Println("Server running on port 8080")

    // Start the server
    if err := r.Run(":8080"); err != nil {
        log.Fatalf("Failed to run server: %v", err)
    }
}
