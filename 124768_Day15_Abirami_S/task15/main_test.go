package main

import (
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestLogin(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := setupRouter()
	q := httptest.NewRequest("POST", "/login", strings.NewReader(`{"username":"admin","password":"admin123"}`))
	q.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, q)
	assert.Equal(t, 200, w.Code)
	assert.Contains(t, w.Body.String(), "token")
}
func TestProtectedWithoutToken(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := setupRouter()
	q := httptest.NewRequest("GET", "/api/employees", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, q)
	assert.Equal(t, 401, w.Code)
}
func TestNotFound(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := setupRouter()
	q := httptest.NewRequest("GET", "/missing", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, q)
	assert.Equal(t, 404, w.Code)
}
