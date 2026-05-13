package model

import "time"

type UserIcon struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID    uint      `gorm:"index;not null" json:"user_id"`
	ImageURL  string    `gorm:"size:500;not null" json:"image_url"`
	LinkURL   string    `gorm:"size:500;not null" json:"link_url"`
	SortOrder int       `gorm:"default:0" json:"sort_order"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (UserIcon) TableName() string {
	return "user_icons"
}
