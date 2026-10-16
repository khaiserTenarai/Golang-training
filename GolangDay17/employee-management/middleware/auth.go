package middleware

import (
	"employee-management/auth"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)


func bearerToken(c *gin.Context) string {
	h := c.GetHeader("Authorization")
	if len(h) > 7 && strings.EqualFold(h[:7], "Bearer ") {
		return strings.TrimSpace(h[7:])
	}
	return ""
}


func attach(c *gin.Context, claims *auth.Claims) {
	c.Request = c.Request.WithContext(auth.WithUser(c.Request.Context(), claims))
}


func Auth(t *auth.TokenService) gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, err := t.Parse(bearerToken(c))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "authentication required"})
			return
		}
		attach(c, claims)
		c.Next()
	}
}


func OptionalAuth(t *auth.TokenService) gin.HandlerFunc {
	return func(c *gin.Context) {
		if token := bearerToken(c); token != "" {
			if claims, err := t.Parse(token); err == nil {
				attach(c, claims)
			}
		}
		c.Next()
	}
}


func RequireRole(role string) gin.HandlerFunc {
	return func(c *gin.Context) {
		if _, err := auth.RequireRole(c.Request.Context(), role); err != nil {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "you do not have permission to do this"})
			return
		}
		c.Next()
	}
}
