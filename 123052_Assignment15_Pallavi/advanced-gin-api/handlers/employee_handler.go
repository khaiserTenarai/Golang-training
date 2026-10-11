package handlers

import (
	"fmt"
	"net/http"
	"path/filepath"
	"strconv"

	"advanced-gin-api/models"

	"github.com/gin-gonic/gin"
)

var employeeDB = []models.Employee{
	{ID: "1", Name: "Pallavi", Role: "Software Engineer", Department: "IT", Salary: 80000},
	{ID: "2", Name: "Khaiser", Role: "Lead Architect", Department: "Engineering", Salary: 110000},
	{ID: "3", Name: "Aarav", Role: "Product Manager", Department: "Product", Salary: 90000},
	{ID: "4", Name: "Bhavana", Role: "QA Specialist", Department: "Testing", Salary: 70000},
}

// Topic 12: Implement Pagination
func GetEmployeesPaginated(c *gin.Context) {
	pageStr := c.DefaultQuery("page", "1")
	limitStr := c.DefaultQuery("limit", "2")

	page, _ := strconv.Atoi(pageStr)
	limit, _ := strconv.Atoi(limitStr)

	if page <= 0 {
		page = 1
	}
	if limit <= 0 {
		limit = 2
	}

	start := (page - 1) * limit
	end := start + limit

	if start >= len(employeeDB) {
		c.JSON(http.StatusOK, gin.H{"page": page, "limit": limit, "total": len(employeeDB), "data": []models.Employee{}})
		return
	}

	if end > len(employeeDB) {
		end = len(employeeDB)
	}

	c.JSON(http.StatusOK, gin.H{
		"page":  page,
		"limit": limit,
		"total": len(employeeDB),
		"data":  employeeDB[start:end],
	})
}

// Topic 11: Request Validation
func CreateEmployee(c *gin.Context) {
	var req models.CreateEmployeeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"validation_error": err.Error()})
		return
	}

	newEmp := models.Employee{
		ID:         fmt.Sprintf("%d", len(employeeDB)+1),
		Name:       req.Name,
		Role:       req.Role,
		Department: req.Department,
		Salary:     req.Salary,
	}

	employeeDB = append(employeeDB, newEmp)
	c.JSON(http.StatusCreated, gin.H{"message": "Employee created successfully", "data": newEmp})
}

// Topic 13: File Upload Implementation
func UploadDocument(c *gin.Context) {
	file, err := c.FormFile("document")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No file uploaded or invalid form field"})
		return
	}

	dst := filepath.Join("uploads", filepath.Base(file.Filename))
	if err := c.SaveUploadedFile(file, dst); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "File uploaded successfully", "filename": file.Filename, "path": dst})
}