package main

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

type Employee struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var employees = []Employee{{1, "Abirami"}, {2, "Priya"}, {3, "Kavi"}, {4, "Anu"}, {5, "Meena"}}

func main() {
	r := gin.Default()
	r.GET("/employees", func(c *gin.Context) {
		p, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
		l, _ := strconv.Atoi(c.DefaultQuery("limit", "2"))
		if p < 1 {
			p = 1
		}
		if l < 1 {
			l = 2
		}
		s := (p - 1) * l
		if s >= len(employees) {
			c.JSON(200, gin.H{"page": p, "limit": l, "data": []Employee{}})
			return
		}
		e := s + l
		if e > len(employees) {
			e = len(employees)
		}
		c.JSON(200, gin.H{"page": p, "limit": l, "data": employees[s:e]})
	})
	r.Run(":8080")
}
