package service

import (
	"cinema-app/internal/model"
	"cinema-app/internal/repository"
	"cinema-app/internal/utils"

	"errors"
)

type ShowtimeService interface {
	GetAll() ([]model.Showtime, error)
	GetByID(id string) (model.Showtime, error)
	CreateShowtime(c *model.Showtime) error
	UpdateShowtime(id string, c *model.Showtime) error
	DeleteShowtime(id string) error
}

type showtimeService struct {
	repo repository.ShowtimeRepository
}

func NewShowtimeService(repo repository.ShowtimeRepository) ShowtimeService {
	return &showtimeService{repo}
}

func (s *showtimeService) GetAll() ([]model.Showtime, error) {
	return s.repo.GetAll()
}

func (s *showtimeService) GetByID(id string) (model.Showtime, error) {
	return s.repo.GetByID(id)
}

func (s *showtimeService) CreateShowtime(c *model.Showtime) error {
	c.ID = utils.GenerateUUID()
	if c.BranchID == "" || c.MovieID == "" {
		return errors.New("branch cinema and movie are required")
	}
	return s.repo.Create(c)
}

func (s *showtimeService) UpdateShowtime(id string, updated *model.Showtime) error {
	c, err := s.repo.GetByID(id)
	if err != nil {
		return errors.New("showtime not found")
	}

	c.BranchID = updated.BranchID
	c.MovieID = updated.MovieID
	c.Price = updated.Price
	c.DateTime = updated.DateTime
	return s.repo.Update(&c)
}

func (s *showtimeService) DeleteShowtime(id string) error {
	return s.repo.Delete(id)
}
