package middleware

import (
	"strconv"
	"time"

	"day18/logging"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	httpRequestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests processed.",
		},
		[]string{"method", "endpoint", "status"},
	)

	httpRequestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "Histogram of request latencies in seconds.",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"method", "endpoint"},
	)
)

func init() {
	prometheus.MustRegister(httpRequestsTotal)
	prometheus.MustRegister(httpRequestDuration)
}

func GinLoggingMiddleware(logger *logging.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		traceID := c.GetHeader("X-Trace-ID")
		if traceID == "" {
			traceID = "trace-" + time.Now().Format("20060102150405")
			c.Header("X-Trace-ID", traceID)
		}

		c.Next()

		duration := time.Since(start)
		status := c.Writer.Status()
		endpoint := c.FullPath()
		if endpoint == "" {
			endpoint = "unknown"
		}

		httpRequestsTotal.WithLabelValues(c.Request.Method, endpoint, strconv.Itoa(status)).Inc()
		httpRequestDuration.WithLabelValues(c.Request.Method, endpoint).Observe(duration.Seconds())

		logger.Info("[%s] %s %d %s - TraceID: %s", c.Request.Method, c.Request.URL.Path, status, duration, traceID)
	}
}