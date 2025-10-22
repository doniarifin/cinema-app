package service

import (
	"cinema-app/internal/dto"
	"cinema-app/internal/model"
	"cinema-app/internal/repository"
	"cinema-app/internal/utils"

	"errors"
)

type SeatService interface {
	GetAll() ([]model.Seat, error)
	GetByID(id string) (*model.Seat, error)
	GetSeat(f dto.Seat) (*[]model.Seat, error)
	CreateSeat(c *model.Seat) error
	UpdateSeat(id string, c *model.Seat) error
	UpdateSeats(f dto.Seat, trxID string, u *model.Seat, s *model.SeatTransaction) error
	DeleteSeat(id string) error
}

type seatService struct {
	repo        repository.SeatRepository
	repoSeatTrx repository.SeatTransactionRepo
}

func NewSeatService(r repository.SeatRepository, rst repository.SeatTransactionRepo) SeatService {
	return &seatService{
		repo:        r,
		repoSeatTrx: rst,
	}
}

func (s *seatService) GetAll() ([]model.Seat, error) {
	return s.repo.GetAll()
}

func (s *seatService) GetByID(id string) (*model.Seat, error) {
	return s.repo.GetByID(id)
}

func (s *seatService) GetSeat(filter dto.Seat) (*[]model.Seat, error) {
	return s.repo.GetSeat(filter)
}

func (s *seatService) CreateSeat(c *model.Seat) error {
	c.ID = utils.GenerateUUID()
	if c.SeatNumber == "" || c.BranchID == "" {
		return errors.New("branch and seat number are required")
	}
	return s.repo.Create(c)
}

func (s *seatService) UpdateSeat(id string, updated *model.Seat) error {
	c, err := s.repo.GetByID(id)
	if err != nil {
		return errors.New("seat not found")
	}

	c.BranchID = updated.BranchID
	// c.ShowtimeID = updated.ShowtimeID
	c.SeatNumber = updated.SeatNumber
	c.Status = updated.Status
	c.IsBooked = updated.IsBooked
	return s.repo.Update(c)
}

func (s *seatService) UpdateSeats(f dto.Seat, trxId string, updated *model.Seat, updateStrx *model.SeatTransaction) error {
	if err := s.repoSeatTrx.UpdateMany(trxId, updateStrx); err != nil {
		return errors.New("internal server error: failed updatemany seats")
	}

	c, err := s.repo.GetSeat(f)
	if err != nil {
		return errors.New("internal server error: failed updatemany seats")
	}

	return s.repo.UpdateMany(*c, *updated)
}

func (s *seatService) DeleteSeat(id string) error {
	return s.repo.Delete(id)
}
