package main

import (
	"github.com/VictorzllDev/ZorpPay/backend/internal/api/handler"
	"github.com/VictorzllDev/ZorpPay/backend/internal/api/repository"
	"github.com/VictorzllDev/ZorpPay/backend/internal/api/routes"
	"github.com/VictorzllDev/ZorpPay/backend/internal/api/service"
	"github.com/VictorzllDev/ZorpPay/backend/internal/config"
	"github.com/VictorzllDev/ZorpPay/backend/internal/database"
	"github.com/gin-gonic/gin"
	"log"
)

// @title           ZorpPay
// @version         1.0.0
// @description     ZorpPay API Documentation

// @contact.name   VictorzllDev
// contact.url    https://github.com/VictorzllDev

// @BasePath  /api/v1/
func main() {
	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatal("Failed to load config:", err)
	}

	// Connect to database
	db, err := database.Connect(cfg.DatabaseURL)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Dependence injection
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	// Routes
	r := gin.Default()
	routes.UserRoutes(r, userHandler)

	// Specific Routes for Development
	if cfg.Env == "development" {
		routes.DocsRoutes(r)
	}

	// Run server
	log.Println("Server running on port", cfg.Port)
	if err := r.Run(":" + cfg.Port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
