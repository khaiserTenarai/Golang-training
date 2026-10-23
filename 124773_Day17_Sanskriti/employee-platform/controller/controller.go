package controller

import (
	"net/http"
	"strconv"

	"employee-platform/model"
	"employee-platform/service"
	"github.com/gin-gonic/gin"
)

type Controller struct{ Service *service.Service }
func New(s *service.Service)*Controller{return &Controller{Service:s}}
func(cn *Controller)Login(c *gin.Context){
	var in model.LoginInput;if err:=c.ShouldBindJSON(&in);err!=nil{c.JSON(400,gin.H{"error":"invalid JSON request"});return}
	token,err:=cn.Service.Login(in.Username,in.Password);if err!=nil{c.JSON(401,gin.H{"error":"invalid credentials"});return}
	c.JSON(200,gin.H{"token":token,"token_type":"Bearer"})
}
func parseID(c *gin.Context)(int64,error){return strconv.ParseInt(c.Param("id"),10,64)}
func(cn *Controller)ListEmployees(c *gin.Context){
	page:=intQuery(c,"page",1);limit:=intQuery(c,"limit",10)
	items,total,err:=cn.Service.ListEmployees(c.Request.Context(),page,limit,c.Query("name"),c.Query("department"),c.DefaultQuery("sortBy","id"),c.DefaultQuery("order","ASC"))
	if err!=nil{c.JSON(400,gin.H{"error":err.Error()});return}
	c.JSON(200,gin.H{"data":items,"page":page,"limit":limit,"total":total})
}
func(cn *Controller)GetEmployee(c *gin.Context){id,e:=parseID(c);if e!=nil||id<1{c.JSON(400,gin.H{"error":"invalid employee id"});return};v,e:=cn.Service.GetEmployee(c.Request.Context(),id);if e!=nil{c.JSON(404,gin.H{"error":"employee not found"});return};c.JSON(200,v)}
func(cn *Controller)CreateEmployee(c *gin.Context){var in model.EmployeeInput;if e:=c.ShouldBindJSON(&in);e!=nil{c.JSON(400,gin.H{"error":"invalid JSON request"});return};v,e:=cn.Service.CreateEmployee(c.Request.Context(),in);if e!=nil{c.JSON(400,gin.H{"error":e.Error()});return};c.JSON(201,v)}
func(cn *Controller)UpdateEmployee(c *gin.Context){id,e:=parseID(c);if e!=nil||id<1{c.JSON(400,gin.H{"error":"invalid employee id"});return};var in model.EmployeeInput;if e=c.ShouldBindJSON(&in);e!=nil{c.JSON(400,gin.H{"error":"invalid JSON request"});return};v,e:=cn.Service.UpdateEmployee(c.Request.Context(),id,in);if e!=nil{c.JSON(400,gin.H{"error":e.Error()});return};c.JSON(200,v)}
func(cn *Controller)DeleteEmployee(c *gin.Context){id,e:=parseID(c);if e!=nil||id<1{c.JSON(400,gin.H{"error":"invalid employee id"});return};if e=cn.Service.DeleteEmployee(c.Request.Context(),id);e!=nil{c.JSON(404,gin.H{"error":e.Error()});return};c.Status(http.StatusNoContent)}
func(cn *Controller)ListDepartments(c *gin.Context){v,e:=cn.Service.ListDepartments(c.Request.Context());if e!=nil{c.JSON(500,gin.H{"error":"database error"});return};c.JSON(200,v)}
func intQuery(c *gin.Context,key string,def int)int{v,e:=strconv.Atoi(c.DefaultQuery(key,strconv.Itoa(def)));if e!=nil{return def};return v}
