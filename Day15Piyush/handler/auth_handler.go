// Auth handlers: register + login (JWT token generation)

package handler

import (
	"net/http"
	"sync"

	"advanced_gin/middleware"
	"advanced_gin/models"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	mu    sync.RWMutex
	users map[string]models.User
	next  int
}

func NewAuthHandler() *AuthHandler {
	h := &AuthHandler{users: make(map[string]models.User), next: 1}
	// Seed users
	seed := []models.User{
		{Username: "admin", Password: "admin123", Role: "admin"},
		{Username: "manager", Password: "manager123", Role: "manager"},
		{Username: "piyush", Password: "piyush123", Role: "user"},
	}
	for _, u := range seed {
		u.ID = h.next
		h.users[u.Username] = u
		h.next++
	}
	return h
}

func (h *AuthHandler) Register(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": err.Error()})
		return
	}
	if user.Role == "" {
		user.Role = "user"
	}
	h.mu.Lock()
	if _, exists := h.users[user.Username]; exists {
		h.mu.Unlock()
		c.JSON(http.StatusConflict, gin.H{"success": false, "message": "Username taken"})
		return
	}
	user.ID = h.next
	h.next++
	h.users[user.Username] = user
	h.mu.Unlock()

	token, _ := middleware.GenerateJWT(user.ID, user.Username, user.Role)
	c.JSON(http.StatusCreated, gin.H{
		"success": true, "message": "Registered", "token": token,
		"user": gin.H{"id": user.ID, "username": user.Username, "role": user.Role},
	})
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req models.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": err.Error()})
		return
	}
	h.mu.RLock()
	user, exists := h.users[req.Username]
	h.mu.RUnlock()

	if !exists || user.Password != req.Password {
		c.JSON(http.StatusUnauthorized, gin.H{"success": false, "message": "Invalid credentials"})
		return
	}

	token, _ := middleware.GenerateJWT(user.ID, user.Username, user.Role)
	c.JSON(http.StatusOK, gin.H{
		"success": true, "message": "Login successful", "token": token,
		"user": gin.H{"id": user.ID, "username": user.Username, "role": user.Role},
	})
}

func (h *AuthHandler) Profile(c *gin.Context) {
	username, _ := c.Get("username")
	role, _ := c.Get("role")
	userID, _ := c.Get("user_id")
	c.JSON(http.StatusOK, gin.H{
		"success": true, "data": gin.H{"id": userID, "username": username, "role": role},
	})
}
