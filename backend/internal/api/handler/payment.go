package handler

import (
	"net/http"
	"strings"
	"time"

	"github.com/VictorzllDev/ZorpPay/backend/internal/api/dto/request"
	"github.com/VictorzllDev/ZorpPay/backend/internal/api/dto/response"
	"github.com/VictorzllDev/ZorpPay/backend/internal/api/service"
	"github.com/VictorzllDev/ZorpPay/backend/internal/domain/entities"
	"github.com/gin-gonic/gin"
)

type PaymentHandler struct {
	service service.PaymentService
}

func NewPaymentHandler(service service.PaymentService) *PaymentHandler {
	return &PaymentHandler{service: service}
}

// @Summary Create a new Payment
// @Description Create a new Payment, associating it to a user
// @Tags Payment
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param Payment body request.CreatePayment true "Payment"
// @Success 201 {object} response.Payment
// @Router /payments [post]
func (h *PaymentHandler) CreatePayment(c *gin.Context) {
	var req request.CreatePayment

	// Validação do payload
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "formato de requisição inválido",
			"details": err.Error(),
		})
		return
	}

	// Converter data (com tratamento de erro)
	parsedPaymentDate, err := time.Parse("02-01-2006", req.PaymentDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":           "formato de data inválido",
			"expected_format": "DD-MM-YYYY",
		})
		return
	}

	// Criar entidade
	payment := entities.Payment{
		PaymentDate: parsedPaymentDate,
		UserID:      req.UserID,
		Amount:      req.Amount,
	}

	// Chamar service
	if err := h.service.CreatePayment(&payment); err != nil {
		// Determinar status code apropriado
		statusCode := http.StatusInternalServerError
		if strings.Contains(err.Error(), "já registrado") {
			statusCode = http.StatusConflict
		}

		c.JSON(statusCode, gin.H{
			"error":   "não foi possível registrar o pagamento",
			"details": err.Error(),
		})
		return
	}

	// Formatar resposta
	res := response.Payment{
		ID:   payment.ID,
		Date: payment.PaymentDate.Format("02-01-2006"), // Formatar data
	}

	c.JSON(http.StatusCreated, res)

}
