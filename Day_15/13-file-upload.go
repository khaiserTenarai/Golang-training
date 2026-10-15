package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

var reader = bufio.NewReader(os.Stdin)

func ask(question string) string {
	fmt.Print(question)
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text)
}

func main() {
	port := ask("Enter port number (press Enter for 8080): ")
	if port == "" {
		port = "8080"
	}

	folder := ask("Which folder should uploaded files be saved in? ")
	if folder == "" {
		fmt.Println("A folder name is needed.")
		return
	}
	if err := os.MkdirAll(folder, 0o755); err != nil {
		fmt.Println("Could not create the folder:", err)
		return
	}

	maxMB, err := strconv.Atoi(ask("Maximum file size in MB: "))
	if err != nil || maxMB < 1 {
		fmt.Println("Please enter a whole number of 1 or more.")
		return
	}
	maxBytes := int64(maxMB) * 1024 * 1024

	router := gin.New()

	router.POST("/upload", func(c *gin.Context) {
		c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, maxBytes)

		file, err := c.FormFile("file")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "send a file in a form field called 'file' (and keep it under the size limit)"})
			return
		}

		name := filepath.Base(file.Filename)
		savePath := filepath.Join(folder, name)

		if err := c.SaveUploadedFile(file, savePath); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "could not save the file"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "file uploaded", "name": name, "size_bytes": file.Size})
	})

	if err := router.Run(":" + port); err != nil {
		fmt.Println("Could not start the server:", err)
	}
}
