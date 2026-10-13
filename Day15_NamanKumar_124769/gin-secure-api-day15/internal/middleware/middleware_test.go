package middleware

import (
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestNextExecutionOrder(t *testing.T) {
	gin.SetMode(gin.TestMode)
	order := []string{}
	r := gin.New()
	r.Use(func(c *gin.Context) {
		order = append(order, "A-before")
		c.Next()
		order = append(order, "A-after")
	})
	r.Use(func(c *gin.Context) {
		order = append(order, "B-before")
		c.Next()
		order = append(order, "B-after")
	})
	r.GET("/", func(c *gin.Context) {
		order = append(order, "handler")
		c.Status(200)
	})
	r.ServeHTTP(httptest.NewRecorder(), httptest.NewRequest("GET", "/", nil))
	want := []string{"A-before", "B-before", "handler", "B-after", "A-after"}
	if !reflect.DeepEqual(order, want) {
		t.Fatalf("order = %#v, want %#v", order, want)
	}
}

func TestAbortStopsLaterHandlers(t *testing.T) {
	gin.SetMode(gin.TestMode)
	called := false
	r := gin.New()
	r.Use(func(c *gin.Context) {
		c.AbortWithStatusJSON(401, gin.H{"error": "blocked"})
	})
	r.GET("/", func(c *gin.Context) {
		called = true
		c.Status(200)
	})
	rec := httptest.NewRecorder()
	r.ServeHTTP(rec, httptest.NewRequest("GET", "/", nil))
	if called {
		t.Fatal("route handler ran after Abort")
	}
	if rec.Code != 401 {
		t.Fatalf("status = %d, want 401", rec.Code)
	}
}
