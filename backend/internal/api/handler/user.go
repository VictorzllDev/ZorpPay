package handler

import (
	"net/http"

	"github.com/VictorzllDev/ZorpPay/backend/internal/api/dto/request"
	"github.com/VictorzllDev/ZorpPay/backend/internal/api/dto/response"
	"github.com/VictorzllDev/ZorpPay/backend/internal/api/service"
	"github.com/VictorzllDev/ZorpPay/backend/internal/domain/entities"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service service.UserService
}

type Config struct {
	R *gin.Engine
}

func NewUserHandler(service service.UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var req request.CreateUser

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := domain.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}

	if err := h.service.CreateUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	res := response.User{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}

	c.JSON(http.StatusCreated, res)
}

func (h *UserHandler) GetUser(c *gin.Context) {
	users, err := h.service.GetAllUser()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	res := []response.User{}
	for _, user := range users {
		res = append(res, response.User{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
		})
	}
	c.JSON(http.StatusOK, res)
}
