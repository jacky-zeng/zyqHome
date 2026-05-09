package repository

import (
	"zyqHome/backProject/internal/database"
	"zyqHome/backProject/internal/model"

	"gorm.io/gorm"
)

type IconRepo struct{}

func NewIconRepo() *IconRepo {
	return &IconRepo{}
}

func (r *IconRepo) FindAll() ([]model.CenterIcon, error) {
	var icons []model.CenterIcon
	err := database.DB.Order("sort_order ASC, id ASC").Find(&icons).Error
	return icons, err
}

func (r *IconRepo) FindActive() ([]model.CenterIcon, error) {
	var icons []model.CenterIcon
	err := database.DB.Where("is_active = ?", true).Order("sort_order ASC, id ASC").Find(&icons).Error
	return icons, err
}

func (r *IconRepo) FindByID(id uint) (*model.CenterIcon, error) {
	var icon model.CenterIcon
	err := database.DB.First(&icon, id).Error
	if err != nil {
		return nil, err
	}
	return &icon, nil
}

func (r *IconRepo) Create(icon *model.CenterIcon) error {
	return database.DB.Create(icon).Error
}

func (r *IconRepo) Update(icon *model.CenterIcon) error {
	return database.DB.Save(icon).Error
}

func (r *IconRepo) Delete(id uint) error {
	return database.DB.Delete(&model.CenterIcon{}, id).Error
}

func (r *IconRepo) BatchUpdateSort(items []model.CenterIcon) error {
	return database.DB.Transaction(func(tx *gorm.DB) error {
		for _, item := range items {
			if err := tx.Model(&model.CenterIcon{}).Where("id = ?", item.ID).Update("sort_order", item.SortOrder).Error; err != nil {
				return err
			}
		}
		return nil
	})
}

func (r *IconRepo) Count() (int64, error) {
	var count int64
	err := database.DB.Model(&model.CenterIcon{}).Count(&count).Error
	return count, err
}
