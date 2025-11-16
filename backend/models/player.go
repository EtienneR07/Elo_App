package models

import (
	"time"

	"gorm.io/gorm"
)

type PlayerLeague struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	UserID    uint           `gorm:"not null;uniqueIndex:idx_user_league" json:"user_id"`
	User      User           `gorm:"foreignKey:UserID" json:"user,omitempty"`
	LeagueID  uint           `gorm:"not null;uniqueIndex:idx_user_league" json:"league_id"`
	League    League         `gorm:"foreignKey:LeagueID" json:"league,omitempty"`
	Wins      uint           `gorm:"not null;default:0" json:"wins"`
	Losses    uint           `gorm:"not null;default:0" json:"losses"`
	Draws     uint           `gorm:"not null;default:0" json:"draws"`
	Elo       uint           `gorm:"not null;default:1200" json:"elo"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
