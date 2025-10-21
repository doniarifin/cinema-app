package main

import (
	"cinema-app/config"
	"cinema-app/internal/database"
	"cinema-app/internal/routes"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	// Init env
	config.InitEnv()

	// Connect DB
	db, err := database.ConnectDB()
	if err != nil {
		log.Fatal("DB connection error: ", err)
	}

	database.RunMigration(db)

	// Init Gin
	app := gin.Default()

	// trusted proxy
	trusted := []string{
		"127.0.0.1",
	}
	err = app.SetTrustedProxies(trusted)
	if err != nil {
		log.Fatal("Proxy config error:", err)
	}

	// Setup all routes
	routes.SetupRoutes(app, db)

	// Run server
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}
	log.Println("Server running on port " + port)
	// app.Run(":" + port)
	if err := app.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
