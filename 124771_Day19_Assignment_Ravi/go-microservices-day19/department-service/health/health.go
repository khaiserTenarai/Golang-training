package health

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

// Register exposes liveness (/health) and readiness (/ready) endpoints.
func Register(r *gin.Engine, db *gorm.DB) {
	r.GET("/health", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"status": "ok"}) })
	r.GET("/ready", func(c *gin.Context) {
		sqlDB, err := db.DB()
		if err != nil || sqlDB.PingContext(c.Request.Context()) != nil {
			c.JSON(http.StatusServiceUnavailable, gin.H{"status": "not_ready", "error": "database unavailable"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "ready", "database": "ok"})
	})
}
