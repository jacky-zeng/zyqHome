package repository

import (
	"zyqHome/backProject/internal/database"
	"zyqHome/backProject/internal/model"

	"gorm.io/gorm"
)

type ImageRepo struct{}

func NewImageRepo() *ImageRepo {
	return &ImageRepo{}
}

func (r *ImageRepo) FindAll(category string, page, pageSize int) ([]model.Image, int64, error) {
	var images []model.Image
	var total int64

	query := database.DB.Model(&model.Image{}).Where("user_id = 0")
	if category != "" {
		query = query.Where("category = ?", category)
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * pageSize
	err := query.Order("created_at DESC").Offset(offset).Limit(pageSize).Find(&images).Error
	return images, total, err
}

func (r *ImageRepo) FindAllByUser(userID uint, category string, page, pageSize int) ([]model.Image, int64, error) {
	var images []model.Image
	var total int64

	query := database.DB.Model(&model.Image{}).Where("user_id = ?", userID)
	if category != "" {
		query = query.Where("category = ?", category)
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * pageSize
	err := query.Order("created_at DESC").Offset(offset).Limit(pageSize).Find(&images).Error
	return images, total, err
}

func (r *ImageRepo) CountByUser(userID uint) (int64, error) {
	var total int64
	err := database.DB.Model(&model.Image{}).Where("user_id = ?", userID).Count(&total).Error
	return total, err
}

func (r *ImageRepo) FindCategories() ([]string, error) {
	var categories []string
	err := database.DB.Model(&model.Image{}).Where("user_id = 0").Select("DISTINCT category").Order("category ASC").Pluck("category", &categories).Error
	return categories, err
}

func (r *ImageRepo) FindCategoriesByUser(userID uint) ([]string, error) {
	var categories []string
	err := database.DB.Model(&model.Image{}).Where("user_id = ?", userID).Select("DISTINCT category").Order("category ASC").Pluck("category", &categories).Error
	return categories, err
}

func (r *ImageRepo) FindByID(id uint) (*model.Image, error) {
	var image model.Image
	err := database.DB.First(&image, id).Error
	if err != nil {
		return nil, err
	}
	return &image, nil
}

func (r *ImageRepo) Create(image *model.Image) error {
	return database.DB.Create(image).Error
}

func (r *ImageRepo) Update(image *model.Image) error {
	return database.DB.Save(image).Error
}

func (r *ImageRepo) Delete(id uint) error {
	return database.DB.Delete(&model.Image{}, id).Error
}

func (r *ImageRepo) DeleteTx(tx *gorm.DB, id uint) error {
	return tx.Delete(&model.Image{}, id).Error
}
