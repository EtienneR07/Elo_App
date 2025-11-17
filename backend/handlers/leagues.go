package handlers

import (
	"backend/entities"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	ErrLeagueNotFound = errors.New("league not found")
	ErrUnauthorized   = errors.New("unauthorized")
)

type CreateLeagueRequest struct {
	Name           string `json:"name" binding:"required"`
	Description    string `json:"description"`
	PlayersPerTeam int    `json:"playersPerTeam" binding:"required,min=1"`
	Discipline     string `json:"discipline"`
}

type UpdateLeagueRequest struct {
	Name           string `json:"name"`
	Description    string `json:"description"`
	PlayersPerTeam int    `json:"playersPerTeam" binding:"min=1"`
	Discipline     string `json:"discipline"`
}

type LeagueHandler struct {
	db *gorm.DB
}

func NewLeagueHandler(db *gorm.DB) *LeagueHandler {
	return &LeagueHandler{db: db}
}

func (h *LeagueHandler) GetLeagues(c *gin.Context) {
	userID := c.MustGet("userID").(uint)

	var leagues []entities.League
	err := h.db.Where("owner_id = ?", userID).Find(&leagues).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch leagues"})
		return
	}
	c.JSON(http.StatusOK, leagues)
}

func (h *LeagueHandler) GetLeague(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid league ID"})
		return
	}

	userID := c.MustGet("userID").(uint)

	var league entities.League
	err = h.db.First(&league, uint(id)).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "League not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch league"})
		return
	}

	if league.OwnerID != userID {
		c.JSON(http.StatusForbidden, gin.H{"error": "You don't have permission to view this league"})
		return
	}

	c.JSON(http.StatusOK, league)
}

func (h *LeagueHandler) CreateLeague(c *gin.Context) {
	var req CreateLeagueRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID := c.MustGet("userID").(uint)

	league := &entities.League{
		Name:           req.Name,
		Description:    req.Description,
		PlayersPerTeam: req.PlayersPerTeam,
		Discipline:     req.Discipline,
		OwnerID:        userID,
	}

	if err := h.db.Create(league).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create league"})
		return
	}

	c.JSON(http.StatusCreated, league)
}

func (h *LeagueHandler) UpdateLeague(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid league ID"})
		return
	}

	userID := c.MustGet("userID").(uint)

	var req UpdateLeagueRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var league entities.League
	err = h.db.First(&league, uint(id)).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "League not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch league"})
		return
	}

	if league.OwnerID != userID {
		c.JSON(http.StatusForbidden, gin.H{"error": "You don't have permission to update this league"})
		return
	}

	league.Name = req.Name
	league.Description = req.Description
	league.PlayersPerTeam = req.PlayersPerTeam
	league.Discipline = req.Discipline

	if err := h.db.Save(&league).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update league"})
		return
	}

	c.JSON(http.StatusOK, league)
}

func (h *LeagueHandler) DeleteLeague(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid league ID"})
		return
	}

	userID := c.MustGet("userID").(uint)

	var league entities.League
	err = h.db.First(&league, uint(id)).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "League not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch league"})
		return
	}

	if league.OwnerID != userID {
		c.JSON(http.StatusForbidden, gin.H{"error": "You don't have permission to delete this league"})
		return
	}

	if err := h.db.Delete(&entities.League{}, uint(id)).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete league"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "League deleted successfully"})
}
