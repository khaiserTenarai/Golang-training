package controller

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"task13_authentication_system/middleware"
	"task13_authentication_system/models"
	"task13_authentication_system/service"

	"github.com/gorilla/mux"
)

type AuthController struct {
	service *service.AuthService
}

func NewAuthController(svc *service.AuthService) *AuthController {
	return &AuthController{service: svc}
}

func (c *AuthController) Register(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		sendJSON(w, http.StatusBadRequest, models.Response{Message: "Invalid request body"})
		return
	}
	created, err := c.service.Register(user)
	if err != nil {
		sendJSON(w, http.StatusBadRequest, models.Response{Message: err.Error()})
		return
	}
	sendJSON(w, http.StatusCreated, models.Response{Message: "User registered successfully", Data: created})
}

func (c *AuthController) Login(w http.ResponseWriter, r *http.Request) {
	var req models.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		sendJSON(w, http.StatusBadRequest, models.Response{Message: "Invalid request body"})
		return
	}
	token, user, err := c.service.Login(req.Username, req.Password)
	if err != nil {
		sendJSON(w, http.StatusUnauthorized, models.LoginResponse{Message: err.Error()})
		return
	}
	sendJSON(w, http.StatusOK, models.LoginResponse{Message: "Login successful", Token: token, User: &user})
}

func (c *AuthController) Logout(w http.ResponseWriter, r *http.Request) {
	token := strings.TrimPrefix(r.Header.Get("Authorization"), "Bearer ")
	if err := c.service.Logout(token); err != nil {
		sendJSON(w, http.StatusInternalServerError, models.Response{Message: "Logout failed"})
		return
	}
	sendJSON(w, http.StatusOK, models.Response{Message: "Logged out successfully"})
}

func (c *AuthController) Profile(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(middleware.UserIDKey).(int)
	user, err := c.service.Repo.GetByID(userID)
	if err != nil {
		sendJSON(w, http.StatusNotFound, models.Response{Message: "User not found"})
		return
	}
	sendJSON(w, http.StatusOK, models.Response{Message: "Profile", Data: user})
}

func (c *AuthController) ListUsers(w http.ResponseWriter, r *http.Request) {
	users, err := c.service.Repo.GetAll()
	if err != nil {
		sendJSON(w, http.StatusInternalServerError, models.Response{Message: "Failed to fetch users"})
		return
	}
	sendJSON(w, http.StatusOK, models.Response{Message: "Users list", Data: users})
}

func (c *AuthController) DeleteUser(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	if err := c.service.Repo.DeleteUser(id); err != nil {
		sendJSON(w, http.StatusInternalServerError, models.Response{Message: "Failed to delete user"})
		return
	}
	sendJSON(w, http.StatusOK, models.Response{Message: "User deleted"})
}

func (c *AuthController) UpdateRole(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	var body struct {
		Role string `json:"role"`
	}
	json.NewDecoder(r.Body).Decode(&body)
	if err := c.service.Repo.UpdateRole(id, body.Role); err != nil {
		sendJSON(w, http.StatusInternalServerError, models.Response{Message: "Failed to update role"})
		return
	}
	sendJSON(w, http.StatusOK, models.Response{Message: "Role updated to " + body.Role})
}

func sendJSON(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(payload)
}
