package routes

import (
 "secure-gin-api/controller"
 "secure-gin-api/middleware"
 "github.com/gin-gonic/gin"
)
func Setup(c *controller.EmployeeController,secret string)*gin.Engine{r:=gin.New();r.Use(middleware.RequestID(),middleware.Recovery(),middleware.Logger());r.GET("/health",c.Health);r.POST("/login",c.Login);r.GET("/demo/next",middleware.DemoFirst(),middleware.DemoSecond(),func(ctx *gin.Context){ctx.JSON(200,gin.H{"message":"middleware chain completed"})});r.GET("/demo/abort",middleware.BlockRequest(),func(ctx *gin.Context){ctx.JSON(200,gin.H{"message":"should not run"})});r.GET("/ready",func(ctx *gin.Context){ctx.JSON(200,gin.H{"status":"UP","note":"readiness is validated at startup; use DB monitoring for live checks"})});protected:=r.Group("/employees",middleware.Authenticate(secret));protected.GET("",c.List);protected.GET("/:id",c.Get);admin:=r.Group("/employees",middleware.Authenticate(secret),middleware.RequireRole("admin"));admin.POST("",c.Create);admin.PUT("/:id",c.Update);admin.DELETE("/:id",c.Delete);return r}
