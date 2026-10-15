package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type PatchUserRequest struct {
	Name  *string `json:"name,omitempty"`
	Email *string `json:"email,omitempty"`
}

func main() {
	r := gin.Default()

	r.PATCH("/users/:id", func(c *gin.Context) {
		id := c.Param("id")
		var req PatchUserRequest

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"id":      id,
			"updated": req,
			"message": "User partially updated",
		})
	})

	r.Run(":8080")
}
