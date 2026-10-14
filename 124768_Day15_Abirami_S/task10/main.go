package main

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var secret = []byte("my-secret-key")

func auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		h := c.GetHeader("Authorization")
		if !strings.HasPrefix(h, "Bearer ") {
			c.AbortWithStatusJSON(401, gin.H{"error": "Unauthorized"})
			return
		}
		t, e := jwt.Parse(strings.TrimPrefix(h, "Bearer "), func(t *jwt.Token) (interface{}, error) { return secret, nil })
		if e != nil || !t.Valid {
			c.AbortWithStatusJSON(401, gin.H{"error": "Invalid token"})
			return
		}
		claims, _ := t.Claims.(jwt.MapClaims)
		c.Set("role", claims["role"])
		c.Next()
	}
}
func adminOnly() gin.HandlerFunc {
	return func(c *gin.Context) {
		role, _ := c.Get("role")
		if role != "admin" {
			c.AbortWithStatusJSON(403, gin.H{"error": "Admin role required"})
			return
		}
		c.Next()
	}
}
func main() {
	r := gin.Default()
	r.GET("/admin", auth(), adminOnly(), func(c *gin.Context) { c.JSON(200, gin.H{"message": "Welcome admin"}) })
	r.Run(":8080")
	_ = http.StatusOK
}
