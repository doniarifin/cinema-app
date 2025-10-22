package service

import (
	"cinema-app/internal/dto"
	"cinema-app/internal/model"
	"cinema-app/internal/repository"
	"cinema-app/internal/utils"

	"errors"
)

type SeatTrxSrv interface {
	GetAll() ([]model.SeatTransaction, error)
	GetByID(id string) (*model.SeatTransaction, error)
	GetsSeatTrx(f dto.SeatTrx) (*[]model.SeatTransaction, error)
	CreateSeatTrx(c *model.SeatTransaction) error
	UpdateSeatTrx(id string, c *model.SeatTransaction) error
	DeleteSeatTrx(id string) error
}

type seatTrxSrv struct {
	repo repository.SeatTransactionRepo
}

func NewSeatTrxSrv(repo repository.SeatTransactionRepo) SeatTrxSrv {
	return &seatTrxSrv{repo}
}

func (s *seatTrxSrv) GetAll() ([]model.SeatTransaction, error) {
	return s.repo.GetAll()
}

func (s *seatTrxSrv) GetByID(id string) (*model.SeatTransaction, error) {
	return s.repo.GetByID(id)
}

func (s *seatTrxSrv) GetsSeatTrx(filter dto.SeatTrx) (*[]model.SeatTransaction, error) {
	return s.repo.GetsSeatTrx(filter)
}

func (s *seatTrxSrv) CreateSeatTrx(c *model.SeatTransaction) error {
	if c.ID == "" {
		c.ID = utils.GenerateUUID()
	}
	return s.repo.Create(c)
}

func (s *seatTrxSrv) UpdateSeatTrx(id string, updated *model.SeatTransaction) error {
	c, err := s.repo.GetByID(id)
	if err != nil {
		return errors.New("SeatTrx not found")
	}

	c.SeatID = updated.SeatID
	c.ShowtimeID = updated.ShowtimeID
	c.Status = updated.Status
	c.TransactionID = updated.TransactionID
	return s.repo.Update(c)
}

func (s *seatTrxSrv) DeleteSeatTrx(id string) error {
	return s.repo.Delete(id)
}
