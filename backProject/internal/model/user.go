package model

import "time"

type User struct {
	ID          uint       `gorm:"primaryKey;autoIncrement" json:"id"`
	Username    string     `gorm:"uniqueIndex;size:50;not null" json:"username"`
	Password    string     `gorm:"size:255;not null" json:"-"`
	Nickname    string     `gorm:"size:100;default:''" json:"nickname"`
	Email       string     `gorm:"size:100;default:'';index" json:"email"`
	Avatar      string     `gorm:"size:255;default:''" json:"avatar"`
	Role        string     `gorm:"size:20;default:'';index" json:"role"` // "admin" / "member"
	Status      int        `gorm:"default:1" json:"status"`           // 1=启用 0=禁用
	LastLoginIP string     `gorm:"size:45;default:''" json:"last_login_ip"`
	LastLoginAt *time.Time `json:"last_login_at"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}

func (User) TableName() string {
	return "users"
}
