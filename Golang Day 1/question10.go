package main

import (
    "fmt"
    "os" 
)

func main() {
    
    appName := os.Getenv("APP_NAME")
    port := os.Getenv("PORT")

    
    if appName == "" {
        appName = "GoApp" 
    }

    
    if port == "" {
        port = "8080" 
    }

    fmt.Println("Application Name :", appName)
    fmt.Println("Running on Port  :", port)
}