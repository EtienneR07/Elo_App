package handlers

import (
	"backend/entities"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type GameHandler struct {
	db *gorm.DB
}

func NewGameHandler(db *gorm.DB) *GameHandler {
	return &GameHandler{db: db}
}

func (h *GameHandler) CreateGame(c *gin.Context) {
	c.JSON(http.StatusOK, entities.Game{})
}

func (h *GameHandler) GetGameById(c *gin.Context) {
	_, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid game ID"})
		return
	}

}

func (h *GameHandler) GetGamesByPlayerId(c *gin.Context) {
}

func (h *GameHandler) GetGamesByLeagueId(c *gin.Context) {
}

// UpdateGame TODO: Update every other game in the league
func (h *GameHandler) UpdateGame(c *gin.Context) {}

// DeleteGame TODO: Update every other game in the league
func (h *GameHandler) DeleteGame(c *gin.Context) {}
