package controller

import (
	"cinema-app/internal/model"
	"cinema-app/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SeatController struct {
	seatService service.SeatService
}

func NewSeatController(seatService service.SeatService) *SeatController {
	return &SeatController{seatService}
}

func (cc *SeatController) GetAll(c *gin.Context) {
	data, err := cc.seatService.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, data)
}

func (cc *SeatController) GetByID(c *gin.Context) {
	idParam := c.Param("id")

	data, err := cc.seatService.GetByID(idParam)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "cinema not found"})
		return
	}

	c.JSON(http.StatusOK, data)
}

func (cc *SeatController) Create(c *gin.Context) {
	var req model.Seat
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := cc.seatService.CreateSeat(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, req)
}

func (cc *SeatController) Update(c *gin.Context) {
	idParam := c.Param("id")

	var req model.Seat
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := cc.seatService.UpdateSeat(idParam, &req); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Seat updated"})
}

func (cc *SeatController) Delete(c *gin.Context) {
	idParam := c.Param("id")

	if err := cc.seatService.DeleteSeat(idParam); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Seat deleted"})
}
