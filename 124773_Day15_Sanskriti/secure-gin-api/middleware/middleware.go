package middleware

import (
 "log"
 "net/http"
 "strings"
 "time"
 "secure-gin-api/utility"
 "github.com/gin-gonic/gin"
 "github.com/google/uuid"
)
func RequestID()gin.HandlerFunc{return func(c *gin.Context){id:=c.GetHeader("X-Request-ID");if id==""{id=uuid.NewString()};c.Set("request_id",id);c.Header("X-Request-ID",id);c.Next()}}
func Recovery()gin.HandlerFunc{return gin.CustomRecovery(func(c *gin.Context, recovered any){log.Printf("panic recovered request_id=%v error=%v",c.GetString("request_id"),recovered);c.AbortWithStatusJSON(http.StatusInternalServerError,gin.H{"error":"internal server error","request_id":c.GetString("request_id")})})}
func Logger()gin.HandlerFunc{return func(c *gin.Context){start:=time.Now();c.Next();log.Printf("request_id=%s method=%s path=%s status=%d duration=%s",c.GetString("request_id"),c.Request.Method,c.Request.URL.Path,c.Writer.Status(),time.Since(start))}}
func Authenticate(secret string)gin.HandlerFunc{return func(c *gin.Context){h:=c.GetHeader("Authorization");parts:=strings.SplitN(h," ",2);if len(parts)!=2||!strings.EqualFold(parts[0],"Bearer"){c.AbortWithStatusJSON(http.StatusUnauthorized,gin.H{"error":"missing or invalid bearer token"});return};claims,err:=utility.ParseToken(secret,parts[1]);if err!=nil{c.AbortWithStatusJSON(http.StatusUnauthorized,gin.H{"error":"invalid or expired token"});return};c.Set("user_id",claims.UserID);c.Set("username",claims.Username);c.Set("role",claims.Role);c.Next()}}
func RequireRole(roles ...string)gin.HandlerFunc{return func(c *gin.Context){role:=c.GetString("role");for _,allowed:=range roles{if role==allowed{c.Next();return}};c.AbortWithStatusJSON(http.StatusForbidden,gin.H{"error":"insufficient permissions"})}}
func DemoFirst()gin.HandlerFunc{return func(c *gin.Context){log.Println("first middleware: before Next");c.Next();log.Println("first middleware: after Next")}}
func DemoSecond()gin.HandlerFunc{return func(c *gin.Context){log.Println("second middleware: before Next");c.Next();log.Println("second middleware: after Next")}}
func BlockRequest()gin.HandlerFunc{return func(c *gin.Context){c.AbortWithStatusJSON(http.StatusForbidden,gin.H{"error":"request blocked by Abort demo"});return}}
