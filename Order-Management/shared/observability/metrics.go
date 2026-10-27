package observability

import (
	"fmt"
	"net/http"
	"sync/atomic"
	"time"

	"github.com/gin-gonic/gin"
	"order-management/shared/logger"
)

var requestCount atomic.Uint64
var errorCount atomic.Uint64
var databaseErrorCount atomic.Uint64
var latencyNanos atomic.Uint64
var latencyBuckets = [11]atomic.Uint64{}
var bucketBounds = [11]float64{0.005, 0.01, 0.025, 0.05, 0.1, 0.25, 0.5, 1, 2.5, 5, 10}

// IncDatabaseError records a database failure from repository/database code.
func IncDatabaseError() { databaseErrorCount.Add(1) }

// Middleware records request metrics and structured request logs.
func Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		duration := time.Since(start)
		status := c.Writer.Status()
		requestCount.Add(1)
		latencyNanos.Add(uint64(duration.Nanoseconds()))
		seconds := duration.Seconds()
		for i, bound := range bucketBounds {
			if seconds <= bound {
				latencyBuckets[i].Add(1)
			}
		}
		requestID, _ := c.Get("request_id")
		requestIDText, _ := requestID.(string)
		errText := ""
		if value, exists := c.Get("error"); exists {
			if text, ok := value.(string); ok {
				errText = text
			}
		}
		if errText == "" && len(c.Errors) > 0 {
			errText = c.Errors.String()
		}
		if status >= http.StatusBadRequest {
			errorCount.Add(1)
			if errText == "" {
				errText = http.StatusText(status)
			}
		}
		logger.Get().Info("http_request",
			"request_id", requestIDText,
			"method", c.Request.Method,
			"path", c.Request.URL.Path,
			"status", status,
			"duration", duration.String(),
			"error", errText,
		)
	}
}

// Register exposes liveness, readiness, and Prometheus-compatible metrics.
// readiness receives a function that checks dependencies such as PostgreSQL.
func Register(router *gin.Engine, serviceName string, readiness func() error) {
	live := func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"service": serviceName, "status": "live"})
	}
	router.GET("/health/live", live)
	router.GET("/health/ready", func(c *gin.Context) {
		if readiness != nil {
			if err := readiness(); err != nil {
				c.JSON(http.StatusServiceUnavailable, gin.H{"service": serviceName, "status": "not_ready", "error": err.Error()})
				return
			}
		}
		c.JSON(http.StatusOK, gin.H{"service": serviceName, "status": "ready"})
	})
	router.GET("/metrics", MetricsHandler())
}

func MetricsHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "text/plain; version=0.0.4; charset=utf-8")
		count := requestCount.Load()
		errors := errorCount.Load()
		dbErrors := databaseErrorCount.Load()
		latency := latencyNanos.Load()
		fmt.Fprintf(c.Writer, "# TYPE http_requests_total counter\nhttp_requests_total %d\n", count)
		fmt.Fprintf(c.Writer, "# TYPE http_request_errors_total counter\nhttp_request_errors_total %d\n", errors)
		fmt.Fprintf(c.Writer, "# TYPE database_errors_total counter\ndatabase_errors_total %d\n", dbErrors)
		fmt.Fprintln(c.Writer, "# TYPE http_request_duration_seconds histogram")
		fmt.Fprintf(c.Writer, "http_request_duration_seconds_sum %.9f\n", float64(latency)/float64(time.Second))
		fmt.Fprintf(c.Writer, "http_request_duration_seconds_count %d\n", count)
		for i, bound := range bucketBounds {
			fmt.Fprintf(c.Writer, "http_request_duration_seconds_bucket{le=\"%.3g\"} %d\n", bound, latencyBuckets[i].Load())
		}
		fmt.Fprintf(c.Writer, "http_request_duration_seconds_bucket{le=\"+Inf\"} %d\n", count)
		c.Status(http.StatusOK)
	}
}
