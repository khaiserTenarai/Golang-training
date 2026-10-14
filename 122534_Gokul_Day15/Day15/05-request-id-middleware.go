package main

import (
	"bufio"
	"crypto/rand"
	"encoding/hex"
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

func newRequestID() string {
	bytes := make([]byte, 8)
	rand.Read(bytes)
	return hex.EncodeToString(bytes)
}

func requestID() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.GetHeader("X-Request-ID")
		if id == "" {
			id = newRequestID()
		}
		c.Set("requestID", id)
		c.Header("X-Request-ID", id)
		c.Next()
	}
}

func main() {
	port := ask("Enter port number (press Enter for 8080): ")
	if port == "" {
		port = "8080"
	}

	router := gin.New()
	router.Use(requestID())

	router.GET("/hello", func(c *gin.Context) {
		id := c.GetString("requestID")
		fmt.Println("Handling request", id)
		c.JSON(http.StatusOK, gin.H{"message": "hello", "request_id": id})
	})

	if err := router.Run(":" + port); err != nil {
		fmt.Println("Could not start the server:", err)
	}
}
