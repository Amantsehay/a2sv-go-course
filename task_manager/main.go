

package main

import (
    "fmt"
    "task_manager/data"
    "task_manager/router"
)

func main() {
    data.Init()
    
    r := router.SetupRouter()
    
    fmt.Println("Server running on port 8080")
    r.Run(":8080") 
}
