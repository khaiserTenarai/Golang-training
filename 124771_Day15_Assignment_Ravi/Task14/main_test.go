package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestHello(t *testing.T) {
	gin.SetMode(gin.TestMode)

	r := gin.New()
	r.GET("/hello", Hello)

	req := httptest.NewRequest("GET", "/hello", nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}

	if w.Body.String() != "Hello World" {
		t.Errorf("Expected Hello World, got %s", w.Body.String())
	}
}
