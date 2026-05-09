package model

import "time"

type CenterIcon struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Title     string    `gorm:"size:100;not null" json:"title"`
	URL       string    `gorm:"size:500;not null" json:"url"`
	Icon      string    `gorm:"size:255;default:''" json:"icon"`
	Color     string    `gorm:"size:7;default:'#1890ff'" json:"color"`
	SortOrder int       `gorm:"default:0;index" json:"sort_order"`
	IsActive  bool      `gorm:"default:true" json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (CenterIcon) TableName() string {
	return "center_icons"
}
