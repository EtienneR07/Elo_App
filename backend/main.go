package main

import (
	"backend/config"
	"backend/entities"
	"backend/handlers"
	"backend/routes"
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

	if err := config.AutoMigrate(&entities.User{}, &entities.League{}, &entities.PlayerLeague{}, &entities.Game{}, &entities.GameParticipant{}); err != nil {
		log.Fatal("Failed to run migrations:", err)
	}

	authHandler := handlers.NewAuthHandler(db)
	leagueHandler := handlers.NewLeagueHandler(db)
	gameHandler := handlers.NewGameHandler(db)

	router := gin.Default()

	routes.SetupRoutes(router, authHandler, leagueHandler, gameHandler)

	// Start server
	log.Println("Server starting on :8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
