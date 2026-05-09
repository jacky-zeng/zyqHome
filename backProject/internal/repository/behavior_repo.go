package repository

import (
	"zyqHome/backProject/internal/database"
	"zyqHome/backProject/internal/model"
)

type BehaviorRepo struct{}

func NewBehaviorRepo() *BehaviorRepo {
	return &BehaviorRepo{}
}

func (r *BehaviorRepo) Create(behavior *model.UserBehavior) error {
	return database.DB.Create(behavior).Error
}

func (r *BehaviorRepo) GetTotalViews() (int64, error) {
	var count int64
	err := database.DB.Model(&model.UserBehavior{}).Where("action = ?", "page_view").Count(&count).Error
	return count, err
}

func (r *BehaviorRepo) GetTodayViews() (int64, error) {
	var count int64
	err := database.DB.Model(&model.UserBehavior{}).
		Where("action = ? AND DATE(created_at) = CURDATE()", "page_view").
		Count(&count).Error
	return count, err
}

func (r *BehaviorRepo) GetRecent(limit int) ([]model.UserBehavior, error) {
	var behaviors []model.UserBehavior
	err := database.DB.Order("created_at DESC").Limit(limit).Find(&behaviors).Error
	return behaviors, err
}
