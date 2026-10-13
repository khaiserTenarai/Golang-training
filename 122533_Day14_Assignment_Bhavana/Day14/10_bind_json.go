package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"
	"sync"

	"github.com/gin-gonic/gin"
)

type NewEmployee struct {
	Name       string  `json:"name" binding:"required,min=2"`
	Email      string  `json:"email" binding:"required,email"`
	Department string  `json:"department" binding:"required"`
	Salary     float64 `json:"salary" binding:"required,gt=0"`
}

type Employee struct {
	ID         int     `json:"id"`
	Name       string  `json:"name"`
	Email      string  `json:"email"`
	Department string  `json:"department"`
	Salary     float64 `json:"salary"`
}

var (
	employees = []Employee{}
	nextID    = 1
	mu        sync.Mutex
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Port to run the server on (press Enter for 8080): ")
	portText, _ := reader.ReadString('\n')
	port := strings.TrimSpace(portText)
	if port == "" {
		port = "8080"
	}

	router := gin.Default()

	router.GET("/employee", func(c *gin.Context) {
		mu.Lock()
		defer mu.Unlock()
		c.JSON(http.StatusOK, employees)
	})

	router.POST("/employee", func(c *gin.Context) {
		var input NewEmployee
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		mu.Lock()
		e := Employee{
			ID:         nextID,
			Name:       input.Name,
			Email:      input.Email,
			Department: input.Department,
			Salary:     input.Salary,
		}
		nextID++
		employees = append(employees, e)
		mu.Unlock()

		c.JSON(http.StatusCreated, e)
	})

	if err := router.Run(":" + port); err != nil {
		fmt.Println("Server stopped:", err)
	}
}
