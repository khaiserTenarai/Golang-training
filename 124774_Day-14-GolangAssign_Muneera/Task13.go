package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Static("/assets", "./assets")
	fmt.Println("Server started at http://localhost:8080")

	r.Run(":8080")
}

// Request URL: http://localhost:8080/assets/hello.txt
// Method: GET

// Output:
// Contents of the assets/hello.txt file
// Example: Hello! This is a static file served by Gin.
