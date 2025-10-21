package service

import (
	"cinema-app/internal/dto"
	model "cinema-app/internal/model"
	"cinema-app/internal/pkg/jwt"
	"cinema-app/internal/repository"
	"cinema-app/internal/utils"
	"errors"
)

type AuthService struct {
	repo repository.UserRepository
}

func NewAuthService(r repository.UserRepository) *AuthService {
	return &AuthService{r}
}

func (s AuthService) Register(req *dto.RegisterRequest) (*dto.RegisterResponse, error) {
	hashedPassword, _ := utils.HashPassword(req.Password)

	user := &model.User{
		ID:       utils.GenerateUUID(),
		Name:     req.Name,
		Email:    req.Email,
		Password: hashedPassword,
	}
	if err := s.repo.Create(user); err != nil {
		return nil, err
	}

	return &dto.RegisterResponse{
		ID:    user.ID,
		Email: user.Email,
		Name:  user.Name,
	}, nil
}

func (s AuthService) Login(req *dto.LoginRequest) (string, error) {
	user, err := s.repo.FindByEmail(req.Email)
	if err != nil {
		return "", err
	}

	if !utils.CheckPasswordHash(req.Password, user.Password) {
		return "", errors.New("invalid email or password")
	}

	token, err := jwt.GenerateJWT(user)
	if err != nil {
		return "", errors.New("error generate token")
	}

	return token, nil
}
