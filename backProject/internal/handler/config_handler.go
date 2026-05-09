package handler

import (
	"github.com/gin-gonic/gin"
	"zyqHome/backProject/internal/service"
	"zyqHome/backProject/pkg/response"
)

type ConfigHandler struct {
	service *service.ConfigService
}

func NewConfigHandler(service *service.ConfigService) *ConfigHandler {
	return &ConfigHandler{service: service}
}

func (h *ConfigHandler) GetPublicConfig(c *gin.Context) {
	cfg, err := h.service.GetPublicConfig()
	if err != nil {
		response.ServerError(c, "获取配置失败")
		return
	}
	response.Success(c, cfg)
}

func (h *ConfigHandler) GetAllConfigs(c *gin.Context) {
	configs, err := h.service.GetAllConfigs()
	if err != nil {
		response.ServerError(c, "获取配置失败")
		return
	}
	response.Success(c, configs)
}

func (h *ConfigHandler) UpdateConfigs(c *gin.Context) {
	var configs map[string]string
	if err := c.ShouldBindJSON(&configs); err != nil {
		response.BadRequest(c, "参数错误")
		return
	}
	if err := h.service.UpdateConfigs(configs); err != nil {
		response.ServerError(c, "更新配置失败")
		return
	}
	response.Success(c, nil)
}
