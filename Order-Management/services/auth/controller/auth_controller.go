package controller

import (
	"net/http"
	"strings"

	"order-management/services/auth/dto"
	"order-management/services/auth/service"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	service *service.AuthService
}

func NewAuthController(s *service.AuthService) *AuthController {
	return &AuthController{service: s}
}

func (c *AuthController) Login(ctx *gin.Context) {
	var req dto.LoginRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "username and password are required",
		})
		return
	}

	req.Username = strings.TrimSpace(req.Username)

	token, role, err := c.service.Login(req.Username, req.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, dto.LoginResponse{
		Token:    token,
		Username: req.Username,
		Role:     role,
	})
}
