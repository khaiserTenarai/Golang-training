/*
Question 14: Implement custom 404/405 handling.

Try: curl localhost:8080/nothing-here       (404)
     curl -X PUT localhost:8080/employees   (405)
*/

package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

func readPort() string {
	fmt.Print("Enter the port (press Enter for 8080): ")
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	text = strings.TrimSpace(text)
	if text == "" {
		return "8080"
	}
	return text
}

func main() {
	router := gin.Default()
	router.HandleMethodNotAllowed = true

	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "route not found",
			"path":  c.Request.URL.Path,
		})
	})

	router.NoMethod(func(c *gin.Context) {
		c.JSON(http.StatusMethodNotAllowed, gin.H{
			"error":  "method not allowed",
			"method": c.Request.Method,
		})
	})

	router.GET("/employees", func(c *gin.Context) {
		c.JSON(http.StatusOK, []string{"Asha", "Ravi"})
	})

	router.POST("/employees", func(c *gin.Context) {
		c.JSON(http.StatusCreated, gin.H{"message": "employee created"})
	})

	router.Run(":" + readPort())
}
