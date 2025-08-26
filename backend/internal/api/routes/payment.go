package routes

import (
	"github.com/VictorzllDev/ZorpPay/backend/internal/api/handler"
	"github.com/VictorzllDev/ZorpPay/backend/internal/api/middleware"
	"github.com/VictorzllDev/ZorpPay/backend/internal/pkg/security"
	"github.com/gin-gonic/gin"
)

func PaymentRoutes(r *gin.Engine, h *handler.PaymentHandler, jwtService security.JWTService) {
	dayGroup := r.Group("/api/v1/payments")
	dayGroup.Use(middleware.AuthMiddleware(jwtService))
	{
		dayGroup.POST("", h.CreatePayment)
	}
}
