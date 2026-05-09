package service

import (
	"encoding/json"
	"time"

	"zyqHome/backProject/internal/database"
	"zyqHome/backProject/internal/model"
	"zyqHome/backProject/internal/repository"
)

const iconCacheKey = "zyqhome:icons:active"

type IconService struct {
	repo *repository.IconRepo
}

func NewIconService(repo *repository.IconRepo) *IconService {
	return &IconService{repo: repo}
}

func (s *IconService) GetActiveIcons() ([]model.CenterIcon, error) {
	cached, err := database.RDB.Get(database.Ctx, iconCacheKey).Result()
	if err == nil && cached != "" {
		var icons []model.CenterIcon
		if json.Unmarshal([]byte(cached), &icons) == nil {
			return icons, nil
		}
	}

	icons, err := s.repo.FindActive()
	if err != nil {
		return nil, err
	}

	if data, err := json.Marshal(icons); err == nil {
		database.RDB.Set(database.Ctx, iconCacheKey, string(data), time.Hour)
	}

	return icons, nil
}

func (s *IconService) GetActiveIconsByMenuID(menuID uint) ([]model.CenterIcon, error) {
	return s.repo.FindActiveByMenuID(menuID)
}

func (s *IconService) GetAllIcons() ([]model.CenterIcon, error) {
	return s.repo.FindAll()
}

func (s *IconService) Create(icon *model.CenterIcon) error {
	err := s.repo.Create(icon)
	if err == nil {
		s.clearCache()
	}
	return err
}

func (s *IconService) Update(icon *model.CenterIcon) error {
	err := s.repo.Update(icon)
	if err == nil {
		s.clearCache()
	}
	return err
}

func (s *IconService) Delete(id uint) error {
	err := s.repo.Delete(id)
	if err == nil {
		s.clearCache()
	}
	return err
}

func (s *IconService) BatchUpdateSort(items []model.CenterIcon) error {
	err := s.repo.BatchUpdateSort(items)
	if err == nil {
		s.clearCache()
	}
	return err
}

func (s *IconService) Count() (int64, error) {
	return s.repo.Count()
}

func (s *IconService) clearCache() {
	database.RDB.Del(database.Ctx, iconCacheKey)
}
