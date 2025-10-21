package repository

import (
	"cinema-app/internal/model"

	"gorm.io/gorm"
)

type MovieRepository interface {
	GetAll() ([]model.Movie, error)
	GetByID(id string) (model.Movie, error)
	Create(c *model.Movie) error
	Update(c *model.Movie) error
	Delete(id string) error
}

type movieRepository struct {
	db *gorm.DB
}

func NewMovieRepository(db *gorm.DB) MovieRepository {
	return &movieRepository{db}
}

func (r *movieRepository) GetAll() ([]model.Movie, error) {
	var movies []model.Movie
	err := r.db.Find(&movies).Error
	return movies, err
}

func (r *movieRepository) GetByID(id string) (model.Movie, error) {
	var c model.Movie
	err := r.db.First(&c, id).Error
	return c, err
}

func (r *movieRepository) Create(c *model.Movie) error {
	return r.db.Create(c).Error
}

func (r *movieRepository) Update(c *model.Movie) error {
	return r.db.Save(c).Error
}

func (r *movieRepository) Delete(id string) error {
	return r.db.Delete(&model.Movie{}, id).Error
}
