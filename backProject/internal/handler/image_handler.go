package handler

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"

	"zyqHome/backProject/internal/model"
	"zyqHome/backProject/internal/service"
	"zyqHome/backProject/pkg/response"
)

type ImageHandler struct {
	service   *service.ImageService
	uploadDir string
}

func NewImageHandler(service *service.ImageService, uploadDir string) *ImageHandler {
	return &ImageHandler{service: service, uploadDir: uploadDir}
}

func (h *ImageHandler) GetList(c *gin.Context) {
	category := c.Query("category")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))

	images, total, err := h.service.GetList(category, page, pageSize)
	if err != nil {
		response.ServerError(c, "获取图片列表失败")
		return
	}

	response.Success(c, gin.H{
		"items": images,
		"total": total,
		"page":  page,
	})
}

func (h *ImageHandler) GetCategories(c *gin.Context) {
	categories, err := h.service.GetCategories()
	if err != nil {
		response.ServerError(c, "获取分类列表失败")
		return
	}
	response.Success(c, categories)
}

type UploadImageReq struct {
	Category string `form:"category"`
}

func (h *ImageHandler) Upload(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		response.BadRequest(c, "请选择要上传的文件")
		return
	}

	// Validate file type
	ext := strings.ToLower(filepath.Ext(file.Filename))
	allowedExts := map[string]bool{".jpg": true, ".jpeg": true, ".png": true, ".gif": true, ".webp": true, ".svg": true, ".ico": true}
	if !allowedExts[ext] {
		response.BadRequest(c, "不支持的文件格式，仅支持 jpg/png/gif/webp/svg/ico")
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

	var req UploadImageReq
	c.ShouldBind(&req)
	category := req.Category
	if category == "" {
		category = "default"
	}

	image := &model.Image{
		Filename: file.Filename,
		URL:      "/uploads/" + filename,
		Category: category,
	}

	if err := h.service.Create(image); err != nil {
		response.ServerError(c, "保存图片记录失败")
		return
	}

	response.Success(c, image)
}

func (h *ImageHandler) Update(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效的ID")
		return
	}

	var req struct {
		Category string `json:"category"`
		Filename string `json:"filename"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数错误")
		return
	}

	image, err := h.service.GetByID(uint(id))
	if err != nil {
		response.NotFound(c, "图片不存在")
		return
	}

	image.Category = req.Category
	if req.Filename != "" {
		image.Filename = req.Filename
	}
	if err := h.service.Update(image); err != nil {
		response.ServerError(c, "更新图片失败")
		return
	}

	response.Success(c, image)
}

func (h *ImageHandler) Crop(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效的ID")
		return
	}

	image, err := h.service.GetByID(uint(id))
	if err != nil {
		response.NotFound(c, "图片不存在")
		return
	}

	file, err := c.FormFile("file")
	if err != nil {
		response.BadRequest(c, "请上传裁剪后的图片文件")
		return
	}

	// Validate file type
	ext := strings.ToLower(filepath.Ext(file.Filename))
	allowedExts := map[string]bool{".jpg": true, ".jpeg": true, ".png": true, ".gif": true, ".webp": true, ".svg": true, ".ico": true}
	if !allowedExts[ext] {
		response.BadRequest(c, "不支持的文件格式，仅支持 jpg/png/gif/webp/svg/ico")
		return
	}

	// Validate file size (10MB max)
	if file.Size > 10*1024*1024 {
		response.BadRequest(c, "文件大小不能超过10MB")
		return
	}

	// Delete old physical file
	oldFilename := filepath.Base(image.URL)
	oldPath := filepath.Join(h.uploadDir, oldFilename)
	os.Remove(oldPath)

	// Save new file
	newFilename := fmt.Sprintf("%d%s", time.Now().UnixNano(), ext)
	savePath := filepath.Join(h.uploadDir, newFilename)
	if err := c.SaveUploadedFile(file, savePath); err != nil {
		response.ServerError(c, "保存裁剪图片失败")
		return
	}

	// Update database record
	image.URL = "/uploads/" + newFilename

	category := c.PostForm("category")
	if category != "" {
		image.Category = category
	}

	if filename := c.PostForm("filename"); filename != "" {
		image.Filename = filename
	}

	if err := h.service.Update(image); err != nil {
		response.ServerError(c, "更新图片记录失败")
		return
	}

	response.Success(c, image)
}

func (h *ImageHandler) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效的ID")
		return
	}

	if err := h.service.Delete(uint(id)); err != nil {
		response.ServerError(c, "删除图片失败")
		return
	}

	response.Success(c, nil)
}
