package handlers

import (
	"backend/entities"
	"backend/services"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CreateGameRequest struct {
	PlayersTeam1Ids []uint `json:"playersTeam1Ids" binding:"required,min=1,dive,min=1"`
	PlayersTeam2Ids []uint `json:"playersTeam2Ids" binding:"required,min=1,dive,min=1"`
	Team1Score      uint   `json:"team1Score" binding:"required,min=0"`
	Team2Score      uint   `json:"team2Score" binding:"required,min=0"`
}

type GameHandler struct {
	db            *gorm.DB
	gameService   *services.GameService
	leagueService *services.LeagueService
}

func NewGameHandler(db *gorm.DB, gameService *services.GameService, leagueService *services.LeagueService) *GameHandler {
	return &GameHandler{
		db:            db,
		gameService:   gameService,
		leagueService: leagueService,
	}
}

func (h *GameHandler) CreateGame(c *gin.Context) {
	var req CreateGameRequest

	leagueId, err := strconv.ParseUint(c.Param("leagueId"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid league ID"})
		return
	}

	userID := c.MustGet("userID").(uint)

	if err := h.validateGame(c, &req, uint(leagueId), userID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	game, err := h.gameService.CreateGame(
		uint(leagueId),
		req.Team1Score,
		req.Team2Score,
		req.PlayersTeam1Ids,
		req.PlayersTeam2Ids,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create game"})
		return
	}

	c.JSON(http.StatusOK, game)
}

func (h *GameHandler) GetGameById(c *gin.Context) {
	leagueId, err := strconv.ParseUint(c.Param("leagueId"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid league ID"})
		return
	}

	gameId, err := strconv.ParseUint(c.Param("gameId"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid game ID"})
		return
	}

	userID := c.MustGet("userID").(uint)

	game, err := h.gameService.GetGameByID(uint(gameId), uint(leagueId), userID)
	if err != nil {
		handleServiceError(c, err)
		return
	}

	c.JSON(http.StatusOK, game)
}

// TODO: When player_routes
func (h *GameHandler) GetGamesByPlayerId(c *gin.Context) {
}

func (h *GameHandler) GetGamesByLeagueId(c *gin.Context) {
	leagueId, err := strconv.ParseUint(c.Param("leagueId"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid league ID"})
		return
	}

	userID := c.MustGet("userID").(uint)

	games, err := h.gameService.GetGamesByLeagueId(uint(leagueId), userID)
	if err != nil {
		handleServiceError(c, err)
		return
	}

	c.JSON(http.StatusOK, games)
}

// UpdateGame TODO: Update every other game in the league
func (h *GameHandler) UpdateGame(c *gin.Context) {}

// DeleteGame TODO: Update every other game in the league
func (h *GameHandler) DeleteGame(c *gin.Context) {}

func hasDuplicates(ids []uint) bool {
	seen := make(map[uint]struct{})
	for _, id := range ids {
		if _, exists := seen[id]; exists {
			return true
		}
		seen[id] = struct{}{}
	}
	return false
}

func hasOverlap(team1, team2 []uint) bool {
	team1Set := make(map[uint]struct{})
	for _, id := range team1 {
		team1Set[id] = struct{}{}
	}

	for _, id := range team2 {
		if _, exists := team1Set[id]; exists {
			return true
		}
	}
	return false
}

func handleServiceError(c *gin.Context, err error) {
	switch {
	case errors.Is(err, services.ErrGameNotFound):
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	case errors.Is(err, services.ErrNotLeagueMember):
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
	default:
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
	}
}

func (h *GameHandler) validateGame(c *gin.Context, req *CreateGameRequest, leagueID, userID uint) error {
	if err := c.ShouldBindJSON(&req); err != nil {
		return err
	}

	// Check if league exists and user has access
	league, err := h.leagueService.GetLeagueByID(leagueID, userID)
	if err != nil {
		if errors.Is(err, services.ErrLeagueNotFound) {
			return fmt.Errorf("league not found")
		}
		if errors.Is(err, services.ErrUnauthorized) {
			return fmt.Errorf("you don't have access to this league")
		}
		return fmt.Errorf("failed to fetch league")
	}

	if hasDuplicates(req.PlayersTeam1Ids) {
		return fmt.Errorf("team 1 has duplicate players")
	}
	if hasDuplicates(req.PlayersTeam2Ids) {
		return fmt.Errorf("team 2 has duplicate players")
	}

	if hasOverlap(req.PlayersTeam1Ids, req.PlayersTeam2Ids) {
		return fmt.Errorf("a player cannot be on both teams")
	}

	if len(req.PlayersTeam1Ids) != len(req.PlayersTeam2Ids) {
		return fmt.Errorf("team 1 and team 2 don't have the same number of players")
	}

	// Check if number of players matches league settings
	if len(req.PlayersTeam1Ids) != league.PlayersPerTeam {
		return fmt.Errorf("each team must have %d players as per league settings", league.PlayersPerTeam)
	}

	// Check if draws are allowed in this league
	if !league.HasDraws && req.Team1Score == req.Team2Score {
		return fmt.Errorf("draws are not allowed in this league")
	}

	// Check if all players are in the league
	allPlayerIDs := append(req.PlayersTeam1Ids, req.PlayersTeam2Ids...)
	for _, playerID := range allPlayerIDs {
		var playerLeague entities.PlayerLeague
		err := h.db.Where("user_id = ? AND league_id = ?", playerID, leagueID).First(&playerLeague).Error
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return fmt.Errorf("player with ID %d is not in this league", playerID)
			}
			return fmt.Errorf("failed to verify player membership")
		}
	}

	return nil
}
