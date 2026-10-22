package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"gin-layered-api/controllers"
	"gin-layered-api/middlewares"
	"gin-layered-api/models"
	"gin-layered-api/repositories"
	"gin-layered-api/services"
	"gin-layered-api/utils"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupTestRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	userRepo := repositories.NewUserRepository()
	productRepo := repositories.NewProductRepository()

	authService := services.NewAuthService(userRepo)
	productService := services.NewProductService(productRepo)

	authController := controllers.NewAuthController(authService)
	productController := controllers.NewProductController(productService)

	r := gin.New()
	r.POST("/login", authController.Login)

	api := r.Group("/api")
	api.Use(middlewares.AuthMiddleware())
	{
		api.GET("/products", productController.GetProducts)
		api.POST("/products", middlewares.RBACMiddleware(models.RoleAdmin), productController.CreateProduct)
	}

	return r
}

func TestLoginValidation(t *testing.T) {
	router := setupTestRouter()

	body, _ := json.Marshal(map[string]string{
		"username": "alice",
	})

	req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestSuccessfulLogin(t *testing.T) {
	router := setupTestRouter()

	body, _ := json.Marshal(map[string]string{
		"username": "alice",
		"password": "password123",
	})

	req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "token")
}

func TestGetProductsPagination(t *testing.T) {
	router := setupTestRouter()
	token, _ := utils.GenerateToken("alice", models.RoleUser)

	req, _ := http.NewRequest("GET", "/api/products?page=1&limit=2", nil)
	req.Header.Set("Authorization", "Bearer "+token)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), `"limit":2`)
	assert.Contains(t, w.Body.String(), `"page":1`)
}

func TestRBAC_UserCannotCreateProduct(t *testing.T) {
	router := setupTestRouter()
	userToken, _ := utils.GenerateToken("alice", models.RoleUser)

	body, _ := json.Marshal(map[string]interface{}{
		"name":  "Desk",
		"price": 150.00,
	})

	req, _ := http.NewRequest("POST", "/api/products", bytes.NewBuffer(body))
	req.Header.Set("Authorization", "Bearer "+userToken)
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusForbidden, w.Code)
}

func TestRBAC_AdminCanCreateProduct(t *testing.T) {
	router := setupTestRouter()
	adminToken, _ := utils.GenerateToken("bob", models.RoleAdmin)

	body, _ := json.Marshal(map[string]interface{}{
		"name":  "Gaming Chair",
		"price": 250.00,
	})

	req, _ := http.NewRequest("POST", "/api/products", bytes.NewBuffer(body))
	req.Header.Set("Authorization", "Bearer "+adminToken)
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Contains(t, w.Body.String(), "Gaming Chair")
}
