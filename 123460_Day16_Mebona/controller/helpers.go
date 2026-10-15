package controller

import (
	"net/http"
	"strconv"

	"employee-management/apperror"

	"github.com/gin-gonic/gin"
)

func respondError(ctx *gin.Context, err error) {
	ae := apperror.As(err)
	status := http.StatusInternalServerError
	switch ae.Code {
	case apperror.CodeNotFound:
		status = http.StatusNotFound
	case apperror.CodeBadInput:
		status = http.StatusBadRequest
	case apperror.CodeConflict:
		status = http.StatusConflict
	}
	body := gin.H{"error": ae.Message, "code": ae.Code}
	if ae.Field != "" {
		body["field"] = ae.Field
	}
	ctx.JSON(status, body)
}

func parseID(ctx *gin.Context) (uint, bool) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil || id == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid employee id"})
		return 0, false
	}
	return uint(id), true
}
