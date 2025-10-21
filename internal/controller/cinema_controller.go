package controller

import (
	cinema "cinema-app/internal/model"
	"cinema-app/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CinemaController struct {
	cinemaService service.CinemaService
}

func NewCinemaController(cinemaService service.CinemaService) *CinemaController {
	return &CinemaController{cinemaService}
}

func (cc *CinemaController) GetAll(c *gin.Context) {
	data, err := cc.cinemaService.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, data)
}

func (cc *CinemaController) GetByID(c *gin.Context) {
	idParam := c.Param("id")

	data, err := cc.cinemaService.GetByID(idParam)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "cinema not found"})
		return
	}

	c.JSON(http.StatusOK, data)
}

func (cc *CinemaController) Create(c *gin.Context) {
	var req cinema.CinemaBranch
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := cc.cinemaService.CreateCinema(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, req)
}

func (cc *CinemaController) Update(c *gin.Context) {
	idParam := c.Param("id")

	var req cinema.CinemaBranch
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := cc.cinemaService.UpdateCinema(idParam, &req); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "cinema updated"})
}

func (cc *CinemaController) Delete(c *gin.Context) {
	idParam := c.Param("id")

	if err := cc.cinemaService.DeleteCinema(idParam); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "cinema deleted"})
}
