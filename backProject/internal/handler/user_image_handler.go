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

type UserImageHandler struct {
	service   *service.ImageService
	uploadDir string
}

func NewUserImageHandler(service *service.ImageService, uploadDir string) *UserImageHandler {
	return &UserImageHandler{service: service, uploadDir: uploadDir}
}

// GET /api/user/images - 列出当前会员自己的图片
func (h *UserImageHandler) ListMyImages(c *gin.Context) {
	userID, _ := c.Get("user_id")
	category := c.Query("category")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))

	images, total, err := h.service.GetUserList(userID.(uint), category, page, pageSize)
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

// GET /api/user/images/categories - 获取当前会员图片的分类列表
func (h *UserImageHandler) GetMyCategories(c *gin.Context) {
	userID, _ := c.Get("user_id")
	categories, err := h.service.GetUserCategories(userID.(uint))
	if err != nil {
		response.ServerError(c, "获取分类列表失败")
		return
	}
	response.Success(c, categories)
}

// POST /api/user/images - 上传图片（归属当前会员）
func (h *UserImageHandler) Upload(c *gin.Context) {
	userID, _ := c.Get("user_id")

	// Limit: max 60 images per user
	count, err := h.service.CountByUser(userID.(uint))
	if err != nil {
		response.ServerError(c, "查询图片数量失败")
		return
	}
	if count >= 60 {
		response.BadRequest(c, "图片数量已达上限（最多60张）")
		return
	}

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

	// Validate file size (500KB max)
	if file.Size > 500*1024 {
		response.BadRequest(c, "文件大小不能超过500KB")
		return
	}

	// Generate unique filename
	filename := fmt.Sprintf("%d%s", time.Now().UnixNano(), ext)
	userDir := strconv.FormatUint(uint64(userID.(uint)), 10)
	saveDir := filepath.Join(h.uploadDir, userDir)
	if err := os.MkdirAll(saveDir, 0755); err != nil {
		response.ServerError(c, "创建目录失败")
		return
	}
	savePath := filepath.Join(saveDir, filename)

	if err := c.SaveUploadedFile(file, savePath); err != nil {
		response.ServerError(c, "文件上传失败")
		return
	}

	var req struct {
		Category string `form:"category"`
	}
	c.ShouldBind(&req)
	category := req.Category
	if category == "" {
		category = "default"
	}

	image := &model.Image{
		UserID:   userID.(uint),
		Filename: file.Filename,
		URL:      "/uploads/" + userDir + "/" + filename,
		Category: category,
	}

	if err := h.service.Create(image); err != nil {
		response.ServerError(c, "保存图片记录失败")
		return
	}

	response.Success(c, image)
}

// DELETE /api/user/images/:id - 删除自己的图片（校验属主）
func (h *UserImageHandler) DeleteMyImage(c *gin.Context) {
	userID, _ := c.Get("user_id")
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效的ID")
		return
	}

	if err := h.service.DeleteByUser(userID.(uint), uint(id)); err != nil {
		response.ServerError(c, err.Error())
		return
	}

	response.Success(c, nil)
}
