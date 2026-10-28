package salary

import (
    "context"
    "net/http"
    "strconv"
    "time"
    "github.com/gin-gonic/gin"
)

type Controller struct { service *Service }
func NewController(service *Service) *Controller { return &Controller{service:service} }
func (c *Controller) Routes() *gin.Engine {
    r := gin.New()
    r.Use(gin.Recovery())
    r.GET("/health", func(ctx *gin.Context) { ctx.JSON(http.StatusOK, gin.H{"service":"salary","status":"UP"}) })
    r.GET("/salaries/:employeeId", c.list)
    return r
}
func (c *Controller) list(ctx *gin.Context) {
    id, err := strconv.ParseInt(ctx.Param("employeeId"), 10, 64)
    if err != nil || id <= 0 { ctx.JSON(http.StatusBadRequest, gin.H{"error":"invalid employee id"}); return }
    requestCtx, cancel := context.WithTimeout(ctx.Request.Context(), 3*time.Second); defer cancel()
    items, err := c.service.ListByEmployee(requestCtx, id)
    if err != nil { ctx.JSON(http.StatusInternalServerError, gin.H{"error":"could not list salaries"}); return }
    ctx.JSON(http.StatusOK, items)
}
