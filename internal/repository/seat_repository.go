package repository

import (
	"cinema-app/internal/model"

	"gorm.io/gorm"
)

type SeatRepository interface {
	GetAll() ([]model.Seat, error)
	GetByID(id string) (*model.Seat, error)
	Create(c *model.Seat) error
	Update(c *model.Seat) error
	Delete(id string) error
}

type seatRepository struct {
	db *gorm.DB
}

func NewSeatRepository(db *gorm.DB) SeatRepository {
	return &seatRepository{db}
}

func (r *seatRepository) GetAll() ([]model.Seat, error) {
	var seats []model.Seat
	err := r.db.Find(&seats).Error
	return seats, err
}

func (r *seatRepository) GetByID(id string) (*model.Seat, error) {
	var c model.Seat
	if err := r.db.First(&c, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &c, nil
}

func (r *seatRepository) Create(c *model.Seat) error {
	return r.db.Create(c).Error
}

func (r *seatRepository) Update(c *model.Seat) error {
	return r.db.Save(c).Error
}

func (r *seatRepository) Delete(id string) error {
	return r.db.Delete(&model.Seat{}, id).Error
}
