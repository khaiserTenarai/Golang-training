package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestHealthRouteExample(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r:=gin.New()
	r.GET("/health",func(c *gin.Context){c.JSON(http.StatusOK,gin.H{"status":"UP"})})
	w:=httptest.NewRecorder()
	req:=httptest.NewRequest(http.MethodGet,"/health",nil)
	r.ServeHTTP(w,req)
	if w.Code!=http.StatusOK {t.Fatalf("expected 200, got %d",w.Code)}
}
