package controller

import (
	"encoding/json"
	"net/http"
	"strings"
	"task15_employee_management_capstone/middleware"
	"task15_employee_management_capstone/models"
	"task15_employee_management_capstone/service"
)

type AuthController struct {
	svc service.AuthService
}

func NewAuthController(svc service.AuthService) *AuthController {
	return &AuthController{svc: svc}
}

func (c *AuthController) Register(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		sendJSON(w, http.StatusBadRequest, models.Response{Message: "Invalid request body"})
		return
	}
	created, err := c.svc.Register(user)
	if err != nil {
		sendJSON(w, http.StatusBadRequest, models.Response{Message: err.Error()})
		return
	}
	sendJSON(w, http.StatusCreated, models.Response{Message: "User registered", Data: created})
}

func (c *AuthController) Login(w http.ResponseWriter, r *http.Request) {
	var req models.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		sendJSON(w, http.StatusBadRequest, models.Response{Message: "Invalid request body"})
		return
	}
	token, user, err := c.svc.Login(req.Username, req.Password)
	if err != nil {
		sendJSON(w, http.StatusUnauthorized, models.LoginResponse{Message: err.Error()})
		return
	}
	sendJSON(w, http.StatusOK, models.LoginResponse{Message: "Login successful", Token: token, User: &user})
}

func (c *AuthController) Logout(w http.ResponseWriter, r *http.Request) {
	token := strings.TrimPrefix(r.Header.Get("Authorization"), "Bearer ")
	c.svc.Logout(token)
	sendJSON(w, http.StatusOK, models.Response{Message: "Logged out"})
}

func (c *AuthController) Profile(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(middleware.UserIDKey).(int)
	user, err := c.svc.GetUserByID(userID)
	if err != nil {
		sendJSON(w, http.StatusNotFound, models.Response{Message: "User not found"})
		return
	}
	sendJSON(w, http.StatusOK, models.Response{Message: "Profile", Data: user})
}
