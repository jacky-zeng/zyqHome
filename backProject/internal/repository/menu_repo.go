package repository

import (
	"zyqHome/backProject/internal/database"
	"zyqHome/backProject/internal/model"

	"gorm.io/gorm"
)

type MenuRepo struct{}

func NewMenuRepo() *MenuRepo {
	return &MenuRepo{}
}

func (r *MenuRepo) FindAll() ([]model.MenuItem, error) {
	var items []model.MenuItem
	err := database.DB.Order("sort_order ASC, id ASC").Find(&items).Error
	return items, err
}

func (r *MenuRepo) FindActive() ([]model.MenuItem, error) {
	var items []model.MenuItem
	err := database.DB.Where("is_active = ?", true).Order("sort_order ASC, id ASC").Find(&items).Error
	return items, err
}

func (r *MenuRepo) FindByID(id uint) (*model.MenuItem, error) {
	var item model.MenuItem
	err := database.DB.First(&item, id).Error
	if err != nil {
		return nil, err
	}
	return &item, nil
}

func (r *MenuRepo) Create(item *model.MenuItem) error {
	return database.DB.Create(item).Error
}

func (r *MenuRepo) Update(item *model.MenuItem) error {
	return database.DB.Save(item).Error
}

func (r *MenuRepo) Delete(id uint) error {
	return database.DB.Delete(&model.MenuItem{}, id).Error
}

func (r *MenuRepo) BatchUpdateSort(items []model.MenuItem) error {
	return database.DB.Transaction(func(tx *gorm.DB) error {
		for _, item := range items {
			if err := tx.Model(&model.MenuItem{}).Where("id = ?", item.ID).Update("sort_order", item.SortOrder).Error; err != nil {
				return err
			}
		}
		return nil
	})
}

func (r *MenuRepo) Count() (int64, error) {
	var count int64
	err := database.DB.Model(&model.MenuItem{}).Count(&count).Error
	return count, err
}
