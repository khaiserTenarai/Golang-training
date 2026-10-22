package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

var reader = bufio.NewReader(os.Stdin)

func ask(question string) string {
	fmt.Print(question)
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text)
}

func timer() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		fmt.Println("Before Next(): request came in for", c.Request.URL.Path)

		c.Next()

		fmt.Println("After Next(): handler finished with status", c.Writer.Status(), "in", time.Since(start))
	}
}

func main() {
	port := ask("Enter port number (press Enter for 8080): ")
	if port == "" {
		port = "8080"
	}

	router := gin.New()
	router.Use(timer())

	router.GET("/work", func(c *gin.Context) {
		milliseconds, _ := strconv.Atoi(c.DefaultQuery("delay", "0"))
		time.Sleep(time.Duration(milliseconds) * time.Millisecond)
		c.JSON(http.StatusOK, gin.H{"message": "work finished"})
	})

	if err := router.Run(":" + port); err != nil {
		fmt.Println("Could not start the server:", err)
	}
}
