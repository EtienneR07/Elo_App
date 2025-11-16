package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID            uint           `gorm:"primaryKey" json:"id"`
	Email         string         `gorm:"uniqueIndex;not null" json:"email"`
	Password      *string        `gorm:"type:text" json:"-"`
	Name          string         `gorm:"not null" json:"name"`
	OAuthProvider string         `gorm:"type:varchar(50)" json:"oauth_provider,omitempty"`
	OAuthID       string         `gorm:"type:varchar(255)" json:"oauth_id,omitempty"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"-"`
}
