package middleware

import (
	"net/http"

	"employee-management/logging"

	"github.com/gin-gonic/gin"
)

func Recovery(logger *logging.Logger) gin.HandlerFunc {
	return gin.CustomRecovery(func(c *gin.Context, recovered any) {
		logger.Error("panic recovered")

		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": "internal server error",
		})
	})
}
