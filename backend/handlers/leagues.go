package handlers

import (
	"backend/models"
	"backend/services"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
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
	leagueService services.LeagueService
}

func NewLeagueHandler(leagueService services.LeagueService) *LeagueHandler {
	return &LeagueHandler{leagueService: leagueService}
}

func (h *LeagueHandler) GetLeagues(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Not authenticated"})
		return
	}

	leagues, err := h.leagueService.GetUserLeagues(userID.(uint))
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

	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Not authenticated"})
		return
	}

	league, err := h.leagueService.GetLeagueByID(uint(id), userID.(uint))
	if err != nil {
		if errors.Is(err, services.ErrLeagueNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "League not found"})
			return
		}
		if errors.Is(err, services.ErrUnauthorized) {
			c.JSON(http.StatusForbidden, gin.H{"error": "You don't have permission to view this league"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch league"})
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

	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Not authenticated"})
		return
	}

	league := &models.League{
		Name:           req.Name,
		Description:    req.Description,
		PlayersPerTeam: req.PlayersPerTeam,
		Discipline:     req.Discipline,
		OwnerID:        userID.(uint),
	}

	if err := h.leagueService.CreateLeague(league); err != nil {
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

	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Not authenticated"})
		return
	}

	var req UpdateLeagueRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	league := &models.League{
		Name:           req.Name,
		Description:    req.Description,
		PlayersPerTeam: req.PlayersPerTeam,
		Discipline:     req.Discipline,
	}

	if err := h.leagueService.UpdateLeague(uint(id), userID.(uint), league); err != nil {
		if errors.Is(err, services.ErrLeagueNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "League not found"})
			return
		}
		if errors.Is(err, services.ErrUnauthorized) {
			c.JSON(http.StatusForbidden, gin.H{"error": "You don't have permission to update this league"})
			return
		}
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

	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Not authenticated"})
		return
	}

	if err := h.leagueService.DeleteLeague(uint(id), userID.(uint)); err != nil {
		if errors.Is(err, services.ErrLeagueNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "League not found"})
			return
		}
		if errors.Is(err, services.ErrUnauthorized) {
			c.JSON(http.StatusForbidden, gin.H{"error": "You don't have permission to delete this league"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete league"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "League deleted successfully"})
}
