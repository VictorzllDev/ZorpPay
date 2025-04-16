package routes

import (
	"github.com/VictorzllDev/ZorpPay/backend/internal/api/handler"
	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.Engine, h *handler.UserHandler) {
	userGroup := r.Group("/users")
	{
		userGroup.POST("", h.CreateUser)
		userGroup.GET("", h.GetUser)
	}
}
