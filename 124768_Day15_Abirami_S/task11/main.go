package main

import "github.com/gin-gonic/gin"

type Employee struct {
	Name   string `json:"name" binding:"required,min=3"`
	Email  string `json:"email" binding:"required,email"`
	Salary int    `json:"salary" binding:"required,min=1"`
}

func main() {
	r := gin.Default()
	r.POST("/employees", func(c *gin.Context) {
		var e Employee
		if err := c.ShouldBindJSON(&e); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		c.JSON(201, e)
	})
	r.Run(":8080")
}
