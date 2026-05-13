package handler

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"zyqHome/backProject/internal/service"
	"zyqHome/backProject/pkg/response"
)

type AuthHandler struct {
	authService *service.AuthService
}

func NewAuthHandler(authService *service.AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "请输入用户名和密码")
		return
	}

	ip := c.ClientIP()
	result, err := h.authService.Login(req.Username, req.Password, ip)
	if err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	response.Success(c, result)
}

func (h *AuthHandler) Register(c *gin.Context) {
	var req service.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "请输入正确的邮箱和密码")
		return
	}

	if req.Email == "" || req.Password == "" {
		response.BadRequest(c, "邮箱和密码不能为空")
		return
	}

	result, err := h.authService.Register(req)
	if err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	response.Success(c, result)
}

func (h *AuthHandler) Me(c *gin.Context) {
	userID, _ := c.Get("user_id")
	user, err := h.authService.GetCurrentUser(userID.(uint))
	if err != nil {
		response.NotFound(c, "用户不存在")
		return
	}
	response.Success(c, user)
}

func (h *AuthHandler) MemberList(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))

	users, total, err := h.authService.GetMemberList(page, pageSize)
	if err != nil {
		response.ServerError(c, "获取会员列表失败")
		return
	}

	response.Success(c, gin.H{
		"items":     users,
		"total":     total,
		"page":      page,
		"page_size": pageSize,
	})
}

func (h *AuthHandler) MemberDetail(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "参数错误")
		return
	}

	user, err := h.authService.GetMemberDetail(uint(id))
	if err != nil {
		response.NotFound(c, err.Error())
		return
	}

	response.Success(c, user)
}

func (h *AuthHandler) UpdateMemberStatus(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "参数错误")
		return
	}

	var req struct {
		Status *int `json:"status" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "请提供状态值")
		return
	}

	if err := h.authService.UpdateMemberStatus(uint(id), *req.Status); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	response.Success(c, nil)
}
