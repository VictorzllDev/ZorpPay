package handler

import (
	"github.com/VictorzllDev/ZorpPay/backend/internal/api/dto/request"
	"github.com/VictorzllDev/ZorpPay/backend/internal/api/dto/response"
	"github.com/VictorzllDev/ZorpPay/backend/internal/api/service"
	"github.com/VictorzllDev/ZorpPay/backend/internal/domain/entities"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AuthHandler struct {
	service service.AuthService
}

func NewAuthHandler(service service.AuthService) *AuthHandler {
	return &AuthHandler{service: service}
}

// @Summary Register new user
// @Description Creates a new user account in the system
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body request.SignUp true "Registration data"
// @Success 201 {object} response.Auth "Account created successfully"
// @Router /auth/signup [post]
func (h *AuthHandler) SignUp(c *gin.Context) {
	req := &request.SignUp{}

	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userToSignUp := &entities.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}

	if err := h.service.SignUp(userToSignUp); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	userToSignIn := &entities.User{
		Email:    req.Email,
		Password: req.Password,
	}

	token, err := h.service.SignIn(userToSignIn)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	res := &response.Auth{
		Token: token.Token,
	}

	c.JSON(http.StatusCreated, res)
}

// @Summary Authenticate user
// @Description Performs user login and returns access token
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body request.SignIn true "Login credentials"
// @Success 200 {object} response.Auth "Login successful"
// @Router /auth/signin [post]
func (h *AuthHandler) SignIn(c *gin.Context) {
	req := &request.SignIn{}

	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := &entities.User{
		Email:    req.Email,
		Password: req.Password,
	}

	token, err := h.service.SignIn(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	res := &response.Auth{
		Token: token.Token,
	}

	c.JSON(http.StatusOK, res)
}
