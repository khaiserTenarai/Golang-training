package handler

import (
	"employee-microservices/auth/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

// AuthHandler is the HTTP/controller layer. Task 6.
type AuthHandler struct{ service *service.AuthService }

func NewAuthHandler(s *service.AuthService) *AuthHandler { return &AuthHandler{service: s} }
func (h *AuthHandler) RegisterRoutes(r *gin.Engine)      { r.POST("/login", h.Login) }
func (h *AuthHandler) Login(c *gin.Context) {
	var input struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON"})
		return
	}
	if err := h.service.Login(input.Username, input.Password); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "demo login successful", "note": "learning-only; no JWT issued"})
}
