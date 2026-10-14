package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var secret = []byte("my-secret-key")

type Claims struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.RegisteredClaims
}

func main() {
	r := gin.Default()
	r.POST("/login", func(c *gin.Context) {
		claims := Claims{Username: "abirami", Role: "admin", RegisteredClaims: jwt.RegisteredClaims{ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour))}}
		t, e := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(secret)
		if e != nil {
			c.JSON(500, gin.H{"error": e.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"token": t})
	})
	r.Run(":8080")
}
