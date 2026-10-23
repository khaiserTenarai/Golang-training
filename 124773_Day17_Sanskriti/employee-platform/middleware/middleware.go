package middleware

import (
	"net/http"
	"os"
	"strings"

	"employee-platform/utility"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func RequestID() gin.HandlerFunc { return func(c *gin.Context){id:=c.GetHeader("X-Request-ID");if id==""{id=uuid.NewString()};c.Header("X-Request-ID",id);c.Set("request_id",id);c.Next()} }
func Authenticate(secret string)gin.HandlerFunc{return func(c *gin.Context){authorize(c,secret,true)}}
func OptionalAuthenticate(secret string)gin.HandlerFunc{return func(c *gin.Context){authorize(c,secret,false)}}
func authorize(c *gin.Context,secret string,required bool){
	h:=strings.TrimSpace(c.GetHeader("Authorization"))
	if h=="" {if required{c.AbortWithStatusJSON(http.StatusUnauthorized,gin.H{"error":"missing bearer token"})};return}
	parts:=strings.SplitN(h," ",2);if len(parts)!=2||!strings.EqualFold(parts[0],"Bearer"){c.AbortWithStatusJSON(401,gin.H{"error":"invalid authorization header"});return}
	u,err:=utility.ParseToken(parts[1],secret);if err!=nil{c.AbortWithStatusJSON(401,gin.H{"error":"invalid or expired token"});return}
	c.Set("username",u.Username);c.Set("role",u.Role);c.Next()
}
func RequireRole(role string)gin.HandlerFunc{return func(c *gin.Context){v,_:=c.Get("role");if v!=role{c.AbortWithStatusJSON(403,gin.H{"error":"forbidden"});return};c.Next()}}
func SecretFromEnv() string {
    return os.Getenv("JWT_SECRET")
}
