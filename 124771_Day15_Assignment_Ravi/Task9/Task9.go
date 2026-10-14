package main

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

const jwtSecret = "my-secret-key"

func authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")
		parts := strings.SplitN(header, " ", 2)

		if len(parts) != 2 || !strings.EqualFold(parts[0], "Bearer") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Token required",
			})
			return
		}

		token, err := jwt.Parse(parts[1], func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrTokenSignatureInvalid
			}
			return []byte(jwtSecret), nil
		})

		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid token",
			})
			return
		}

		c.Next()
	}
}

func main() {
	r := gin.Default()

	// Public route
	r.GET("/public", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Public route"})
	})

	// Protected route
	protected := r.Group("/api")
	protected.Use(authMiddleware())

	protected.GET("/profile", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Protected profile"})
	})

	r.Run(":8080")
}
