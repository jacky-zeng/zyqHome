package service

import (
	"zyqHome/backProject/internal/model"
	"zyqHome/backProject/internal/repository"
)

type BehaviorService struct {
	repo *repository.BehaviorRepo
}

func NewBehaviorService(repo *repository.BehaviorRepo) *BehaviorService {
	return &BehaviorService{repo: repo}
}

func (s *BehaviorService) Track(ip, action, target, userAgent, referer string) error {
	behavior := &model.UserBehavior{
		IPAddress: ip,
		Action:    action,
		Target:    target,
		UserAgent: userAgent,
		Referer:   referer,
	}
	return s.repo.Create(behavior)
}

type DashboardStats struct {
	TotalViews     int64                `json:"total_views"`
	TodayViews     int64                `json:"today_views"`
	TotalIcons     int64                `json:"total_icons"`
	TotalMenus     int64                `json:"total_menus"`
	RecentBehaviors []model.UserBehavior `json:"recent_behaviors"`
}
