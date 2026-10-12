package controller

import (
 "context"
 "database/sql"
 "errors"
 "net/http"
 "strconv"
 "strings"
 "time"

 "employee-management/model"
 "employee-management/repository"
 "employee-management/service"
 "github.com/gin-gonic/gin"
)

type Controller struct { svc service.EmployeeService; db *sql.DB }
func New(s service.EmployeeService, db *sql.DB)*Controller{return &Controller{svc:s,db:db}}
func(c *Controller) Routes(r *gin.Engine){
 r.GET("/health",func(x *gin.Context){x.JSON(200,gin.H{"status":"ok"})})
 r.GET("/ready",c.ready)
 g:=r.Group("/api/employees")
 g.POST("",c.create); g.GET("",c.list); g.GET("/:id",c.get); g.PUT("/:id",c.update); g.PATCH("/:id",c.patch); g.DELETE("/:id",c.delete)
}
func(c *Controller) ready(x *gin.Context){
 ctx,cancel:=context.WithTimeout(x.Request.Context(),2*time.Second);defer cancel()
 if err:=c.db.PingContext(ctx);err!=nil{x.JSON(503,gin.H{"status":"not_ready"});return}
 x.JSON(200,gin.H{"status":"ready","database":"connected"})
}
func(c *Controller) create(x *gin.Context){
 var in model.EmployeeInput;if x.ShouldBindJSON(&in)!=nil{x.JSON(400,gin.H{"error":"invalid JSON body"});return}
 e,err:=c.svc.Create(x.Request.Context(),in);if err!=nil{c.fail(x,err);return};x.JSON(201,gin.H{"message":"employee created","employee":e})
}
func(c *Controller) list(x *gin.Context){
 page,e:=positive(x.DefaultQuery("page","1"));if e!=nil{x.JSON(400,gin.H{"error":"page must be positive"});return}
 limit,e:=positive(x.DefaultQuery("limit","10"));if e!=nil||limit>100{x.JSON(400,gin.H{"error":"limit must be 1..100"});return}
 data,total,err:=c.svc.List(x.Request.Context(),page,limit,x.Query("search"),strings.ToLower(x.DefaultQuery("sort","id")))
 if err!=nil{x.JSON(500,gin.H{"error":"could not list employees"});return}
 x.JSON(200,gin.H{"data":data,"pagination":gin.H{"page":page,"limit":limit,"total":total}})
}
func(c *Controller) get(x *gin.Context){id,ok:=idParam(x);if !ok{return};e,err:=c.svc.Get(x.Request.Context(),id);if err!=nil{c.fail(x,err);return};x.JSON(200,gin.H{"employee":e})}
func(c *Controller) update(x *gin.Context){id,ok:=idParam(x);if !ok{return};var in model.EmployeeInput;if x.ShouldBindJSON(&in)!=nil{x.JSON(400,gin.H{"error":"invalid JSON body"});return};e,err:=c.svc.Update(x.Request.Context(),id,in);if err!=nil{c.fail(x,err);return};x.JSON(200,gin.H{"message":"employee updated","employee":e})}
func(c *Controller) patch(x *gin.Context){id,ok:=idParam(x);if !ok{return};var p model.EmployeePatch;if x.ShouldBindJSON(&p)!=nil{x.JSON(400,gin.H{"error":"invalid JSON body"});return};e,err:=c.svc.Patch(x.Request.Context(),id,p);if err!=nil{c.fail(x,err);return};x.JSON(200,gin.H{"message":"employee updated","employee":e})}
func(c *Controller) delete(x *gin.Context){id,ok:=idParam(x);if !ok{return};if err:=c.svc.Delete(x.Request.Context(),id);err!=nil{c.fail(x,err);return};x.JSON(200,gin.H{"message":"employee deleted"})}
func idParam(x *gin.Context)(int64,bool){id,e:=strconv.ParseInt(x.Param("id"),10,64);if e!=nil||id<1{x.JSON(400,gin.H{"error":"id must be positive integer"});return 0,false};return id,true}
func positive(s string)(int,error){n,e:=strconv.Atoi(s);if e!=nil||n<1{return 0,errors.New("must be positive")};return n,nil}
func(c *Controller) fail(x *gin.Context,e error){
 switch {case errors.Is(e,service.ErrInvalid):x.JSON(400,gin.H{"error":"invalid employee: name required, valid email required, age >= 18, salary >= 0"})
 case errors.Is(e,repository.ErrNotFound):x.JSON(404,gin.H{"error":"employee not found"})
 case strings.Contains(strings.ToLower(e.Error()),"duplicate key"),strings.Contains(strings.ToLower(e.Error()),"unique constraint"):x.JSON(409,gin.H{"error":"email already exists"})
 default:x.JSON(http.StatusInternalServerError,gin.H{"error":"internal server error"})}
}
