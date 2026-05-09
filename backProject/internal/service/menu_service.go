package service

import (
	"encoding/json"
	"time"

	"zyqHome/backProject/internal/database"
	"zyqHome/backProject/internal/model"
	"zyqHome/backProject/internal/repository"
)

const menuCacheKey = "zyqhome:menus:active"

type MenuService struct {
	repo *repository.MenuRepo
}

func NewMenuService(repo *repository.MenuRepo) *MenuService {
	return &MenuService{repo: repo}
}

func (s *MenuService) GetActiveMenus() ([]model.MenuItem, error) {
	// Try Redis cache first
	cached, err := database.RDB.Get(database.Ctx, menuCacheKey).Result()
	if err == nil && cached != "" {
		var menus []model.MenuItem
		if json.Unmarshal([]byte(cached), &menus) == nil {
			return menus, nil
		}
	}

	// Query from DB
	items, err := s.repo.FindActive()
	if err != nil {
		return nil, err
	}

	tree := buildMenuTree(items)

	// Write to cache
	if data, err := json.Marshal(tree); err == nil {
		database.RDB.Set(database.Ctx, menuCacheKey, string(data), time.Hour)
	}

	return tree, nil
}

func (s *MenuService) GetAllMenus() ([]model.MenuItem, error) {
	items, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}
	return buildMenuTree(items), nil
}

func (s *MenuService) Create(item *model.MenuItem) error {
	err := s.repo.Create(item)
	if err == nil {
		s.clearCache()
	}
	return err
}

func (s *MenuService) Update(item *model.MenuItem) error {
	err := s.repo.Update(item)
	if err == nil {
		s.clearCache()
	}
	return err
}

func (s *MenuService) Delete(id uint) error {
	err := s.repo.Delete(id)
	if err == nil {
		s.clearCache()
	}
	return err
}

func (s *MenuService) BatchUpdateSort(items []model.MenuItem) error {
	// Convert []model.MenuItem to []model.MenuItem for BatchUpdateSort
	err := s.repo.BatchUpdateSort(items)
	if err == nil {
		s.clearCache()
	}
	return err
}

func (s *MenuService) Count() (int64, error) {
	return s.repo.Count()
}

func (s *MenuService) clearCache() {
	database.RDB.Del(database.Ctx, menuCacheKey)
}

// buildMenuTree assembles flat menu items into a tree structure based on parent_id
func buildMenuTree(items []model.MenuItem) []model.MenuItem {
	menuMap := make(map[uint]*model.MenuItem)
	var roots []model.MenuItem

	for i := range items {
		items[i].Children = []model.MenuItem{}
		menuMap[items[i].ID] = &items[i]
	}

	for i := range items {
		if items[i].ParentID == 0 {
			roots = append(roots, items[i])
		} else if parent, ok := menuMap[items[i].ParentID]; ok {
			parent.Children = append(parent.Children, items[i])
		} else {
			roots = append(roots, items[i])
		}
	}

	return roots
}
