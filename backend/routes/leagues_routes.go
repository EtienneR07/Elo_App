package routes

import (
	"backend/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterLeagueRoutes(rg *gin.RouterGroup, leagueHandler *handlers.LeagueHandler, gameHandler *handlers.GameHandler) {
	leagues := rg.Group("/leagues")
	{
		leagues.GET("", leagueHandler.GetLeagues)
		leagues.POST("", leagueHandler.CreateLeague)

		leaguesId := leagues.Group("/:id")
		{
			leaguesId.GET("", leagueHandler.GetLeague)
			leaguesId.PUT("", leagueHandler.UpdateLeague)
			leaguesId.DELETE("", leagueHandler.DeleteLeague)

			RegisterGameRoutes(leaguesId, gameHandler)
		}
	}
}
