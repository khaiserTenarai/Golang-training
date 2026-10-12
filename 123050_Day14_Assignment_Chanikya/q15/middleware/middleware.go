package middleware

import (
 "log"
 "time"

 "github.com/gin-gonic/gin"
 "github.com/google/uuid"
)

func RequestID() gin.HandlerFunc {
 return func(c *gin.Context) {
  id:=c.GetHeader("X-Request-ID"); if id=="" { id=uuid.NewString() }
  c.Set("request_id",id); c.Header("X-Request-ID",id); c.Next()
 }
}
func RequestLogger() gin.HandlerFunc {
 return func(c *gin.Context) {
  start:=time.Now(); c.Next()
  log.Printf("request_id=%s method=%s path=%s status=%d duration=%s",c.GetString("request_id"),c.Request.Method,c.Request.URL.Path,c.Writer.Status(),time.Since(start))
 }
}
