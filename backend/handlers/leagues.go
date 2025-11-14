package handlers

import (
	"elo-app-backend/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// In-memory league storage (replace with database in production)
var leagues = make(map[string]*models.League)

type CreateLeagueRequest struct {
	Name           string `json:"name" binding:"required"`
	Description    string `json:"description"`
	PlayersPerTeam int    `json:"playersPerTeam" binding:"required,min=1"`
	Discipline     string `json:"discipline"`
	OwnerID        string `json:"ownerId" binding:"required"`
}

type UpdateLeagueRequest struct {
	Name           string `json:"name"`
	Description    string `json:"description"`
	PlayersPerTeam int    `json:"playersPerTeam" binding:"min=1"`
	Discipline     string `json:"discipline"`
}

func GetLeagues(c *gin.Context) {
	result := make([]*models.League, 0, len(leagues))
	for _, league := range leagues {
		result = append(result, league)
	}
	c.JSON(http.StatusOK, result)
}

func GetLeague(c *gin.Context) {
	id := c.Param("id")
	league, exists := leagues[id]
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "League not found"})
		return
	}
	c.JSON(http.StatusOK, league)
}

func CreateLeague(c *gin.Context) {
	var req CreateLeagueRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Verify user is authenticated
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Not authenticated"})
		return
	}

	league := &models.League{
		ID:             uuid.New().String(),
		Name:           req.Name,
		Description:    req.Description,
		PlayersPerTeam: req.PlayersPerTeam,
		Discipline:     req.Discipline,
		OwnerID:        userID.(string),
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}

	leagues[league.ID] = league
	c.JSON(http.StatusCreated, league)
}

func UpdateLeague(c *gin.Context) {
	id := c.Param("id")
	league, exists := leagues[id]
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "League not found"})
		return
	}

	// Verify user owns the league
	userID, _ := c.Get("userID")
	if league.OwnerID != userID.(string) {
		c.JSON(http.StatusForbidden, gin.H{"error": "You don't have permission to update this league"})
		return
	}

	var req UpdateLeagueRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update fields
	if req.Name != "" {
		league.Name = req.Name
	}
	if req.Description != "" {
		league.Description = req.Description
	}
	if req.PlayersPerTeam > 0 {
		league.PlayersPerTeam = req.PlayersPerTeam
	}
	if req.Discipline != "" {
		league.Discipline = req.Discipline
	}
	league.UpdatedAt = time.Now()

	c.JSON(http.StatusOK, league)
}

func DeleteLeague(c *gin.Context) {
	id := c.Param("id")
	league, exists := leagues[id]
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "League not found"})
		return
	}

	// Verify user owns the league
	userID, _ := c.Get("userID")
	if league.OwnerID != userID.(string) {
		c.JSON(http.StatusForbidden, gin.H{"error": "You don't have permission to delete this league"})
		return
	}

	delete(leagues, id)
	c.JSON(http.StatusOK, gin.H{"message": "League deleted successfully"})
}
