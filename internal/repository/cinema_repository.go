package repository

import (
	cinema "cinema-app/internal/model"

	"gorm.io/gorm"
)

type CinemaRepository interface {
	GetAll() ([]cinema.CinemaBranch, error)
	GetByID(id string) (cinema.CinemaBranch, error)
	Create(c *cinema.CinemaBranch) error
	Update(c *cinema.CinemaBranch) error
	Delete(id string) error
}

type cinemaRepository struct {
	db *gorm.DB
}

func NewCinemaRepository(db *gorm.DB) CinemaRepository {
	return &cinemaRepository{db}
}

func (r *cinemaRepository) GetAll() ([]cinema.CinemaBranch, error) {
	var cinemas []cinema.CinemaBranch
	err := r.db.Find(&cinemas).Error
	return cinemas, err
}

func (r *cinemaRepository) GetByID(id string) (cinema.CinemaBranch, error) {
	var c cinema.CinemaBranch
	err := r.db.First(&c, id).Error
	return c, err
}

func (r *cinemaRepository) Create(c *cinema.CinemaBranch) error {
	return r.db.Create(c).Error
}

func (r *cinemaRepository) Update(c *cinema.CinemaBranch) error {
	return r.db.Save(c).Error
}

func (r *cinemaRepository) Delete(id string) error {
	return r.db.Delete(&cinema.CinemaBranch{}, id).Error
}
