/*
Question 7: Create route groups.

Try: curl localhost:8080/api/v1/employees
     curl localhost:8080/api/v2/employees
     curl localhost:8080/admin/stats
     curl -H "X-Admin: yes" localhost:8080/admin/stats
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
	ID         int     `json:"id"`
	Name       string  `json:"name"`
	Department string  `json:"department"`
	Salary     float64 `json:"salary"`
}

var employees = []Employee{
	{1, "Asha", "Sales", 30000},
	{2, "Ravi", "Tech", 45000},
	{3, "Meena", "HR", 38000},
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

func adminOnly() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.GetHeader("X-Admin") != "yes" {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "admins only"})
			return
		}
		c.Next()
	}
}

func main() {
	router := gin.Default()

	v1 := router.Group("/api/v1")
	{
		v1.GET("/employees", func(c *gin.Context) {
			names := []string{}
			for _, e := range employees {
				names = append(names, e.Name)
			}
			c.JSON(http.StatusOK, names)
		})
	}

	v2 := router.Group("/api/v2")
	{
		v2.GET("/employees", func(c *gin.Context) {
			c.JSON(http.StatusOK, employees)
		})
	}

	admin := router.Group("/admin", adminOnly())
	{
		admin.GET("/stats", func(c *gin.Context) {
			total := 0.0
			for _, e := range employees {
				total += e.Salary
			}
			c.JSON(http.StatusOK, gin.H{"employees": len(employees), "total_salary": total})
		})
	}

	router.Run(":" + readPort())
}
