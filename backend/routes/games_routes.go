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

		games.GET("/:id", gameHandler.GetGameById)
		games.PUT("/:id", gameHandler.UpdateGame)
		games.DELETE("/:id", gameHandler.DeleteGame)
	}
}
