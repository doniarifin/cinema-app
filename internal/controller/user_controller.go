package controller

import (
	"cinema-app/internal/dto"
	model "cinema-app/internal/model"
	"cinema-app/internal/service"
	"cinema-app/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	service *service.UserService
}

func NewUserController(s *service.UserService) *UserController {
	return &UserController{s}
}

// @BasePath /api/v1
// GetAllUsers godoc
// @Security BearerAuth
// @Description get all users
// @Tags User
// @Accept json
// @Produce json
// @Success 200 {object} model.User
// @Router /api/v1/GetAllUsers [get]
func (h UserController) GetAllUsers(c *gin.Context) {
	user, err := h.service.Gets()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": utils.FormatValidationError(err),
		})
		return
	}
	c.JSON(http.StatusOK, user)
}

// @BasePath /api/v1
// GetCurrentUser godoc
// @Security BearerAuth
// @Description get current user
// @Tags User
// @Accept json
// @Produce json
// @Success 200 {object} model.User
// @Router /api/v1/GetCurrentUser [get]
func (h UserController) GetCurrentUser(c *gin.Context) {
	id := c.GetString("user_id")
	user, err := h.service.Get(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": utils.FormatValidationError(err),
		})
		return
	}
	c.JSON(http.StatusOK, user)
}

type UpdateResponse struct {
	Message string
	Data    dto.UserReponse
}

// @BasePath /api/v1
// UpdateUser godoc
// @Security BearerAuth
// @Description Update a user's email
// @Tags User
// @Accept json
// @Produce json
// @Param id path string true "ID"
// @Param user body dto.UserRequest true "User data"
// @Success 200 {object} dto.UserReponse
// @Router /api/v1/Update/{id} [put]
func (h UserController) Update(c *gin.Context) {
	//get user_id and role from jwt
	userID := c.GetString("user_id")
	userRole := c.GetString("role")
	//get param
	paramID := c.Param("id")

	var req dto.UserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": utils.FormatValidationError(err),
		})
		return
	}

	currentUser, err := h.service.Get(paramID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": utils.FormatValidationError(err),
		})
		return
	}

	//check autorization
	if userID != paramID && userRole != model.AdminRole.String() {
		c.JSON(http.StatusForbidden, gin.H{"error": "cannot update another user"})
		return
	}

	user := &model.User{
		ID:       currentUser.ID,
		Password: currentUser.Password,
		Role:     currentUser.Role,
		Email:    req.Email,
		Name:     req.Name,
	}

	if err := h.service.Update(user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": utils.FormatValidationError(err),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "update user success",
		"data":    user,
	})
}

type DeleteReponse struct {
	Message string `json:"message"`
}

// @BasePath /api/v1
// DeleteUserByID godoc
// @Security BearerAuth
// @Description Delete a user given their ID
// @Tags User
// @Accept json
// @Produce json
// @Param id path string true "ID"
// @Success 200 {object} DeleteReponse
// @Router /api/v1/Delete/{id} [delete]
func (h UserController) Delete(c *gin.Context) {
	id := c.Param("id")
	err := h.service.Delete(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": utils.FormatValidationError(err),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "delete user success",
	})
}
