package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()
    router.GET("/employees/:id", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "employee_id": c.Param("id"),
        })
    })
    router.GET("/departments/:departmentID/employees/:employeeID", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "department_id": c.Param("departmentID"),
            "employee_id": c.Param("employeeID"),
        })
    })
    router.Run(":10014")
}
