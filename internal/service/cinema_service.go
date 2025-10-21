package service

import (
	cinema "cinema-app/internal/model"
	"cinema-app/internal/repository"
	"cinema-app/internal/utils"

	"errors"
)

type CinemaService interface {
	GetAll() ([]cinema.CinemaBranch, error)
	GetByID(id string) (cinema.CinemaBranch, error)
	CreateCinema(c *cinema.CinemaBranch) error
	UpdateCinema(id string, c *cinema.CinemaBranch) error
	DeleteCinema(id string) error
}

type cinemaService struct {
	repo repository.CinemaRepository
}

func NewCinemaService(repo repository.CinemaRepository) CinemaService {
	return &cinemaService{repo}
}

func (s *cinemaService) GetAll() ([]cinema.CinemaBranch, error) {
	return s.repo.GetAll()
}

func (s *cinemaService) GetByID(id string) (cinema.CinemaBranch, error) {
	return s.repo.GetByID(id)
}

func (s *cinemaService) CreateCinema(c *cinema.CinemaBranch) error {
	c.ID = utils.GenerateUUID()
	if c.BranchName == "" {
		return errors.New("name is required")
	}
	return s.repo.Create(c)
}

func (s *cinemaService) UpdateCinema(id string, updated *cinema.CinemaBranch) error {
	c, err := s.repo.GetByID(id)
	if err != nil {
		return errors.New("cinema not found")
	}

	c.BranchName = updated.BranchName
	c.City = updated.City
	return s.repo.Update(&c)
}

func (s *cinemaService) DeleteCinema(id string) error {
	return s.repo.Delete(id)
}
