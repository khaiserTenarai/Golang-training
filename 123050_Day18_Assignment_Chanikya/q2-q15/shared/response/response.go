package response

import "github.com/gin-gonic/gin"

// Error sends a standard JSON error response.
func Error(c *gin.Context, status int, message string) {
	c.JSON(status, gin.H{
		"error": message,
	})
}
