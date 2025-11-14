package models

import "time"

type League struct {
	ID             string    `json:"id"`
	Name           string    `json:"name"`
	Description    string    `json:"description,omitempty"`
	PlayersPerTeam int       `json:"playersPerTeam"`
	Discipline     string    `json:"discipline,omitempty"`
	OwnerID        string    `json:"ownerId"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
}
