package post

import (
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Post struct {
	ID        int            `json:"id" gorm:"primaryKey"`
	Title     string         `json:"title" gorm:"column:title"`
	Slug      string         `json:"slug" gorm:"column:slug"`
	Body      string         `json:"body" gorm:"column:body"`
	Image     string         `json:"image" gorm:"column:image"`
	ReadTime  string         `json:"readTime" gorm:"column:read_time"`
	Views     int            `json:"views" gorm:"column:views;default:0"`
	Tags      datatypes.JSON `json:"tags" gorm:"column:tags"`
	Category  string         `json:"category" gorm:"column:category"`
	LikeCount int            `json:"likeCount" gorm:"column:like_count;default:0"`
	CreatedAt time.Time      `json:"createdAt" gorm:"column:created_at"`
	UpdatedAt time.Time      `json:"updatedAt" gorm:"column:updated_at"`
	DeletedAt gorm.DeletedAt `json:"deletedAt" gorm:"column:deleted_at"`
}
