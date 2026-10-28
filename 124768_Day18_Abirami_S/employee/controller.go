package employee

import (
    "net/http"
    "time"

    "github.com/gin-gonic/gin"
)

type Controller struct { service *Service }
func NewController(service *Service) *Controller { return &Controller{service: service} }
func (c *Controller) Routes() *gin.Engine {
    r := gin.New()
    r.Use(gin.Recovery())
    r.GET("/health", func(ctx *gin.Context) { ctx.JSON(http.StatusOK, gin.H{"service":"employee", "status":"UP"}) })
    r.GET("/employees", c.list)
    r.POST("/employees", c.create)
    return r
}
func (c *Controller) list(ctx *gin.Context) {
    requestCtx, cancel := contextWithTimeout(ctx, 3*time.Second)
    defer cancel()
    employees, err := c.service.List(requestCtx)
    if err != nil { ctx.JSON(http.StatusInternalServerError, gin.H{"error":"could not list employees"}); return }
    ctx.JSON(http.StatusOK, employees)
}
func (c *Controller) create(ctx *gin.Context) {
    var input CreateInput
    if err := ctx.ShouldBindJSON(&input); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error":"invalid JSON or required fields missing"})
        return
    }
    requestCtx, cancel := contextWithTimeout(ctx, 3*time.Second)
    defer cancel()
    employee, err := c.service.Create(requestCtx, input)
    if err != nil { ctx.JSON(http.StatusBadRequest, gin.H{"error":err.Error()}); return }
    ctx.JSON(http.StatusCreated, employee)
}
