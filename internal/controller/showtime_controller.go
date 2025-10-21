package controller

import (
	"cinema-app/internal/model"
	"cinema-app/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ShowtimeController struct {
	showtimeService service.ShowtimeService
}

func NewShowtimeController(showtimeService service.ShowtimeService) *ShowtimeController {
	return &ShowtimeController{showtimeService}
}

func (cc *ShowtimeController) GetAll(c *gin.Context) {
	data, err := cc.showtimeService.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, data)
}

func (cc *ShowtimeController) GetByID(c *gin.Context) {
	idParam := c.Param("id")

	data, err := cc.showtimeService.GetByID(idParam)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "cinema not found"})
		return
	}

	c.JSON(http.StatusOK, data)
}

func (cc *ShowtimeController) Create(c *gin.Context) {
	var req model.Showtime
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := cc.showtimeService.CreateShowtime(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, req)
}

func (cc *ShowtimeController) Update(c *gin.Context) {
	idParam := c.Param("id")

	var req model.Showtime
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := cc.showtimeService.UpdateShowtime(idParam, &req); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "showtime updated"})
}

func (cc *ShowtimeController) Delete(c *gin.Context) {
	idParam := c.Param("id")

	if err := cc.showtimeService.DeleteShowtime(idParam); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "showtime deleted"})
}
