package controller

import (
	"cinema-app/internal/model"
	"cinema-app/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MovieControlller struct {
	movieService service.MovieService
}

func NewMovieControlller(movieService service.MovieService) *MovieControlller {
	return &MovieControlller{movieService}
}

func (cc *MovieControlller) GetAll(c *gin.Context) {
	data, err := cc.movieService.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, data)
}

func (cc *MovieControlller) GetByID(c *gin.Context) {
	idParam := c.Param("id")

	data, err := cc.movieService.GetByID(idParam)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "cinema not found"})
		return
	}

	c.JSON(http.StatusOK, data)
}

func (cc *MovieControlller) Create(c *gin.Context) {
	var req model.Movie
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := cc.movieService.CreateMovie(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, req)
}

func (cc *MovieControlller) Update(c *gin.Context) {
	idParam := c.Param("id")

	var req model.Movie
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := cc.movieService.UpdateMovie(idParam, &req); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "movie updated"})
}

func (cc *MovieControlller) Delete(c *gin.Context) {
	idParam := c.Param("id")

	if err := cc.movieService.DeleteMovie(idParam); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "movie deleted"})
}
