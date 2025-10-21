package repository

import (
	"cinema-app/internal/model"

	"gorm.io/gorm"
)

type ShowtimeRepository interface {
	GetAll() ([]model.Showtime, error)
	GetByID(id string) (model.Showtime, error)
	Create(c *model.Showtime) error
	Update(c *model.Showtime) error
	Delete(id string) error
}

type showtimeRepository struct {
	db *gorm.DB
}

func NewShowtimeRepository(db *gorm.DB) ShowtimeRepository {
	return &showtimeRepository{db}
}

func (r *showtimeRepository) GetAll() ([]model.Showtime, error) {
	var showtimes []model.Showtime
	err := r.db.Find(&showtimes).Error
	return showtimes, err
}

func (r *showtimeRepository) GetByID(id string) (model.Showtime, error) {
	var c model.Showtime
	err := r.db.First(&c, id).Error
	return c, err
}

func (r *showtimeRepository) Create(c *model.Showtime) error {
	return r.db.Create(c).Error
}

func (r *showtimeRepository) Update(c *model.Showtime) error {
	return r.db.Save(c).Error
}

func (r *showtimeRepository) Delete(id string) error {
	return r.db.Delete(&model.Showtime{}, id).Error
}
