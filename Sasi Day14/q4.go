package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UpdateUserRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {
	r := gin.Default()

	r.PUT("/users/:id", func(c *gin.Context) {
		id := c.Param("id")
		var req UpdateUserRequest

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"id":      id,
			"name":    req.Name,
			"email":   req.Email,
			"status":  "updated (full replacement)",
		})
	})

	r.Run(":8080")
}
