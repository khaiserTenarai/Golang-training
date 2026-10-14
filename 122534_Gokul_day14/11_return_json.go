/*
Question 11: Return JSON responses.

Try: curl localhost:8080/map
     curl localhost:8080/struct
     curl localhost:8080/list
     curl localhost:8080/pretty
     curl -i localhost:8080/accepted
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

type Employee struct {
	ID     int     `json:"id"`
	Name   string  `json:"name"`
	Salary float64 `json:"salary"`
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

	router.GET("/map", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "this came from a map", "ok": true})
	})

	router.GET("/struct", func(c *gin.Context) {
		c.JSON(http.StatusOK, Employee{ID: 1, Name: "Asha", Salary: 30000})
	})

	router.GET("/list", func(c *gin.Context) {
		c.JSON(http.StatusOK, []Employee{
			{1, "Asha", 30000},
			{2, "Ravi", 45000},
		})
	})

	router.GET("/pretty", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, gin.H{
			"company": "Acme",
			"staff":   []string{"Asha", "Ravi"},
		})
	})

	router.GET("/accepted", func(c *gin.Context) {
		c.JSON(http.StatusAccepted, gin.H{"status": "queued"})
	})

	router.Run(":" + readPort())
}
