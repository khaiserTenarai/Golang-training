package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SharedService struct{}

func (s *SharedService) GetData() string {
	return "Shared Core Service Data"
}

func main() {
	svc := &SharedService{}
	r := gin.Default()

	// REST Endpoint
	r.GET("/api/data", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": svc.GetData()})
	})

	// GraphQL Endpoint
	r.POST("/graphql", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": map[string]string{"gql": svc.GetData()}})
	})

	fmt.Println("Service exposed on REST (/api/data) and GraphQL (/graphql)")
}
