package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/upload", func(c *gin.Context) {
		// Get uploaded file
		file, err := c.FormFile("file")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Please select a file",
			})
			return
		}

		// Save file
		err = c.SaveUploadedFile(file, "./"+file.Filename)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Upload failed",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "File uploaded successfully!",
		})
	})

	r.Run(":8080")
}
