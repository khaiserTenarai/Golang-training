// Package controller contains Gin HTTP handlers.
package controller

import (
	"employee-management/model"
	"employee-management/repository"
	"employee-management/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// EmployeeController handles employee requests.
type EmployeeController struct{ service service.EmployeeService }

// NewEmployeeController creates the controller.
func NewEmployeeController(s service.EmployeeService) *EmployeeController {
	return &EmployeeController{s}
}

// RegisterRoutes registers employee REST routes.
func (c *EmployeeController) RegisterRoutes(r *gin.Engine) {
	r.GET("/employees", c.GetAll)
	r.POST("/employees", c.Create)
	r.GET("/employees/:id", c.GetByID)
	r.PUT("/employees/:id", c.Update)
	r.DELETE("/employees/:id", c.Delete)
}

// GetAll handles the sort query parameter.
func (c *EmployeeController) GetAll(ctx *gin.Context) {
	e, err := c.service.GetAll(ctx.Request.Context(), ctx.Query("sort"))
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, e)
}

// Create binds a JSON body and creates an employee.
func (c *EmployeeController) Create(ctx *gin.Context) {
	var e model.Employee
	if err := ctx.ShouldBindJSON(&e); err != nil {
		ctx.JSON(400, gin.H{"error": "invalid JSON request body"})
		return
	}
	if err := c.service.Create(ctx.Request.Context(), &e); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, e)
}

// GetByID handles the path parameter.
func (c *EmployeeController) GetByID(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "employee id must be numeric"})
		return
	}
	e, err := c.service.GetByID(ctx.Request.Context(), id)
	if repository.IsNotFound(err) {
		ctx.JSON(404, gin.H{"error": "employee not found"})
		return
	}
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, e)
}

// Update updates an employee.
func (c *EmployeeController) Update(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "employee id must be numeric"})
		return
	}
	var e model.Employee
	if err := ctx.ShouldBindJSON(&e); err != nil {
		ctx.JSON(400, gin.H{"error": "invalid JSON request body"})
		return
	}
	u, err := c.service.Update(ctx.Request.Context(), id, &e)
	if repository.IsNotFound(err) {
		ctx.JSON(404, gin.H{"error": "employee not found"})
		return
	}
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, u)
}

// Delete deletes an employee.
func (c *EmployeeController) Delete(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "employee id must be numeric"})
		return
	}
	if err := c.service.Delete(ctx.Request.Context(), id); err != nil {
		if err.Error() == "employee not found" {
			ctx.JSON(404, gin.H{"error": "employee not found"})
			return
		}
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.Status(204)
}

// -----------

// ---------------------------

/*

// Package controller contains Gin HTTP handlers.
package controller

import (
	"employee-management/model"
	"employee-management/repository"
	"employee-management/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// EmployeeController handles employee requests.
type EmployeeController struct {
	service service.EmployeeService
}

// NewEmployeeController creates the controller.
func NewEmployeeController(s service.EmployeeService) *EmployeeController {
	return &EmployeeController{s}
}

// RegisterRoutes registers employee REST routes.
func (c *EmployeeController) RegisterRoutes(r *gin.Engine) {
	r.GET("/employees", c.GetAll)
	r.POST("/employees", c.Create)
	r.GET("/employees/:id", c.GetByID)
	r.PUT("/employees/:id", c.Update)
	r.DELETE("/employees/:id", c.Delete)
}

// GetAll godoc
// @Summary Get all employees
// @Description Retrieves all employees. Use the sort query parameter to sort results.
// @Tags Employees
// @Produce json
// @Param sort query string false "Sort employees"
// @Success 200 {array} model.Employee
// @Failure 500 {object} map[string]string
// @Router /employees [get]
func (c *EmployeeController) GetAll(ctx *gin.Context) {
	e, err := c.service.GetAll(ctx.Request.Context(), ctx.Query("sort"))
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, e)
}

// Create godoc
// @Summary Create an employee
// @Description Creates a new employee using the JSON request body.
// @Tags Employees
// @Accept json
// @Produce json
// @Param employee body model.Employee true "Employee details"
// @Success 201 {object} model.Employee
// @Failure 400 {object} map[string]string
// @Router /employees [post]
func (c *EmployeeController) Create(ctx *gin.Context) {
	var e model.Employee

	if err := ctx.ShouldBindJSON(&e); err != nil {
		ctx.JSON(400, gin.H{"error": "invalid JSON request body"})
		return
	}

	if err := c.service.Create(ctx.Request.Context(), &e); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, e)
}

// GetByID godoc
// @Summary Get employee by ID
// @Description Retrieves an employee using the employee ID.
// @Tags Employees
// @Produce json
// @Param id path int true "Employee ID"
// @Success 200 {object} model.Employee
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /employees/{id} [get]
func (c *EmployeeController) GetByID(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "employee id must be numeric"})
		return
	}

	e, err := c.service.GetByID(ctx.Request.Context(), id)
	if repository.IsNotFound(err) {
		ctx.JSON(404, gin.H{"error": "employee not found"})
		return
	}
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, e)
}

// Update godoc
// @Summary Update an employee
// @Description Updates an existing employee using the employee ID and JSON request body.
// @Tags Employees
// @Accept json
// @Produce json
// @Param id path int true "Employee ID"
// @Param employee body model.Employee true "Updated employee details"
// @Success 200 {object} model.Employee
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /employees/{id} [put]
func (c *EmployeeController) Update(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "employee id must be numeric"})
		return
	}

	var e model.Employee
	if err := ctx.ShouldBindJSON(&e); err != nil {
		ctx.JSON(400, gin.H{"error": "invalid JSON request body"})
		return
	}

	u, err := c.service.Update(ctx.Request.Context(), id, &e)
	if repository.IsNotFound(err) {
		ctx.JSON(404, gin.H{"error": "employee not found"})
		return
	}
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, u)
}

// Delete godoc
// @Summary Delete an employee
// @Description Deletes an employee using the employee ID.
// @Tags Employees
// @Param id path int true "Employee ID"
// @Success 204 "Employee deleted successfully"
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /employees/{id} [delete]
func (c *EmployeeController) Delete(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "employee id must be numeric"})
		return
	}

	if err := c.service.Delete(ctx.Request.Context(), id); err != nil {
		if err.Error() == "employee not found" {
			ctx.JSON(404, gin.H{"error": "employee not found"})
			return
		}
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(204)
}

*/
