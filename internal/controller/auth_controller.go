package controller

import (
	"cinema-app/internal/dto"
	"cinema-app/internal/service"
	"cinema-app/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	service *service.AuthService
}

func NewAuthController(s *service.AuthService) *AuthController {
	return &AuthController{s}
}

// Register godoc
// @Description
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body dto.RegisterRequest false "Register Request"
// @Success 200 {object} dto.RegisterResponse
// @Router /register [post]
func (h AuthController) Register(c *gin.Context) {
	var req dto.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": utils.FormatValidationError(err),
		})
		return
	}

	resp, err := h.service.Register(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Registration successful",
		"data":    resp,
	})
}

type LoginResponse struct {
	Message string `json:"message"`
	Token   string `json:"token"`
}

// Login godoc
// @Description
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body dto.LoginRequest false "Login Request"
// @Success 200 {object} LoginResponse
// @Router /login [post]
func (h AuthController) Login(c *gin.Context) {
	var req dto.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": utils.FormatValidationError(err),
		})
		return
	}

	token, err := h.service.Login(&req)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": utils.FormatValidationError(err),
		})
		return
	}

	c.JSON(http.StatusOK, LoginResponse{
		Message: "Success",
		Token:   token,
	})

}
