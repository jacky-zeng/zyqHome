package repository

import (
	"time"

	"zyqHome/backProject/internal/database"
	"zyqHome/backProject/internal/model"
)

type UserRepo struct{}

func NewUserRepo() *UserRepo {
	return &UserRepo{}
}

func (r *UserRepo) FindByUsername(username string) (*model.User, error) {
	var user model.User
	err := database.DB.Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepo) FindByID(id uint) (*model.User, error) {
	var user model.User
	err := database.DB.First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepo) Create(user *model.User) error {
	return database.DB.Create(user).Error
}

func (r *UserRepo) UpdateLastLogin(id uint, ip string) error {
	return database.DB.Model(&model.User{}).Where("id = ?", id).Updates(map[string]interface{}{
		"last_login_ip": ip,
		"last_login_at": time.Now(),
	}).Error
}
