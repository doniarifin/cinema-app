package repository

import (
	"cinema-app/internal/dto"
	"cinema-app/internal/model"

	"gorm.io/gorm"
)

type SeatRepository interface {
	GetAll() ([]model.Seat, error)
	GetByID(id string) (*model.Seat, error)
	GetSeat(filter dto.Seat) (*[]model.Seat, error)
	Create(c *model.Seat) error
	Update(c *model.Seat) error
	UpdateMany(c []model.Seat, u model.Seat) error
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

func (r *seatRepository) GetSeat(filter dto.Seat) (*[]model.Seat, error) {
	var seats []model.Seat
	query := r.db.Model(&model.Seat{})

	if filter.ID != "" {
		query = query.Where("id = ?", filter.ID)
	}
	if filter.IDs != nil {
		query = query.Where("id IN ?", filter.IDs)
	}
	if filter.ShowtimeID != "" {
		query = query.Where("showtime_id = ?", filter.ShowtimeID)
	}
	if filter.SeatNumber != "" {
		query = query.Where("seat_number = ?", filter.SeatNumber)
	}
	if filter.Status != "" {
		query = query.Where("status = ?", filter.Status)
	}

	if filter.IsBooked != nil {
		query = query.Where("is_booked = ?", *filter.IsBooked)
	}

	err := query.Find(&seats).Error
	return &seats, err

}

func (r *seatRepository) Create(c *model.Seat) error {
	return r.db.Create(c).Error
}

func (r *seatRepository) Update(c *model.Seat) error {
	return r.db.Save(c).Error
}

func (r *seatRepository) UpdateMany(seats []model.Seat, updated model.Seat) error {
	if len(seats) == 0 {
		return nil
	}

	updateData := map[string]any{
		"status":    updated.Status,
		"is_booked": updated.IsBooked,
	}

	return r.db.Model(seats).
		Updates(updateData).Error
}

func (r *seatRepository) Delete(id string) error {
	return r.db.Delete(&model.Seat{}, id).Error
}
