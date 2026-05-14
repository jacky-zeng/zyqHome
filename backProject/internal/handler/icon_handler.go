package handler

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"zyqHome/backProject/internal/model"
	"zyqHome/backProject/internal/service"
	"zyqHome/backProject/pkg/response"
)

// PublicIconResponse 公开接口返回的图标字段，隐藏内部字段
type PublicIconResponse struct {
	ID    uint   `json:"id"`
	Title string `json:"title"`
	URL   string `json:"url"`
	Icon  string `json:"icon"`
	Color string `json:"color"`
}

func toPublicIcon(icon model.CenterIcon) PublicIconResponse {
	return PublicIconResponse{
		ID:    icon.ID,
		Title: icon.Title,
		URL:   icon.URL,
		Icon:  icon.Icon,
		Color: icon.Color,
	}
}

type IconHandler struct {
	service *service.IconService
}

func NewIconHandler(service *service.IconService) *IconHandler {
	return &IconHandler{service: service}
}

func (h *IconHandler) GetActiveIcons(c *gin.Context) {
	var icons []model.CenterIcon
	var err error

	menuIDStr := c.Query("menu_id")
	if menuIDStr != "" {
		var menuID uint64
		menuID, err = strconv.ParseUint(menuIDStr, 10, 64)
		if err != nil {
			response.BadRequest(c, "无效的 menu_id")
			return
		}
		icons, err = h.service.GetActiveIconsByMenuID(uint(menuID))
	} else {
		icons, err = h.service.GetActiveIcons()
	}
	if err != nil {
		response.ServerError(c, "获取图标失败")
		return
	}

	result := make([]PublicIconResponse, len(icons))
	for i, icon := range icons {
		result[i] = toPublicIcon(icon)
	}
	response.Success(c, result)
}

func (h *IconHandler) GetAllIcons(c *gin.Context) {
	icons, err := h.service.GetAllIcons()
	if err != nil {
		response.ServerError(c, "获取图标失败")
		return
	}
	response.Success(c, icons)
}

func (h *IconHandler) Create(c *gin.Context) {
	var icon model.CenterIcon
	if err := c.ShouldBindJSON(&icon); err != nil {
		response.BadRequest(c, "参数错误")
		return
	}
	if err := h.service.Create(&icon); err != nil {
		response.ServerError(c, "创建图标失败")
		return
	}
	response.Success(c, icon)
}

func (h *IconHandler) Update(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效的ID")
		return
	}

	var icon model.CenterIcon
	if err := c.ShouldBindJSON(&icon); err != nil {
		response.BadRequest(c, "参数错误")
		return
	}
	icon.ID = uint(id)

	if err := h.service.Update(&icon); err != nil {
		response.ServerError(c, "更新图标失败")
		return
	}
	response.Success(c, icon)
}

func (h *IconHandler) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效的ID")
		return
	}
	if err := h.service.Delete(uint(id)); err != nil {
		response.ServerError(c, "删除图标失败")
		return
	}
	response.Success(c, nil)
}

func (h *IconHandler) BatchSort(c *gin.Context) {
	var req SortRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数错误")
		return
	}

	items := make([]model.CenterIcon, len(req.Items))
	for i, item := range req.Items {
		items[i] = model.CenterIcon{
			ID:        item.ID,
			SortOrder: item.SortOrder,
		}
	}

	if err := h.service.BatchUpdateSort(items); err != nil {
		response.ServerError(c, "排序失败")
		return
	}
	response.Success(c, nil)
}
