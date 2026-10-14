package main

import (
	"math"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	ID    uint   `json:"id" gorm:"primaryKey"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type PaginationRequest struct {
	Page  int `form:"page,default=1" binding:"gte=1"`
	Limit int `form:"limit,default=10" binding:"gte=1,lte=100"`
}

type PaginationMeta struct {
	Page       int   `json:"page"`
	Limit      int   `json:"limit"`
	Total      int64 `json:"total"`
	TotalPages int   `json:"total_pages"`
}

type UserListResponse struct {
	Data       []User         `json:"data"`
	Pagination PaginationMeta `json:"pagination"`
}

func main() {
	db, err := gorm.Open(sqlite.Open("app.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	if err := db.AutoMigrate(&User{}); err != nil {
		panic(err)
	}

	r := gin.Default()

	r.GET("/users", func(c *gin.Context) {
		var req PaginationRequest

		if err := c.ShouldBindQuery(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "page must be >= 1 and limit must be between 1 and 100",
			})
			return
		}

		offset := (req.Page - 1) * req.Limit

		var total int64
		if err := db.Model(&User{}).Count(&total).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to count users",
			})
			return
		}

		var users []User
		if err := db.Order("id ASC").
			Limit(req.Limit).
			Offset(offset).
			Find(&users).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to fetch users",
			})
			return
		}

		totalPages := int(math.Ceil(float64(total) / float64(req.Limit)))

		c.JSON(http.StatusOK, UserListResponse{
			Data: users,
			Pagination: PaginationMeta{
				Page:       req.Page,
				Limit:      req.Limit,
				Total:      total,
				TotalPages: totalPages,
			},
		})
	})

	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}

//GET /users?page=2&limit=10
