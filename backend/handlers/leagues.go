package handlers

import (
	"backend/services"
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type CreateLeagueRequest struct {
	Name           string    `json:"name" binding:"required"`
	Description    string    `json:"description"`
	PlayersPerTeam int       `json:"playersPerTeam" binding:"required,min=1"`
	Discipline     string    `json:"discipline"`
	HasDraws       bool      `json:"hasDraws"`
	Start          time.Time `json:"start" binding:"required"`
	End            time.Time `json:"end" binding:"required"`
}

type UpdateLeagueRequest struct {
	Name           string `json:"name"`
	Description    string `json:"description"`
	PlayersPerTeam int    `json:"playersPerTeam" binding:"min=1"`
	Discipline     string `json:"discipline"`
}

type LeagueHandler struct {
	leagueService *services.LeagueService
}

func NewLeagueHandler(leagueService *services.LeagueService) *LeagueHandler {
	return &LeagueHandler{leagueService: leagueService}
}

func (h *LeagueHandler) GetLeagues(c *gin.Context) {
	userID := c.MustGet("userID").(uint)

	leagues, err := h.leagueService.GetLeaguesByOwner(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch leagues"})
		return
	}
	c.JSON(http.StatusOK, leagues)
}

func (h *LeagueHandler) GetLeague(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("leagueId"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid league ID"})
		return
	}

	userID := c.MustGet("userID").(uint)

	league, err := h.leagueService.GetLeagueByID(uint(id), userID)
	if err != nil {
		handleLeagueServiceError(c, err)
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

	league, err := h.leagueService.CreateLeague(
		userID,
		req.Name,
		req.Description,
		req.Discipline,
		req.PlayersPerTeam,
		req.HasDraws,
		req.Start,
		req.End,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create league"})
		return
	}

	c.JSON(http.StatusCreated, league)
}

func (h *LeagueHandler) UpdateLeague(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("leagueId"), 10, 32)
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

	league, err := h.leagueService.UpdateLeague(
		uint(id),
		userID,
		req.Name,
		req.Description,
	)
	if err != nil {
		handleLeagueServiceError(c, err)
		return
	}

	c.JSON(http.StatusOK, league)
}

func (h *LeagueHandler) DeleteLeague(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("leagueId"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid league ID"})
		return
	}

	userID := c.MustGet("userID").(uint)

	err = h.leagueService.DeleteLeague(uint(id), userID)
	if err != nil {
		handleLeagueServiceError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "League deleted successfully"})
}

func handleLeagueServiceError(c *gin.Context, err error) {
	switch {
	case errors.Is(err, services.ErrLeagueNotFound):
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	case errors.Is(err, services.ErrUnauthorized):
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
	default:
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
	}
}
