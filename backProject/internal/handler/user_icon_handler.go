package handler

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"zyqHome/backProject/internal/model"
	"zyqHome/backProject/internal/service"
	"zyqHome/backProject/pkg/response"
)

type UserIconHandler struct {
	service *service.UserIconService
}

func NewUserIconHandler(service *service.UserIconService) *UserIconHandler {
	return &UserIconHandler{service: service}
}

// GET /api/user/icons - 获取当前用户的图标
func (h *UserIconHandler) ListMyIcons(c *gin.Context) {
	userID, _ := c.Get("user_id")
	icons, err := h.service.GetUserIcons(userID.(uint))
	if err != nil {
		response.ServerError(c, "获取图标失败")
		return
	}
	if icons == nil {
		icons = []model.UserIcon{}
	}
	response.Success(c, icons)
}

type CreateUserIconRequest struct {
	ImageURL string `json:"image_url" binding:"required"`
	LinkURL  string `json:"link_url" binding:"required"`
}

// POST /api/user/icons - 创建图标
func (h *UserIconHandler) CreateMyIcon(c *gin.Context) {
	var req CreateUserIconRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "请提供图片URL和跳转URL")
		return
	}
	userID, _ := c.Get("user_id")
	icon, err := h.service.Create(userID.(uint), req.ImageURL, req.LinkURL)
	if err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	response.Success(c, icon)
}

// DELETE /api/user/icons/:id - 删除自己的图标
func (h *UserIconHandler) DeleteMyIcon(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效的ID")
		return
	}
	userID, _ := c.Get("user_id")
	if err := h.service.Delete(userID.(uint), uint(id)); err != nil {
		response.ServerError(c, err.Error())
		return
	}
	response.Success(c, nil)
}

// GET /api/admin/users/:id/icons - 管理员获取指定用户的图标
func (h *UserIconHandler) AdminGetUserIcons(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效的用户ID")
		return
	}
	icons, err := h.service.GetUserIcons(uint(id))
	if err != nil {
		response.ServerError(c, "获取图标失败")
		return
	}
	if icons == nil {
		icons = []model.UserIcon{}
	}
	response.Success(c, icons)
}
