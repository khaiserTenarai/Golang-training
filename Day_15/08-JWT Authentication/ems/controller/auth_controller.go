package controller

import (
	"net/http"

	"ems/auth"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	auth *auth.AuthService
}

func NewAuthController(a *auth.AuthService) *AuthController {
	return &AuthController{auth: a}
}

// POST /login
func (c *AuthController) Login(ctx *gin.Context) {
	var input struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(400, gin.H{"error": "username and password are required"})
		return
	}

	token, err := c.auth.Login(input.Username, input.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "invalid username or password",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"token":     token,
		"tokenType": "Bearer",
	})
}
