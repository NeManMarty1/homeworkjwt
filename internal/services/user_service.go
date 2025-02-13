package services

import (
	"errors"
	"homeworkjwt/internal/models"
	"homeworkjwt/internal/repository"
	"homeworkjwt/internal/utils"
	"os"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) Register(user models.User) (models.User, error) {
	_, err := s.repo.FindByEmail(user.Email)
	if err == nil {
		return models.User{}, errors.New("email already registered")
	}

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return models.User{}, err
	}
	user.Password = hashedPassword

	return s.repo.Create(user), nil
}

func (s *UserService) Login(email, password string) (string, error) {
	user, err := s.repo.FindByEmail(email)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	if !utils.CheckPassword(password, user.Password) {
		return "", errors.New("invalid credentials")
	}

	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		return "", errors.New("JWT secret is not set")
	}

	token, err := utils.GenerateJWT(user.ID, jwtSecret)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *UserService) GetByID(id int) (models.User, error) {
	return s.repo.FindByID(id)
}
