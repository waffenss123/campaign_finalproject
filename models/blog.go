package models

import (
	"time"

	"gorm.io/gorm"
)

type BlogDB struct {
	ID         string         `gorm:"primaryKey" json:"id"`
	UserID     int32          `json:"user_id"`
	CampaignID string         `json:"campaign_id"`
	Content    string         `json:"content"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `json:"-"` // soft-delete kalau mau
}

func (BlogDB) TableName() string {
	return "blogs.blogs"
}
