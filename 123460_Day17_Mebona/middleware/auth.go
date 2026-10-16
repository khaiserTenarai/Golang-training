package middleware

import (
	"employee-management/auth"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// bearerToken reads the token from the header   Authorization: Bearer <token>
func bearerToken(c *gin.Context) string {
	h := c.GetHeader("Authorization")
	if len(h) > 7 && strings.EqualFold(h[:7], "Bearer ") {
		return strings.TrimSpace(h[7:])
	}
	return ""
}

// attach stores the logged-in user inside the request, so controllers and GraphQL resolvers can read it.
func attach(c *gin.Context, claims *auth.Claims) {
	c.Request = c.Request.WithContext(auth.WithUser(c.Request.Context(), claims))
}

// Auth = AUTHENTICATION. The request must contain a valid token, otherwise 401.
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

// OptionalAuth is used for GraphQL. A token is read if it is there, but nothing is rejected here:
// the "login" mutation must work without a token, and each resolver decides what it needs.
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

// RequireRole = AUTHORIZATION. The logged-in user must have the given role, otherwise 403.
// Always put it AFTER Auth.
func RequireRole(role string) gin.HandlerFunc {
	return func(c *gin.Context) {
		if _, err := auth.RequireRole(c.Request.Context(), role); err != nil {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "you do not have permission to do this"})
			return
		}
		c.Next()
	}
}
