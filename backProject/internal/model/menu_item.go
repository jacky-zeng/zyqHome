package model

import "time"

type MenuItem struct {
	ID        uint       `gorm:"primaryKey;autoIncrement" json:"id"`
	Title     string     `gorm:"size:100;not null" json:"title"`
	URL       string     `gorm:"size:500;not null" json:"url"`
	Icon      string     `gorm:"size:100;default:''" json:"icon"`
	ParentID  uint       `gorm:"default:0;index" json:"parent_id"`
	SortOrder int        `gorm:"default:0;index" json:"sort_order"`
	IsActive  bool       `gorm:"default:true" json:"is_active"`
	Target    string     `gorm:"size:10;default:'_blank'" json:"target"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	Children  []MenuItem `gorm:"-" json:"children,omitempty"`
}

func (MenuItem) TableName() string {
	return "menu_items"
}
