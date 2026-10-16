package controller

import (
	"employee-management/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AuthController struct{ service service.AuthService }

func NewAuthController(s service.AuthService) *AuthController { return &AuthController{s} }

type loginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Login: POST /login  {"username":"admin","password":"admin123"}  ->  {"token":"..."}
func (c *AuthController) Login(ctx *gin.Context) {
	var req loginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result, err := c.service.Login(req.Username, req.Password)
	if err != nil {
		fail(ctx, err, "user")
		return
	}
	ctx.JSON(http.StatusOK, result)
}
