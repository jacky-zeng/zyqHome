package handler

import (
	"fmt"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
	"zyqHome/backProject/pkg/response"
)

type UploadHandler struct {
	uploadDir string
}

func NewUploadHandler(uploadDir string) *UploadHandler {
	return &UploadHandler{uploadDir: uploadDir}
}

func (h *UploadHandler) Upload(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		response.BadRequest(c, "请选择要上传的文件")
		return
	}

	// Validate file type
	ext := filepath.Ext(file.Filename)
	allowedExts := map[string]bool{".jpg": true, ".jpeg": true, ".png": true, ".gif": true, ".webp": true}
	if !allowedExts[ext] {
		response.BadRequest(c, "不支持的文件格式，仅支持 jpg/png/gif/webp")
		return
	}

	// Validate file size (10MB max)
	if file.Size > 10*1024*1024 {
		response.BadRequest(c, "文件大小不能超过10MB")
		return
	}

	// Generate unique filename
	filename := fmt.Sprintf("%d%s", time.Now().UnixNano(), ext)
	savePath := filepath.Join(h.uploadDir, filename)

	if err := c.SaveUploadedFile(file, savePath); err != nil {
		response.ServerError(c, "文件上传失败")
		return
	}

	response.Success(c, gin.H{
		"url": "/uploads/" + filename,
	})
}
