package models

import (
	"time"

	"gorm.io/gorm"
)

type GameParticipant struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	GameID    uint           `gorm:"not null;index" json:"game_id"`
	Game      Game           `gorm:"foreignKey:GameID" json:"game,omitempty"`
	PlayerID  uint           `gorm:"not null;index" json:"player_id"`
	Player    PlayerLeague   `gorm:"foreignKey:PlayerID" json:"player,omitempty"`
	Team      int            `gorm:"not null;check:team IN (1, 2)" json:"team"`
	EloBefore uint           `gorm:"not null" json:"elo_before"`
	EloAfter  uint           `gorm:"not null" json:"elo_after"`
	EloChange int            `gorm:"not null" json:"elo_change"`
	Result    string         `gorm:"type:varchar(20);not null;check:result IN ('win', 'loss', 'draw')" json:"result"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
