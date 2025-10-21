package service

import (
	"cinema-app/internal/model"
	"cinema-app/internal/repository"
	"cinema-app/internal/utils"

	"errors"
)

type SeatService interface {
	GetAll() ([]model.Seat, error)
	GetByID(id string) (model.Seat, error)
	CreateSeat(c *model.Seat) error
	UpdateSeat(id string, c *model.Seat) error
	DeleteSeat(id string) error
}

type seatService struct {
	repo repository.SeatRepository
}

func NewSeatService(repo repository.SeatRepository) SeatService {
	return &seatService{repo}
}

func (s *seatService) GetAll() ([]model.Seat, error) {
	return s.repo.GetAll()
}

func (s *seatService) GetByID(id string) (model.Seat, error) {
	return s.repo.GetByID(id)
}

func (s *seatService) CreateSeat(c *model.Seat) error {
	c.ID = utils.GenerateUUID()
	if c.SeatNumber == "" || c.ShowtimeID == "" {
		return errors.New("showtime and seat number are required")
	}
	return s.repo.Create(c)
}

func (s *seatService) UpdateSeat(id string, updated *model.Seat) error {
	c, err := s.repo.GetByID(id)
	if err != nil {
		return errors.New("seat not found")
	}

	c.ShowtimeID = updated.ShowtimeID
	c.SeatNumber = updated.SeatNumber
	c.IsBooked = updated.IsBooked
	return s.repo.Update(&c)
}

func (s *seatService) DeleteSeat(id string) error {
	return s.repo.Delete(id)
}
