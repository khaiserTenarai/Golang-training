package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

var reader = bufio.NewReader(os.Stdin)

func ask(question string) string {
	fmt.Print(question)
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text)
}

func requestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Request received:", c.Request.Method, c.Request.URL.Path)
		c.Header("X-Powered-By", "My Gin App")
	}
}

func main() {
	port := ask("Enter port number (press Enter for 8080): ")
	if port == "" {
		port = "8080"
	}

	router := gin.New()
	router.Use(requestLogger())

	router.GET("/hello", func(c *gin.Context) {
		name := c.DefaultQuery("name", "stranger")
		c.JSON(http.StatusOK, gin.H{"message": "Hello, " + name})
	})

	err := router.Run(":" + port)
	if err != nil {
		fmt.Println("Could not start the server:", err)
	}
}
