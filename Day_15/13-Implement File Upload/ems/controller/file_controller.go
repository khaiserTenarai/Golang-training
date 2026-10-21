package controller

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
)

const maxFileSize int64 = 5 << 20 // 5 MB

// UploadEmployeeFile handles employee document uploads.
func UploadEmployeeFile(c *gin.Context) {
	// Get the uploaded file from the multipart form.
	header, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "File is required. Use the form field 'file'.",
		})
		return
	}

	// Validate file size.
	if header.Size <= 0 || header.Size > maxFileSize {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "File size must be between 1 byte and 5 MB.",
		})
		return
	}

	// Allow only specific file extensions.
	ext := strings.ToLower(filepath.Ext(header.Filename))
	allowedExtensions := map[string]bool{
		".pdf":  true,
		".jpg":  true,
		".jpeg": true,
		".png":  true,
	}

	if !allowedExtensions[ext] {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Only PDF, JPG, JPEG, and PNG files are allowed.",
		})
		return
	}

	// Open the uploaded file.
	src, err := header.Open()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Unable to read uploaded file.",
		})
		return
	}
	defer src.Close()

	// Detect actual content type instead of trusting the filename.
	buffer := make([]byte, 512)
	n, err := io.ReadFull(src, buffer)
	if err != nil && err != io.EOF && err != io.ErrUnexpectedEOF {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Unable to read file content.",
		})
		return
	}

	contentType := http.DetectContentType(buffer[:n])

	allowedTypes := map[string]bool{
		".pdf":  contentType == "application/pdf",
		".jpg":  contentType == "image/jpeg",
		".jpeg": contentType == "image/jpeg",
		".png":  contentType == "image/png",
	}

	if !allowedTypes[ext] {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "File content does not match an allowed file type.",
		})
		return
	}

	// Generate a random filename to avoid collisions and unsafe names.
	randomBytes := make([]byte, 16)
	if _, err := rand.Read(randomBytes); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Unable to generate file name.",
		})
		return
	}

	filename := hex.EncodeToString(randomBytes) + ext
	uploadDir := "uploads"

	if err := os.MkdirAll(uploadDir, 0750); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Unable to create upload directory.",
		})
		return
	}

	filePath := filepath.Join(uploadDir, filename)

	// O_EXCL prevents accidentally overwriting an existing file.
	dst, err := os.OpenFile(
		filePath,
		os.O_WRONLY|os.O_CREATE|os.O_EXCL,
		0640,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Unable to save uploaded file.",
		})
		return
	}

	// Copy the already-checked header bytes, then the remaining content.
	_, copyErr := dst.Write(buffer[:n])
	if copyErr == nil {
		_, copyErr = io.Copy(dst, src)
	}
	closeErr := dst.Close()

	if copyErr != nil || closeErr != nil {
		_ = os.Remove(filePath)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Unable to save uploaded file.",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message":  "File uploaded successfully",
		"filename": filename,
		"size":     header.Size,
		"path":     fmt.Sprintf("%s/%s", uploadDir, filename),
	})
}
