package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/upload", func(c *gin.Context) {
		f, e := c.FormFile("file")
		if e != nil {
			c.JSON(400, gin.H{"error": "File is required"})
			return
		}
		if e = c.SaveUploadedFile(f, "./"+f.Filename); e != nil {
			c.JSON(500, gin.H{"error": "Unable to save file"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "File uploaded successfully", "filename": f.Filename})
	})
	r.Run(":8080")
}
