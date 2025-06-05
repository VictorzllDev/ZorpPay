package routes

import (
	"github.com/VictorzllDev/ZorpPay/backend/internal/api/handler"
	"github.com/gin-gonic/gin"
)

func PaymentRoutes(r *gin.Engine, h *handler.PaymentHandler) {
	dayGroup := r.Group("/api/v1/payments")
	{
		dayGroup.POST("", h.CreatePayment)
	}
}
