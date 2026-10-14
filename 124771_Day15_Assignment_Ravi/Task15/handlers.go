package main

import (
	"net/http"
	"sort"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type RegisterRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type CreateTaskRequest struct {
	Title string `json:"title" binding:"required"`
}

func Register(c *gin.Context) {
	var req RegisterRequest

	if err := c.ShouldBindJSON(&req); err != nil ||
		!validEmail(req.Email) ||
		!validPassword(req.Password) {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "provide a valid email and a password of 8-72 bytes",
		})
		return
	}

	email := strings.ToLower(strings.TrimSpace(req.Email))

	if _, exists := users[email]; exists {
		c.JSON(http.StatusConflict, gin.H{
			"error": "email already registered",
		})
		return
	}

	hash, err := hashPassword(req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "password hashing failed",
		})
		return
	}

	// Public registration always creates a regular user.
	users[email] = User{
		ID:       email,
		Email:    email,
		Password: hash,
		Role:     "user",
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "registered successfully",
		"user": gin.H{
			"id":    email,
			"email": email,
			"role":  "user",
		},
	})
}

func Login(c *gin.Context) {
	var req LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "email and password are required",
		})
		return
	}

	email := strings.ToLower(strings.TrimSpace(req.Email))
	user, exists := users[email]

	if !exists || !checkPassword(user.Password, req.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "invalid email or password",
		})
		return
	}

	token, err := generateToken(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "token generation failed",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

func CreateTask(c *gin.Context) {
	var req CreateTaskRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})
		return
	}

	if err := validateTitle(req.Title); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	task := Task{
		ID:        nextTaskID,
		Title:     strings.TrimSpace(req.Title),
		Completed: false,
		OwnerID:   c.GetString("user_id"),
	}

	tasks[nextTaskID] = task
	nextTaskID++

	c.JSON(http.StatusCreated, task)
}

func ListTasks(c *gin.Context) {
	page, errPage := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, errLimit := strconv.Atoi(c.DefaultQuery("limit", "10"))

	if errPage != nil || errLimit != nil ||
		page < 1 || limit < 1 || limit > 100 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "page must be positive and limit must be between 1 and 100",
		})
		return
	}

	userID := c.GetString("user_id")
	role := c.GetString("role")

	all := make([]Task, 0)

	for _, task := range tasks {
		if role == "admin" || task.OwnerID == userID {
			all = append(all, task)
		}
	}

	sort.Slice(all, func(i, j int) bool {
		return all[i].ID < all[j].ID
	})

	total := len(all)
	start := (page - 1) * limit

	if start >= total {
		c.JSON(http.StatusOK, gin.H{
			"data":  []Task{},
			"page":  page,
			"limit": limit,
			"total": total,
		})
		return
	}

	end := start + limit
	if end > total {
		end = total
	}

	c.JSON(http.StatusOK, gin.H{
		"data":  all[start:end],
		"page":  page,
		"limit": limit,
		"total": total,
	})
}

func DeleteTask(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil || id < 1 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid task ID",
		})
		return
	}

	task, exists := tasks[id]

	if !exists {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "task not found",
		})
		return
	}

	userID := c.GetString("user_id")
	role := c.GetString("role")

	if role != "admin" && task.OwnerID != userID {
		c.JSON(http.StatusForbidden, gin.H{
			"error": "permission denied",
		})
		return
	}

	delete(tasks, id)

	c.JSON(http.StatusOK, gin.H{
		"message": "task deleted successfully",
	})
}

func ListUsers(c *gin.Context) {
	result := make([]gin.H, 0, len(users))

	for _, user := range users {
		result = append(result, gin.H{
			"id":    user.ID,
			"email": user.Email,
			"role":  user.Role,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"data": result,
	})
}
