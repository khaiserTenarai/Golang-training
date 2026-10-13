package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("What message should unknown URLs get? ")
	messageText, _ := reader.ReadString('\n')
	notFoundMessage := strings.TrimSpace(messageText)
	if notFoundMessage == "" {
		notFoundMessage = "that route doesn't exist"
	}

	fmt.Print("Port to run the server on (press Enter for 8080): ")
	portText, _ := reader.ReadString('\n')
	port := strings.TrimSpace(portText)
	if port == "" {
		port = "8080"
	}

	router := gin.Default()
	router.HandleMethodNotAllowed = true

	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"error": notFoundMessage,
			"path":  c.Request.URL.Path,
		})
	})

	router.NoMethod(func(c *gin.Context) {
		c.JSON(http.StatusMethodNotAllowed, gin.H{
			"error":  "that URL doesn't support this HTTP method",
			"method": c.Request.Method,
			"path":   c.Request.URL.Path,
		})
	})

	router.GET("/employee", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "the employee list would be here"})
	})
	router.POST("/employee", func(c *gin.Context) {
		c.JSON(http.StatusCreated, gin.H{"message": "an employee would be created here"})
	})

	if err := router.Run(":" + port); err != nil {
		fmt.Println("Server stopped:", err)
	}
}
