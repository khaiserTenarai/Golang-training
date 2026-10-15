package controller

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"

	"secure-banking-api/model"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func success(c *gin.Context, status int, data any) {
	c.JSON(status, gin.H{"success": true, "data": data})
}

func failure(c *gin.Context, status int, msg string) {
	c.JSON(status, gin.H{"success": false, "message": msg})
}


func bindError(c *gin.Context, err error) {
	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		msgs := make([]string, 0, len(ve))
		for _, fe := range ve {
			msgs = append(msgs, fmt.Sprintf("%s failed '%s' validation", fe.Field(), fe.Tag()))
		}
		failure(c, http.StatusBadRequest, strings.Join(msgs, "; "))
		return
	}
	failure(c, http.StatusBadRequest, "invalid request")
}

func handleError(c *gin.Context, err error) {
	switch {
	case errors.Is(err, model.ErrNotFound):
		failure(c, http.StatusNotFound, err.Error())
	case errors.Is(err, model.ErrForbidden):
		failure(c, http.StatusForbidden, err.Error())
	case errors.Is(err, model.ErrEmailExists):
		failure(c, http.StatusConflict, err.Error())
	case errors.Is(err, model.ErrInvalidCredentials):
		failure(c, http.StatusUnauthorized, err.Error())
	default:
		log.Printf("internal error: %v", err)
		failure(c, http.StatusInternalServerError, "internal server error")
	}
}
