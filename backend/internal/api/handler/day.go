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

type DayHandler struct {
	service service.DayService
}

func NewDayHandler(service service.DayService) *DayHandler {
	return &DayHandler{service: service}
}

// @Summary Create a new day
// @Description Create a new Day, associating it to a user
// @Tags Day
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param Day body request.CreateDay true "Day"
// @Success 201 {object} response.Day
// @Router /days [post]
func (h *DayHandler) CreateDay(c *gin.Context) {
	var req request.CreateDay

	// Validação do payload
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "formato de requisição inválido",
			"details": err.Error(),
		})
		return
	}

	// Converter data (com tratamento de erro)
	parsedDate, err := time.Parse("02-01-2006", req.Date)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":           "formato de data inválido",
			"expected_format": "DD-MM-YYYY",
		})
		return
	}

	// Criar entidade
	day := entities.Day{
		Date:   parsedDate,
		UserID: req.UserID,
	}

	// Chamar service
	if err := h.service.CreateDay(&day); err != nil {
		// Determinar status code apropriado
		statusCode := http.StatusInternalServerError
		if strings.Contains(err.Error(), "já registrado") {
			statusCode = http.StatusConflict
		}

		c.JSON(statusCode, gin.H{
			"error":   "não foi possível registrar o dia",
			"details": err.Error(),
		})
		return
	}

	// Formatar resposta
	res := response.Day{
		ID:   day.ID,
		Date: day.Date.Format("02-01-2006"), // Formatar data
	}

	c.JSON(http.StatusCreated, res)
}
