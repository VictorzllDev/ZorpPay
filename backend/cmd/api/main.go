package main

import (
	"log"

	"github.com/VictorzllDev/ZorpPay/backend/internal/api/handler"
	"github.com/VictorzllDev/ZorpPay/backend/internal/api/repository"
	"github.com/VictorzllDev/ZorpPay/backend/internal/api/routes"
	"github.com/VictorzllDev/ZorpPay/backend/internal/api/service"
	"github.com/VictorzllDev/ZorpPay/backend/internal/config"
	"github.com/VictorzllDev/ZorpPay/backend/internal/database"
	"github.com/VictorzllDev/ZorpPay/backend/internal/pkg/security"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// @title           ZorpPay
// @version         0.6.0
// @description     ZorpPay API Documentation

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type: Bearer token. Example: "Bearer your_jwt_token_here"

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
	authRepo := repository.NewAuthRepository(db)
	authService := service.NewAuthService(authRepo)
	authHandler := handler.NewAuthHandler(authService)

	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	dayRepo := repository.NewDayRepository(db)
	dayService := service.NewDayService(dayRepo)
	dayHandler := handler.NewDayHandler(dayService)

	paymentRepo := repository.NewPaymentRepository(db)
	paymentService := service.NewPaymentService(paymentRepo)
	paymentHandler := handler.NewPaymentHandler(paymentService)

	// JWT
	jwtService := security.NewJWT()

	// Routes
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
	}))
	routes.AuthRoutes(r, authHandler)
	routes.UserRoutes(r, userHandler, jwtService)
	routes.DayRoutes(r, dayHandler, jwtService)
	routes.PaymentRoutes(r, paymentHandler, jwtService)

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
