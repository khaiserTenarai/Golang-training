/*
Question 9: Implement query parameters.

Try: curl "localhost:8080/search?name=asha"
     curl "localhost:8080/search?dept=Tech&min_salary=40000"
     curl "localhost:8080/search?sort=salary"
*/

package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"sort"
	"strconv"
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
	{4, "Karthik", "Tech", 52000},
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

	router.GET("/search", func(c *gin.Context) {
		name := c.Query("name")
		dept := c.DefaultQuery("dept", "all")
		minSalary, err := strconv.ParseFloat(c.DefaultQuery("min_salary", "0"), 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "min_salary must be a number"})
			return
		}

		found := []Employee{}
		for _, e := range employees {
			if name != "" && !strings.Contains(strings.ToLower(e.Name), strings.ToLower(name)) {
				continue
			}
			if dept != "all" && !strings.EqualFold(e.Department, dept) {
				continue
			}
			if e.Salary < minSalary {
				continue
			}
			found = append(found, e)
		}

		if sortBy, given := c.GetQuery("sort"); given && sortBy == "salary" {
			sort.Slice(found, func(i, j int) bool { return found[i].Salary > found[j].Salary })
		}

		c.JSON(http.StatusOK, gin.H{"count": len(found), "employees": found})
	})

	router.Run(":" + readPort())
}
