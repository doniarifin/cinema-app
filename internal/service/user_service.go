package service

import (
	model "cinema-app/internal/model"
	"cinema-app/internal/repository"
)

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(r repository.UserRepository) *UserService {
	return &UserService{r}
}

func (s UserService) Gets() ([]model.User, error) {
	user, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s UserService) Get(id string) (*model.User, error) {
	return s.repo.FindByID(id)
}

func (s UserService) Update(m *model.User) error {
	return s.repo.Update(m)
}

func (s UserService) Delete(id string) error {
	ids := []string{}
	ids = append(ids, id)
	return s.repo.Delete(ids)
}

func (s UserService) DeleteMany(ids []string) error {
	return s.repo.Delete(ids)
}
