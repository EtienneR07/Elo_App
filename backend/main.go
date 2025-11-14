package main

import (
	"elo-app-backend/handlers"
	"elo-app-backend/middleware"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// Create Gin router
	router := gin.Default()

	// Set trusted proxies (for local development, trust none)
	router.SetTrustedProxies(nil)

	// CORS configuration
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:5173", "http://localhost:5174"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Authorization"}
	config.AllowCredentials = true
	router.Use(cors.New(config))

	// Public routes (no auth required)
	auth := router.Group("/api/auth")
	{
		auth.POST("/login", handlers.Login)
		auth.POST("/register", handlers.Register)
		auth.POST("/logout", handlers.Logout)
	}

	// Protected routes (auth required)
	api := router.Group("/api")
	api.Use(middleware.AuthMiddleware())
	{
		// Auth session check
		api.GET("/auth/session", handlers.GetSession)

		// Leagues
		api.GET("/leagues", handlers.GetLeagues)
		api.POST("/leagues", handlers.CreateLeague)
		api.GET("/leagues/:id", handlers.GetLeague)
		api.PUT("/leagues/:id", handlers.UpdateLeague)
		api.DELETE("/leagues/:id", handlers.DeleteLeague)
	}

	// Start server
	log.Println("Server starting on :8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
