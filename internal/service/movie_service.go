package service

import (
	"cinema-app/internal/model"
	"cinema-app/internal/repository"
	"cinema-app/internal/utils"

	"errors"
)

type MovieService interface {
	GetAll() ([]model.Movie, error)
	GetByID(id string) (*model.Movie, error)
	CreateMovie(c *model.Movie) error
	UpdateMovie(id string, c *model.Movie) error
	DeleteMovie(id string) error
}

type movieService struct {
	repo repository.MovieRepository
}

func NewMovieService(repo repository.MovieRepository) MovieService {
	return &movieService{repo}
}

func (s *movieService) GetAll() ([]model.Movie, error) {
	return s.repo.GetAll()
}

func (s *movieService) GetByID(id string) (*model.Movie, error) {
	return s.repo.GetByID(id)
}

func (s *movieService) CreateMovie(c *model.Movie) error {
	c.ID = utils.GenerateUUID()
	if c.Title == "" {
		return errors.New("name is required")
	}
	return s.repo.Create(c)
}

func (s *movieService) UpdateMovie(id string, updated *model.Movie) error {
	c, err := s.repo.GetByID(id)
	if err != nil {
		return errors.New("movie not found")
	}

	c.Title = updated.Title
	c.Genre = updated.Genre
	c.Duration = updated.Duration
	c.Synopsis = updated.Synopsis
	return s.repo.Update(c)
}

func (s *movieService) DeleteMovie(id string) error {
	return s.repo.Delete(id)
}
