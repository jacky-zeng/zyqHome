package model

import "time"

type UserBehavior struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	IPAddress string    `gorm:"size:45;not null;index" json:"ip_address"`
	Action    string    `gorm:"size:50;not null;index" json:"action"`
	Target    string    `gorm:"size:500;default:''" json:"target"`
	UserAgent string    `gorm:"size:500;default:''" json:"user_agent"`
	Referer   string    `gorm:"size:500;default:''" json:"referer"`
	CreatedAt time.Time `gorm:"index" json:"created_at"`
}

func (UserBehavior) TableName() string {
	return "user_behaviors"
}
