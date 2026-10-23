package gateway

import (
    "io"
    "net/http"
    "net/http/httputil"
    "net/url"
    "strings"
    "log"

    "github.com/gin-gonic/gin"
    "employee-microservices/discovery"
    "employee-microservices/saga"
)

func NewRouter(registry *discovery.Registry, sagaController *saga.Controller, logger *log.Logger) *gin.Engine {
    r := gin.New()
    r.Use(gin.Recovery(), RequestID())
    r.GET("/health", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"service":"gateway","status":"UP"}) })
    r.POST("/sagas/employee", sagaController.Run)

    r.Any("/employees", proxyTo(registry, "employee", "/employees"))
    r.Any("/employees/*path", proxyTo(registry, "employee", "/employees"))
    r.Any("/departments", proxyTo(registry, "department", "/departments"))
    r.Any("/departments/*path", proxyTo(registry, "department", "/departments"))
    r.Any("/salaries/*path", proxyTo(registry, "salary", "/salaries"))
    r.Any("/salaries", proxyTo(registry, "salary", "/salaries"))
    _ = logger
    return r
}

func proxyTo(registry *discovery.Registry, serviceName, prefix string) gin.HandlerFunc {
    return func(c *gin.Context) {
        address, ok := registry.Resolve(serviceName)
        if !ok {
            c.JSON(http.StatusServiceUnavailable, gin.H{"error":"service unavailable"})
            return
        }
        target, err := url.Parse(address)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error":"invalid service address"})
            return
        }
        proxy := httputil.NewSingleHostReverseProxy(target)
        originalDirector := proxy.Director
        proxy.Director = func(req *http.Request) {
            originalDirector(req)
            req.URL.Path = strings.TrimRight(prefix, "/") + strings.TrimPrefix(c.Request.URL.Path, prefix)
            req.URL.RawQuery = c.Request.URL.RawQuery
            req.Header.Set("X-Request-ID", c.GetHeader("X-Request-ID"))
        }
        proxy.ErrorHandler = func(w http.ResponseWriter, req *http.Request, err error) {
            w.Header().Set("Content-Type", "application/json")
            w.WriteHeader(http.StatusBadGateway)
            _, _ = io.WriteString(w, `{"error":"upstream service unavailable"}`)
        }
        proxy.ServeHTTP(c.Writer, c.Request)
    }
}
