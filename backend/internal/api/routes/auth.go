package routes

import (
	"github.com/VictorzllDev/ZorpPay/backend/internal/api/handler"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.Engine, h *handler.AuthHandler) {
	userGroup := r.Group("/api/v1/auth")
	{
		userGroup.POST("/signup", h.SignUp)
		userGroup.POST("/signin", h.SignIn)
	}
}
