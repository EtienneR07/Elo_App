package main

import (
	"backend/config"
	"backend/handlers"
	"backend/models"
	"backend/repository"
	"backend/routes"
	"backend/services"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	config.InitGoogleOAuth()

	if err := config.InitDatabase(); err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	db := config.GetDB()
	
	if err := config.AutoMigrate(&models.User{}, &models.League{}, &models.PlayerLeague{}, &models.Game{}, &models.GameParticipant{}); err != nil {
		log.Fatal("Failed to run migrations:", err)
	}

	userRepo := repository.NewUserRepository(db)
	leagueRepo := repository.NewLeagueRepository(db)

	authService := services.NewAuthService(userRepo)
	leagueService := services.NewLeagueService(leagueRepo)

	authHandler := handlers.NewAuthHandler(authService)
	leagueHandler := handlers.NewLeagueHandler(leagueService)

	router := gin.Default()

	routes.SetupRoutes(router, authHandler, leagueHandler)

	// Start server
	log.Println("Server starting on :8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
