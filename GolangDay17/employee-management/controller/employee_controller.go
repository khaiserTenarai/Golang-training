package controller

import (
	"employee-management/model"
	"employee-management/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type EmployeeController struct{ service service.EmployeeService }

func NewEmployeeController(s service.EmployeeService) *EmployeeController {
	return &EmployeeController{s}
}


type employeeRequest struct {
	Name         string  `json:"name"`
	Email        string  `json:"email"`
	Salary       float64 `json:"salary"`
	DepartmentID uint    `json:"departmentId"`
}

func (r employeeRequest) toModel() model.Employee {
	return model.Employee{Name: r.Name, Email: r.Email, Salary: r.Salary, DepartmentID: r.DepartmentID}
}


type listQuery struct {
	Page         int      `form:"page"`
	PageSize     int      `form:"pageSize"`
	Search       string   `form:"search"`
	DepartmentID uint     `form:"departmentId"`
	MinSalary    *float64 `form:"minSalary"`
	MaxSalary    *float64 `form:"maxSalary"`
	SortBy       string   `form:"sortBy"`
	SortOrder    string   `form:"sortOrder"`
}

func parseID(ctx *gin.Context) (uint, bool) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil || id == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid employee id"})
		return 0, false
	}
	return uint(id), true
}

func (c *EmployeeController) List(ctx *gin.Context) {
	var q listQuery
	if err := ctx.ShouldBindQuery(&q); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	page, err := c.service.List(model.EmployeeFilter{
		DepartmentID: q.DepartmentID,
		Search:       q.Search,
		MinSalary:    q.MinSalary,
		MaxSalary:    q.MaxSalary,
		SortBy:       q.SortBy,
		SortOrder:    q.SortOrder,
		Page:         q.Page,
		PageSize:     q.PageSize,
	})
	if err != nil {
		fail(ctx, err, "employee")
		return
	}
	ctx.JSON(http.StatusOK, page)
}
func (c *EmployeeController) Create(ctx *gin.Context) {
	var req employeeRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	e := req.toModel()
	if err := c.service.Create(&e); err != nil {
		fail(ctx, err, "employee")
		return
	}
	ctx.JSON(http.StatusCreated, e)
}
func (c *EmployeeController) GetByID(ctx *gin.Context) {
	id, ok := parseID(ctx)
	if !ok {
		return
	}
	e, err := c.service.GetByID(id)
	if err != nil {
		fail(ctx, err, "employee")
		return
	}
	ctx.JSON(http.StatusOK, e)
}
func (c *EmployeeController) Update(ctx *gin.Context) {
	id, ok := parseID(ctx)
	if !ok {
		return
	}
	var req employeeRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	e := req.toModel()
	x, err := c.service.Update(id, &e)
	if err != nil {
		fail(ctx, err, "employee")
		return
	}
	ctx.JSON(http.StatusOK, x)
}
func (c *EmployeeController) Delete(ctx *gin.Context) {
	id, ok := parseID(ctx)
	if !ok {
		return
	}
	if err := c.service.Delete(id); err != nil {
		fail(ctx, err, "employee")
		return
	}
	ctx.Status(http.StatusNoContent)
}
