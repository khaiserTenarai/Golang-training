package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func hello(c *gin.Context) { c.JSON(200, gin.H{"message": "Hello"}) }
func TestHello(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.GET("/hello", hello)
	req := httptest.NewRequest(http.MethodGet, "/hello", nil)
	rec := httptest.NewRecorder()
	r.ServeHTTP(rec, req)
	assert.Equal(t, 200, rec.Code)
	assert.Contains(t, rec.Body.String(), "Hello")
}
