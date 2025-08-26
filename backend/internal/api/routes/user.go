package routes

import (
	"github.com/VictorzllDev/ZorpPay/backend/internal/api/handler"
	"github.com/VictorzllDev/ZorpPay/backend/internal/api/middleware"
	"github.com/VictorzllDev/ZorpPay/backend/internal/pkg/security"
	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.Engine, h *handler.UserHandler, jwtService security.JWTService) {
	userGroup := r.Group("/api/v1/users")
	userGroup.Use(middleware.AuthMiddleware(jwtService))
	{
		userGroup.POST("", h.CreateUser)
		userGroup.GET("", h.GetAllUser)
	}
}
