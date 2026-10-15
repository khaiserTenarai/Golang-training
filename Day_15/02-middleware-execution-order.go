package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

var reader = bufio.NewReader(os.Stdin)

func ask(question string) string {
	fmt.Print(question)
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text)
}

func tracer(name string) gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println(name, "- before handler")
		c.Next()
		fmt.Println(name, "- after handler")
	}
}

func main() {
	port := ask("Enter port number (press Enter for 8080): ")
	if port == "" {
		port = "8080"
	}

	count, err := strconv.Atoi(ask("How many middlewares do you want to chain? "))
	if err != nil || count < 1 {
		fmt.Println("Not a valid number, using 3 middlewares.")
		count = 3
	}

	router := gin.New()
	for i := 1; i <= count; i++ {
		router.Use(tracer(fmt.Sprintf("Middleware %d", i)))
	}

	router.GET("/hello", func(c *gin.Context) {
		fmt.Println("Handler is running")
		name := c.DefaultQuery("name", "stranger")
		c.JSON(http.StatusOK, gin.H{"message": "Hello, " + name})
	})

	if err := router.Run(":" + port); err != nil {
		fmt.Println("Could not start the server:", err)
	}
}
