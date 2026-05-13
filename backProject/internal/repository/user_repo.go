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

func (r *UserRepo) FindByEmail(email string) (*model.User, error) {
	var user model.User
	err := database.DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepo) ExistsByUsername(username string) (bool, error) {
	var count int64
	err := database.DB.Model(&model.User{}).Where("username = ?", username).Count(&count).Error
	return count > 0, err
}

func (r *UserRepo) ExistsByEmail(email string) (bool, error) {
	var count int64
	err := database.DB.Model(&model.User{}).Where("email = ?", email).Count(&count).Error
	return count > 0, err
}

func (r *UserRepo) FindMembers(page, pageSize int) ([]model.User, int64, error) {
	var users []model.User
	var total int64

	query := database.DB.Model(&model.User{}).Where("role = ?", "member")

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * pageSize
	if err := query.Offset(offset).Limit(pageSize).Order("id DESC").Find(&users).Error; err != nil {
		return nil, 0, err
	}

	return users, total, nil
}

func (r *UserRepo) UpdateFields(id uint, fields map[string]interface{}) error {
	return database.DB.Model(&model.User{}).Where("id = ?", id).Updates(fields).Error
}
