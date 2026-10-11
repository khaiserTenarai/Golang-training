package middleware

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var JwtSecretKey = []byte("super_secret_jwt_key")

// Topic 1, 2, 3: Custom Middleware, Execution Order & Next()
func ExecutionOrderLogger(name string) gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Printf("[%s] Pre-processing request: %s %s\n", name, c.Request.Method, c.Request.URL.Path)
		c.Next() // Topic 3: Pass control to the next handler/middleware
		fmt.Printf("[%s] Post-processing response\n", name)
	}
}

// Topic 5: Request ID Middleware
func RequestIDMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		reqID := fmt.Sprintf("req-%d", time.Now().UnixNano())
		c.Set("RequestID", reqID)
		c.Header("X-Request-ID", reqID)
		c.Next()
	}
}

// Topic 6: Recovery Middleware (Catches runtime panics gracefully)
func CustomRecovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"error":   "Internal Server Error",
					"details": fmt.Sprintf("%v", err),
				})
				c.Abort() // Topic 4: Abort request chain on panic
			}
		}()
		c.Next()
	}
}

// Topic 7: CORS Middleware Configuration
func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent) // Topic 4: Abort pre-flight request directly
			return
		}
		c.Next()
	}
}

// Topic 8 & 9: JWT Authentication Middleware & Protecting Routes
func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token required"})
			c.Abort() // Topic 4: Halt pipeline if unauthenticated
			return
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method")
			}
			return JwtSecretKey, nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			c.Abort()
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			c.Set("username", claims["username"])
			c.Set("role", claims["role"])
		}

		c.Next()
	}
}

// Topic 10: Role-Based Access Control (RBAC)
func RequireRole(allowedRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userRole, exists := c.Get("role")
		if !exists {
			c.JSON(http.StatusForbidden, gin.H{"error": "User role not identified"})
			c.Abort()
			return
		}

		for _, role := range allowedRoles {
			if strings.EqualFold(fmt.Sprintf("%v", userRole), role) {
				c.Next()
				return
			}
		}

		c.JSON(http.StatusForbidden, gin.H{"error": "Access denied: insufficient permissions"})
		c.Abort()
	}
}