package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

// appLogger uses Go's standard log package instead of slog.
type appLogger struct {
	level string
}

func newLogger(level string) *appLogger {
	return &appLogger{level: strings.ToLower(level)}
}

func (l *appLogger) Info(message string, args ...any) {
	if l.level == "debug" || l.level == "info" {
		log.Printf("[INFO] %s%s", message, formatLogArgs(args...))
	}
}

func (l *appLogger) Error(message string, args ...any) {
	log.Printf("[ERROR] %s%s", message, formatLogArgs(args...))
}

func formatLogArgs(args ...any) string {
	if len(args) == 0 {
		return ""
	}
	parts := make([]string, 0, len(args)/2+1)
	for i := 0; i < len(args); i += 2 {
		if i+1 < len(args) {
			parts = append(parts, fmt.Sprintf("%v=%v", args[i], args[i+1]))
		} else {
			parts = append(parts, fmt.Sprint(args[i]))
		}
	}
	return " " + strings.Join(parts, " ")
}

// recoveryMiddleware catches a panic from a request handler and returns HTTP 500.
func recoveryMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("[PANIC RECOVERY] Recovered from panic: %v", err)
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"error": "An internal error occurred",
				})
			}
		}()
		c.Next()
	}
}

// loggingMiddleware logs the HTTP method, path, status, and request duration.
func loggingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		log.Printf("[HTTP] %s %s %d %s",
			c.Request.Method,
			c.Request.URL.RequestURI(),
			c.Writer.Status(),
			time.Since(start),
		)
	}
}
