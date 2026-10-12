package main

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func newRequestID() string {
	bytes := make([]byte, 8)
	rand.Read(bytes)
	return hex.EncodeToString(bytes)
}


func requestID() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := newRequestID()
		c.Set("requestID", id)
		c.Header("X-Request-ID", id)
		c.Next()
	}
}


func logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		fmt.Printf("[%s] %s %s -> %d (%v)\n",
			c.GetString("requestID"), c.Request.Method, c.Request.URL.Path,
			c.Writer.Status(), time.Since(start))
	}
}


func recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if problem := recover(); problem != nil {
				fmt.Printf("[%s] recovered from panic: %v\n", c.GetString("requestID"), problem)
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "something went wrong on the server"})
			}
		}()
		c.Next()
	}
}
