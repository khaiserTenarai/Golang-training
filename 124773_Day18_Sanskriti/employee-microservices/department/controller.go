package department

import (
    "context"
    "fmt"
    "net/http"
    "time"

    "github.com/gin-gonic/gin"
)

type Controller struct { service *Service }
func NewController(service *Service) *Controller { return &Controller{service:service} }
func (c *Controller) Routes() *gin.Engine {
    r := gin.New()
    r.Use(gin.Recovery())
    r.GET("/health", func(ctx *gin.Context) { ctx.JSON(http.StatusOK, gin.H{"service":"department","status":"UP"}) })
    r.GET("/departments", c.list)
    r.GET("/departments/:id", c.get)
    return r
}
func (c *Controller) list(ctx *gin.Context) {
    requestCtx, cancel := context.WithTimeout(ctx.Request.Context(), 3*time.Second); defer cancel()
    items, err := c.service.List(requestCtx)
    if err != nil { ctx.JSON(http.StatusInternalServerError, gin.H{"error":"could not list departments"}); return }
    ctx.JSON(http.StatusOK, items)
}
func (c *Controller) get(ctx *gin.Context) {
    var id int64
    if _, err := fmt.Sscan(ctx.Param("id"), &id); err != nil || id <= 0 {
        ctx.JSON(http.StatusBadRequest, gin.H{"error":"invalid department id"}); return
    }
    requestCtx, cancel := context.WithTimeout(ctx.Request.Context(), 3*time.Second); defer cancel()
    item, err := c.service.Get(requestCtx, id)
    if err != nil { ctx.JSON(http.StatusNotFound, gin.H{"error":"department not found"}); return }
    ctx.JSON(http.StatusOK, item)
}
