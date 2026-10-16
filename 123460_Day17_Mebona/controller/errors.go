package controller

import (
	"employee-management/service"
	"errors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// fail turns a service error into the right HTTP status code.
func fail(ctx *gin.Context, err error, entity string) {
	switch {
	case errors.Is(err, service.ErrInvalidCredentials):
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
	case service.IsNotFound(err):
		ctx.JSON(http.StatusNotFound, gin.H{"error": entity + " not found"})
	case service.IsValidation(err):
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	case service.IsDuplicate(err):
		ctx.JSON(http.StatusConflict, gin.H{"error": err.Error()})
	default:
		log.Println("internal error:", err) // details stay in the server log
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "something went wrong"})
	}
}
