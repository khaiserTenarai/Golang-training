package main

import (
    "net/http"
    "strconv"
    "github.com/gin-gonic/gin"
)

type Employee struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
}

var employees = map[int]Employee{
    1: {ID: 1, Name: "Asha"},
    2: {ID: 2, Name: "Rahul"},
}

func main() {
    router := gin.Default()
    router.DELETE("/employees/:id", func(c *gin.Context) {
        id, err := strconv.Atoi(c.Param("id"))
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "id must be an integer"})
            return
        }
        if _, ok := employees[id]; !ok {
            c.JSON(http.StatusNotFound, gin.H{"error": "employee not found"})
            return
        }
        delete(employees, id)
        c.JSON(http.StatusOK, gin.H{"message": "Employee deleted", "id": id})
    })
    router.Run(":10014")
}
