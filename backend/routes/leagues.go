package routes

import (
	"backend/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterLeagueRoutes(rg *gin.RouterGroup, leagueHandler *handlers.LeagueHandler) {
	leagues := rg.Group("/leagues")
	{
		leagues.GET("", leagueHandler.GetLeagues)
		leagues.POST("", leagueHandler.CreateLeague)
		leagues.GET("/:id", leagueHandler.GetLeague)
		leagues.PUT("/:id", leagueHandler.UpdateLeague)
		leagues.DELETE("/:id", leagueHandler.DeleteLeague)
	}
}
