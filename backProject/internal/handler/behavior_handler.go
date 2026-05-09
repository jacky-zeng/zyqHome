package handler

import (
	"github.com/gin-gonic/gin"
	"zyqHome/backProject/internal/service"
	"zyqHome/backProject/pkg/response"
)

type BehaviorHandler struct {
	service *service.BehaviorService
}

func NewBehaviorHandler(service *service.BehaviorService) *BehaviorHandler {
	return &BehaviorHandler{service: service}
}

type TrackRequest struct {
	Action string `json:"action" binding:"required"`
	Target string `json:"target"`
}

func (h *BehaviorHandler) Track(c *gin.Context) {
	var req TrackRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数错误")
		return
	}

	ip := c.ClientIP()
	userAgent := c.GetHeader("User-Agent")
	referer := c.GetHeader("Referer")

	if err := h.service.Track(ip, req.Action, req.Target, userAgent, referer); err != nil {
		response.ServerError(c, "记录失败")
		return
	}
	response.Success(c, nil)
}

func (h *BehaviorHandler) Dashboard(c *gin.Context) {
	// This will be used by the admin dashboard
	// For now return basic stats
	response.Success(c, nil)
}
