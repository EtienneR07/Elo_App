package models

import (
	"time"

	"gorm.io/gorm"
)

type League struct {
	ID             uint           `gorm:"primaryKey" json:"id"`
	Name           string         `gorm:"not null" json:"name"`
	Description    string         `gorm:"type:text" json:"description,omitempty"`
	PlayersPerTeam int            `gorm:"not null" json:"players_per_team"`
	Discipline     string         `gorm:"type:varchar(100)" json:"discipline,omitempty"`
	Start          time.Time      `gorm:"not null" json:"start"`
	End            time.Time      `gorm:"not null" json:"end"`
	OwnerID        uint           `gorm:"not null;index" json:"owner_id"`
	Owner          User           `gorm:"foreignKey:OwnerID" json:"owner,omitempty"` // Relationship to User
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
	DeletedAt      gorm.DeletedAt `gorm:"index" json:"-"` // Soft delete support
}
