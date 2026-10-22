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

// Abort stops the chain: the middlewares and handler after this one
// will not run at all.
func passwordCheck(password string) gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.GetHeader("X-Password") != password {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "wrong or missing password"})
			return
		}
		c.Next()
	}
}

func announce() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Password check passed, moving on")
		c.Next()
	}
}

func main() {
	port := ask("Enter port number (press Enter for 8080): ")
	if port == "" {
		port = "8080"
	}

	password := ask("Set a password that callers must send in the X-Password header: ")
	if password == "" {
		fmt.Println("The password cannot be empty.")
		return
	}

	router := gin.New()

	router.GET("/public", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "anyone can see this"})
	})

	private := router.Group("/private")
	private.Use(passwordCheck(password), announce())
	private.GET("/data", func(c *gin.Context) {
		fmt.Println("Private handler is running")
		c.JSON(http.StatusOK, gin.H{"message": "here is the private data"})
	})

	if err := router.Run(":" + port); err != nil {
		fmt.Println("Could not start the server:", err)
	}
}
