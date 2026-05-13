package service

import (
	"errors"
	"zyqHome/backProject/internal/model"
	"zyqHome/backProject/internal/repository"
)

type UserIconService struct {
	repo *repository.UserIconRepo
}

func NewUserIconService(repo *repository.UserIconRepo) *UserIconService {
	return &UserIconService{repo: repo}
}

func (s *UserIconService) GetUserIcons(userID uint) ([]model.UserIcon, error) {
	return s.repo.FindByUserID(userID)
}

func (s *UserIconService) Create(userID uint, imageURL, linkURL string) (*model.UserIcon, error) {
	if imageURL == "" {
		return nil, errors.New("图片URL不能为空")
	}
	if linkURL == "" {
		return nil, errors.New("跳转URL不能为空")
	}
	icon := &model.UserIcon{
		UserID:   userID,
		ImageURL: imageURL,
		LinkURL:  linkURL,
	}
	if err := s.repo.Create(icon); err != nil {
		return nil, errors.New("创建图标失败")
	}
	return icon, nil
}

func (s *UserIconService) Delete(userID, iconID uint) error {
	icon, err := s.repo.FindByID(iconID)
	if err != nil {
		return errors.New("图标不存在")
	}
	if icon.UserID != userID {
		return errors.New("无权限删除此图标")
	}
	return s.repo.Delete(iconID)
}
