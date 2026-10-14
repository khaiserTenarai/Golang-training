package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type RegisterRequest struct {
	Name     string `json:"name" binding:"required,min=3,max=50"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
	Age      int    `json:"age" binding:"required,gte=18,lte=100"`
}

func Register(c *gin.Context) {
	var req RegisterRequest

	// Bind and validate the request body
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Validation failed",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User registered successfully",
		"data": gin.H{
			"name":  req.Name,
			"email": req.Email,
			"age":   req.Age,
		},
	})
}

func main() {
	r := gin.Default()

	r.POST("/register", Register)

	r.Run(":8080")
}
