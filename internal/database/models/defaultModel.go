package models

import (
	"time"
)

// gorm.Model definition
type Default struct {
	Id        uint      `json:"id" gorm:"primaryKey;not null" autoIncrement:"true"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime:true"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime:true"`
}
