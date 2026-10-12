package main

import "github.com/gin-gonic/gin"

func main() {
    router := gin.Default()
    // Serve files from the local "static" folder at /static.
    router.Static("/static", "./static")
    router.GET("/", func(c *gin.Context) {
        c.File("./static/index.html")
    })
    router.Run(":10014")
}
