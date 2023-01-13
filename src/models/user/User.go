package user

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	// gorm.Model
	Id        int            `json:"id" gorm:"column:id;primary_key"`
	Username  string         `json:"username" gorm:"column:username;unique;not null"`
	Email     string         `json:"email" gorm:"column:email;unique;not null"`
	Password  string         `json:"password" gorm:"column:password;not null"`
	CreatedAt time.Time      `json:"createdAt" gorm:"column:created_at"`
	UpdatedAt time.Time      `json:"updatedAt" gorm:"column:updated_at"`
	DeletedAt gorm.DeletedAt `json:"deletedAt" gorm:"column:deleted_at"`
}
