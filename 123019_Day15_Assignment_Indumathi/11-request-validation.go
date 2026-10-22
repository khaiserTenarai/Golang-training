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

type EmployeeRequest struct {
	Name   string  `json:"name" binding:"required,min=2"`
	Email  string  `json:"email" binding:"required,email"`
	Age    int     `json:"age" binding:"required,gte=18,lte=65"`
	Salary float64 `json:"salary" binding:"required,gt=0"`
}

func main() {
	port := ask("Enter port number (press Enter for 8080): ")
	if port == "" {
		port = "8080"
	}

	router := gin.New()

	router.POST("/employees", func(c *gin.Context) {
		var input EmployeeRequest
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, gin.H{"message": "employee is valid", "employee": input})
	})

	if err := router.Run(":" + port); err != nil {
		fmt.Println("Could not start the server:", err)
	}
}
