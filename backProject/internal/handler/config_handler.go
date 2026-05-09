package handler

import (
	"fmt"

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
	var raw map[string]interface{}
	if err := c.ShouldBindJSON(&raw); err != nil {
		response.BadRequest(c, "参数错误")
		return
	}
	configs := make(map[string]string, len(raw))
	for k, v := range raw {
		if k == "" {
			continue
		}
		configs[k] = fmt.Sprintf("%v", v)
	}
	if err := h.service.UpdateConfigs(configs); err != nil {
		response.ServerError(c, "更新配置失败")
		return
	}
	response.Success(c, nil)
}
