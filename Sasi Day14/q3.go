package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateUserRequest struct {
	Username string `json:"username"`
	Role     string `json:"role"`
}

func main() {
	r := gin.Default()

	r.POST("/users", func(c *gin.Context) {
		var req CreateUserRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, gin.H{
			"message":  "User created successfully",
			"username": req.Username,
			"role":     req.Role,
		})
	})

	r.Run(":8080")
}
