package handler

import (
	"github.com/VictorzllDev/ZorpPay/backend/internal/api/dto/request"
	"github.com/VictorzllDev/ZorpPay/backend/internal/api/dto/response"
	"github.com/VictorzllDev/ZorpPay/backend/internal/api/service"
	"github.com/VictorzllDev/ZorpPay/backend/internal/domain/entities"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserHandler struct {
	service service.UserService
}

func NewUserHandler(service service.UserService) *UserHandler {
	return &UserHandler{service: service}
}

// @Summary Create a new user
// @Description Create a new CreateUser
// @Tags User
// @Accept json
// @Produce json
// @Param CreateUser body request.CreateUser true "CreateUser"
// @Success 201 {object} response.User
// @Router /users [post]
func (h *UserHandler) CreateUser(c *gin.Context) {
	var req request.CreateUser

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := entities.User{
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

// @Summary Get all users
// @Description Get all users
// @Tags User
// @Accept json
// @Produce json
// @Success 200 {array} response.User
// @Router /users [get]
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
