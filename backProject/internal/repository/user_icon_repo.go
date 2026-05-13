package repository

import (
	"zyqHome/backProject/internal/database"
	"zyqHome/backProject/internal/model"
)

type UserIconRepo struct{}

func NewUserIconRepo() *UserIconRepo {
	return &UserIconRepo{}
}

func (r *UserIconRepo) FindByUserID(userID uint) ([]model.UserIcon, error) {
	var icons []model.UserIcon
	err := database.DB.Where("user_id = ?", userID).
		Order("sort_order ASC, id ASC").Find(&icons).Error
	return icons, err
}

func (r *UserIconRepo) FindByID(id uint) (*model.UserIcon, error) {
	var icon model.UserIcon
	err := database.DB.First(&icon, id).Error
	if err != nil {
		return nil, err
	}
	return &icon, nil
}

func (r *UserIconRepo) Create(icon *model.UserIcon) error {
	return database.DB.Create(icon).Error
}

func (r *UserIconRepo) Delete(id uint) error {
	return database.DB.Delete(&model.UserIcon{}, id).Error
}
