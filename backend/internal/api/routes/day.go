package routes

import (
	"github.com/VictorzllDev/ZorpPay/backend/internal/api/handler"
	"github.com/VictorzllDev/ZorpPay/backend/internal/api/middleware"
	"github.com/VictorzllDev/ZorpPay/backend/internal/pkg/security"
	"github.com/gin-gonic/gin"
)

func DayRoutes(r *gin.Engine, h *handler.DayHandler, jwtService security.JWTService) {
	dayGroup := r.Group("/api/v1/days")
	dayGroup.Use(middleware.AuthMiddleware(jwtService))
	{
		dayGroup.POST("", h.CreateDay)
	}
}
