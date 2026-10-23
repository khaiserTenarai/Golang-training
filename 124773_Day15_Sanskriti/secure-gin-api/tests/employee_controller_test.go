package tests

import (
 "net/http"
 "net/http/httptest"
 "testing"
 "github.com/gin-gonic/gin"
)
func TestHealthHandlerWithHTTPTest(t *testing.T){gin.SetMode(gin.TestMode);r:=gin.New();r.GET("/health",func(c *gin.Context){c.JSON(http.StatusOK,gin.H{"status":"UP"})});req:=httptest.NewRequest(http.MethodGet,"/health",nil);rec:=httptest.NewRecorder();r.ServeHTTP(rec,req);if rec.Code!=http.StatusOK{t.Fatalf("want 200, got %d",rec.Code)};if rec.Body.String()==""{t.Fatal("expected response body")}}
func TestAbortStopsFollowingHandler(t *testing.T){gin.SetMode(gin.TestMode);r:=gin.New();ran:=false;r.GET("/abort",func(c *gin.Context){c.AbortWithStatusJSON(http.StatusForbidden,gin.H{"error":"blocked"})},func(c *gin.Context){ran=true});rec:=httptest.NewRecorder();r.ServeHTTP(rec,httptest.NewRequest(http.MethodGet,"/abort",nil));if rec.Code!=http.StatusForbidden{t.Fatalf("want 403, got %d",rec.Code)};if ran{t.Fatal("handler after Abort ran")}}
