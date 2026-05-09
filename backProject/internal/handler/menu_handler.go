package handler

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"zyqHome/backProject/internal/model"
	"zyqHome/backProject/internal/service"
	"zyqHome/backProject/pkg/response"
)

type MenuHandler struct {
	service *service.MenuService
}

func NewMenuHandler(service *service.MenuService) *MenuHandler {
	return &MenuHandler{service: service}
}

// GetActiveMenus is the public API - returns only active menus
func (h *MenuHandler) GetActiveMenus(c *gin.Context) {
	menus, err := h.service.GetActiveMenus()
	if err != nil {
		response.ServerError(c, "获取菜单失败")
		return
	}
	response.Success(c, menus)
}

// GetAllMenus is the admin API - returns all menus
func (h *MenuHandler) GetAllMenus(c *gin.Context) {
	menus, err := h.service.GetAllMenus()
	if err != nil {
		response.ServerError(c, "获取菜单失败")
		return
	}
	response.Success(c, menus)
}

func (h *MenuHandler) Create(c *gin.Context) {
	var item model.MenuItem
	if err := c.ShouldBindJSON(&item); err != nil {
		response.BadRequest(c, "参数错误")
		return
	}
	if err := h.service.Create(&item); err != nil {
		response.ServerError(c, "创建菜单失败")
		return
	}
	response.Success(c, item)
}

func (h *MenuHandler) Update(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效的ID")
		return
	}

	var item model.MenuItem
	if err := c.ShouldBindJSON(&item); err != nil {
		response.BadRequest(c, "参数错误")
		return
	}
	item.ID = uint(id)

	if err := h.service.Update(&item); err != nil {
		response.ServerError(c, "更新菜单失败")
		return
	}
	response.Success(c, item)
}

func (h *MenuHandler) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效的ID")
		return
	}
	if err := h.service.Delete(uint(id)); err != nil {
		response.ServerError(c, "删除菜单失败")
		return
	}
	response.Success(c, nil)
}

type SortRequest struct {
	Items []SortItem `json:"items" binding:"required"`
}

type SortItem struct {
	ID        uint `json:"id"`
	SortOrder int  `json:"sort_order"`
}

func (h *MenuHandler) BatchSort(c *gin.Context) {
	var req SortRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数错误")
		return
	}

	items := make([]model.MenuItem, len(req.Items))
	for i, item := range req.Items {
		items[i] = model.MenuItem{
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
