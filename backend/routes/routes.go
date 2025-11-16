package routes

import (
	"backend/handlers"
	"backend/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, authHandler *handlers.AuthHandler, leagueHandler *handlers.LeagueHandler) {
	// Set trusted proxies (for local development, trust none)
	err := router.SetTrustedProxies(nil)

	if err != nil {
		return
	}

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://localhost:5173", "http://localhost:5174"}
	corsConfig.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	corsConfig.AllowHeaders = []string{"Origin", "Content-Type", "Authorization"}
	corsConfig.AllowCredentials = true

	router.Use(cors.New(corsConfig))

	api := router.Group("/api")
	{
		// Public routes
		RegisterAuthRoutes(api, authHandler)

		// Protected routes
		protected := api.Group("")
		protected.Use(middleware.AuthMiddleware())
		{
			protected.GET("/auth/session", authHandler.GetSession)

			RegisterLeagueRoutes(protected, leagueHandler)
		}
	}
}
