package controller

import (
	"net/http"

	"secure-banking-api/model"
	"secure-banking-api/service"

	"github.com/gin-gonic/gin"
)

type AuthController struct{ svc service.AuthService }

func NewAuthController(svc service.AuthService) *AuthController {
	return &AuthController{svc: svc}
}

func (h *AuthController) Register(c *gin.Context) {
	var req model.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		bindError(c, err)
		return
	}
	user, err := h.svc.Register(c.Request.Context(), req)
	if err != nil {
		handleError(c, err)
		return
	}
	success(c, http.StatusCreated, user)
}

func (h *AuthController) Login(c *gin.Context) {
	var req model.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		bindError(c, err)
		return
	}
	token, err := h.svc.Login(c.Request.Context(), req)
	if err != nil {
		handleError(c, err)
		return
	}
	success(c, http.StatusOK, gin.H{"token": token})
}
