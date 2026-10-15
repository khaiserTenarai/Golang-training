package controller

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"secure-banking-api/middleware"
	"secure-banking-api/model"
	"secure-banking-api/service"

	"github.com/gin-gonic/gin"
)

const maxUploadSize = 2 << 20 

var allowedExt = map[string]bool{".pdf": true, ".png": true, ".jpg": true, ".jpeg": true}

type AccountController struct {
	svc       service.AccountService
	uploadDir string
}

func NewAccountController(svc service.AccountService, uploadDir string) *AccountController {
	return &AccountController{svc: svc, uploadDir: uploadDir}
}

func (h *AccountController) Create(c *gin.Context) {
	var req model.CreateAccountRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		bindError(c, err)
		return
	}
	acc, err := h.svc.Create(c.Request.Context(), c.GetInt64(middleware.CtxUserID), req)
	if err != nil {
		handleError(c, err)
		return
	}
	success(c, http.StatusCreated, acc)
}

func (h *AccountController) Get(c *gin.Context) {
	id, ok := parseID(c)
	if !ok {
		return
	}
	acc, err := h.svc.GetByID(c.Request.Context(), id, c.GetInt64(middleware.CtxUserID), c.GetString(middleware.CtxRole))
	if err != nil {
		handleError(c, err)
		return
	}
	success(c, http.StatusOK, acc)
}


func (h *AccountController) ListMine(c *gin.Context) {
	q, ok := pageQuery(c)
	if !ok {
		return
	}
	res, err := h.svc.ListMine(c.Request.Context(), c.GetInt64(middleware.CtxUserID), q)
	if err != nil {
		handleError(c, err)
		return
	}
	success(c, http.StatusOK, res)
}


func (h *AccountController) ListAll(c *gin.Context) {
	q, ok := pageQuery(c)
	if !ok {
		return
	}
	res, err := h.svc.ListAll(c.Request.Context(), q)
	if err != nil {
		handleError(c, err)
		return
	}
	success(c, http.StatusOK, res)
}


func (h *AccountController) UploadDocument(c *gin.Context) {
	id, ok := parseID(c)
	if !ok {
		return
	}
	ctx := c.Request.Context()

	
	if _, err := h.svc.GetByID(ctx, id, c.GetInt64(middleware.CtxUserID), c.GetString(middleware.CtxRole)); err != nil {
		handleError(c, err)
		return
	}

	file, err := c.FormFile("document")
	if err != nil {
		failure(c, http.StatusBadRequest, "form field 'document' is required")
		return
	}
	if file.Size > maxUploadSize {
		failure(c, http.StatusBadRequest, "file too large (max 2 MB)")
		return
	}
	ext := strings.ToLower(filepath.Ext(file.Filename))
	if !allowedExt[ext] {
		failure(c, http.StatusBadRequest, "only pdf, png, jpg, jpeg files are allowed")
		return
	}

	if err := os.MkdirAll(h.uploadDir, 0o755); err != nil {
		handleError(c, err)
		return
	}
	
	dst := filepath.Join(h.uploadDir, fmt.Sprintf("acc%d_%d%s", id, time.Now().UnixNano(), ext))
	if err := c.SaveUploadedFile(file, dst); err != nil {
		handleError(c, err)
		return
	}
	if err := h.svc.AttachDocument(ctx, id, dst); err != nil {
		handleError(c, err)
		return
	}
	success(c, http.StatusOK, gin.H{"document_path": dst})
}

func parseID(c *gin.Context) (int64, bool) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil || id <= 0 {
		failure(c, http.StatusBadRequest, "invalid id")
		return 0, false
	}
	return id, true
}

func pageQuery(c *gin.Context) (model.PageQuery, bool) {
	var q model.PageQuery
	if err := c.ShouldBindQuery(&q); err != nil {
		bindError(c, err)
		return q, false
	}
	q.Normalize()
	return q, true
}
