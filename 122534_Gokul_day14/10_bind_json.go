/*
Question 10: Bind JSON using Gin.

Try: curl -X POST localhost:8080/employees -d '{"name":"Divya","email":"divya@example.com","age":28,"salary":42000}'
     curl -X POST localhost:8080/employees -d '{"name":"D","email":"not-an-email","age":10}'
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

type NewEmployee struct {
	Name   string  `json:"name" binding:"required,min=2"`
	Email  string  `json:"email" binding:"required,email"`
	Age    int     `json:"age" binding:"required,gte=18,lte=65"`
	Salary float64 `json:"salary" binding:"gte=0"`
}

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

	router.POST("/employees", func(c *gin.Context) {
		var input NewEmployee
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, gin.H{
			"message":  "employee accepted",
			"employee": input,
		})
	})

	router.Run(":" + readPort())
}
