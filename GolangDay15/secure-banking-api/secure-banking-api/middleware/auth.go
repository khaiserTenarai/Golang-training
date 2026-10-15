package middleware

import (
	"net/http"
	"strings"

	"secure-banking-api/util"

	"github.com/gin-gonic/gin"
)

func unauthorized(c *gin.Context, msg string) {
	
	c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"success": false, "message": msg})
}


func JWTAuth(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")
		parts := strings.SplitN(header, " ", 2)
		if len(parts) != 2 || !strings.EqualFold(parts[0], "Bearer") {
			unauthorized(c, "missing or malformed Authorization header")
			return
		}

		claims, err := util.ParseToken(secret, parts[1])
		if err != nil {
			unauthorized(c, "invalid or expired token")
			return
		}

		c.Set(CtxUserID, claims.UserID)
		c.Set(CtxRole, claims.Role)
		c.Next()
	}
}
