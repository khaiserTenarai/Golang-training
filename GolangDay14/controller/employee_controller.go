package controller

import (
	"net/http"
	"employee_management/model"
	"employee_management/service"

	"github.com/gin-gonic/gin"
)

type EmployeeController struct {
	Service service.EmployeeService
}

func (ctrl *EmployeeController) GetEmployees(c *gin.Context) {
	department := c.Query("department")
	employees, _ := ctrl.Service.GetAll(department)
	if employees == nil {
		employees = []model.Employee{} 
	}
	c.JSON(http.StatusOK, gin.H{"data": employees})
}

func (ctrl *EmployeeController) GetEmployeeByID(c *gin.Context) {
	id := c.Param("id")
	employee, err := ctrl.Service.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Employee not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": employee})
}

func (ctrl *EmployeeController) CreateEmployee(c *gin.Context) {
	var input model.Employee
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := ctrl.Service.Create(&input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": input})
}

func (ctrl *EmployeeController) UpdateEmployeePUT(c *gin.Context) {
	id := c.Param("id")
	employee, err := ctrl.Service.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Not found"})
		return
	}

	var input model.Employee
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	employee.Name = input.Name
	employee.Department = input.Department
	employee.Role = input.Role

	ctrl.Service.Update(&employee)
	c.JSON(http.StatusOK, gin.H{"data": employee})
}

func (ctrl *EmployeeController) UpdateEmployeePATCH(c *gin.Context) {
	id := c.Param("id")
	employee, err := ctrl.Service.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Not found"})
		return
	}

	var input map[string]interface{}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if name, ok := input["name"].(string); ok { employee.Name = name }
	if dept, ok := input["department"].(string); ok { employee.Department = dept }
	if role, ok := input["role"].(string); ok { employee.Role = role }

	ctrl.Service.Update(&employee)
	c.JSON(http.StatusOK, gin.H{"data": employee})
}

func (ctrl *EmployeeController) DeleteEmployee(c *gin.Context) {
	id := c.Param("id")
	ctrl.Service.Delete(id)
	c.JSON(http.StatusOK, gin.H{"message": "Deleted successfully"})
}

func (ctrl *EmployeeController) SubmitForm(c *gin.Context) {
	name := c.PostForm("name")
	department := c.PostForm("department")
	c.JSON(http.StatusOK, gin.H{"status": "form received", "name": name, "department": department})
}