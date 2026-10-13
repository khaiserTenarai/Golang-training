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

	fmt.Print("What should your app be called? ")
	nameText, _ := reader.ReadString('\n')
	appName := strings.TrimSpace(nameText)

	fmt.Print("Port to run the server on (press Enter for 8080): ")
	portText, _ := reader.ReadString('\n')
	port := strings.TrimSpace(portText)
	if port == "" {
		port = "8080"
	}

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"app":     appName,
			"message": "Gin is up and running!",
		})
	})

	if err := router.Run(":" + port); err != nil {
		fmt.Println("Server stopped:", err)
	}
}
