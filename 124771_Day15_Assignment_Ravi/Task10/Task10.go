package main

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type User struct {
	Username string
	Role     string
}

// Authentication middleware: validates a demo Bearer token.
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if !strings.HasPrefix(authHeader, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Missing or invalid authorization token",
			})
			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")

		// Demo tokens only. Replace with JWT or database validation.
		users := map[string]User{
			"admin-token":   {Username: "alice", Role: "admin"},
			"user-token":    {Username: "bob", Role: "user"},
			"manager-token": {Username: "carol", Role: "manager"},
		}

		user, exists := users[token]
		if !exists {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid token",
			})
			return
		}

		c.Set("username", user.Username)
		c.Set("role", user.Role)
		c.Next()
	}
}

// Role-based authorization middleware.
func RequireRoles(allowedRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		roleValue, exists := c.Get("role")
		if !exists {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Authentication required",
			})
			return
		}

		role, ok := roleValue.(string)
		if !ok {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"error": "Invalid user role",
			})
			return
		}

		for _, allowedRole := range allowedRoles {
			if role == allowedRole {
				c.Next()
				return
			}
		}

		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
			"error": "You do not have permission to access this resource",
		})
	}
}

func main() {
	r := gin.Default()

	// Public endpoint
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome! This endpoint is public.",
		})
	})

	// All routes in this group require authentication.
	api := r.Group("/api")
	api.Use(AuthMiddleware())

	// Any authenticated user
	api.GET("/profile", func(c *gin.Context) {
		username, _ := c.Get("username")
		role, _ := c.Get("role")

		c.JSON(http.StatusOK, gin.H{
			"username": username,
			"role":     role,
		})
	})

	// Admin only
	api.GET("/admin", RequireRoles("admin"), func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to the admin dashboard",
		})
	})

	// Admin or manager
	api.GET("/reports", RequireRoles("admin", "manager"), func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "You can access reports",
		})
	})

	// Admin only: delete a resource
	api.DELETE("/users/:id", RequireRoles("admin"), func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "User deleted successfully",
			"user_id": c.Param("id"),
		})
	})

	// Any authenticated user
	api.GET("/dashboard", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to your dashboard",
		})
	})

	r.Run(":8080")
}
