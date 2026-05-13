package model

import "time"

type Image struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID    uint      `gorm:"default:0;index" json:"user_id"`
	Filename  string    `gorm:"size:255;not null" json:"filename"`
	URL       string    `gorm:"size:500;not null" json:"url"`
	Category  string    `gorm:"size:100;default:'default';index" json:"category"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (Image) TableName() string {
	return "images"
}
