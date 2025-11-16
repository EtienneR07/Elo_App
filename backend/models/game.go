package models

import (
	"time"

	"gorm.io/gorm"
)

type Game struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	LeagueID  uint           `gorm:"not null;index" json:"league_id"`
	League    League         `gorm:"foreignKey:LeagueID" json:"league,omitempty"`
	PlayedAt  time.Time      `gorm:"index" json:"played_at"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
