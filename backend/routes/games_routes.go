package routes

import (
	"backend/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterGameRoutes(rg *gin.RouterGroup, gameHandler *handlers.GameHandler) {
	games := rg.Group("/games")
	{
		games.GET("", gameHandler.GetGamesByLeagueId)
		games.POST("", gameHandler.CreateGame)

		games.GET("/:gameId", gameHandler.GetGameById)
		games.PUT("/:gameId", gameHandler.UpdateGame)
		games.DELETE("/:gameId", gameHandler.DeleteGame)
	}
}
