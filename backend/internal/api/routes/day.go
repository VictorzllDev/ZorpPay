package routes

import (
	"github.com/VictorzllDev/ZorpPay/backend/internal/api/handler"
	"github.com/gin-gonic/gin"
)

func DayRoutes(r *gin.Engine, h *handler.DayHandler) {
	dayGroup := r.Group("/api/v1/day")
	{
		dayGroup.POST("", h.CreateDay)
	}
}
