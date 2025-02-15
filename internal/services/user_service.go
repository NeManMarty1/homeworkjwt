package services

import (
	"errors"
	"homeworkjwt/internal/models"

	// "homeworkjwt/internal/repository"
	"homeworkjwt/internal/pgdb"
	"homeworkjwt/internal/utils"
)

type UserService struct {
	Repositories *pgdb.Repositories
}

func NewUserService(repositories *pgdb.Repositories) *UserService {
	return &UserService{
		Repositories: repositories,
	}
}

func (s *UserService) Register(user models.User) (models.User, error) {
	_, err := s.Repositories.User.FindByEmail(user.Email)
	if err == nil {
		return models.User{}, errors.New("email already registered")
	}

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return models.User{}, err
	}
	user.Password = hashedPassword

	return s.Repositories.User.Create(user)
}

func (s *UserService) Login(email, password string) (string, error) {
	user, err := s.Repositories.User.FindByEmail(email)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	if !utils.CheckPassword(password, user.Password) {
		return "", errors.New("invalid credentials")
	}

	token, err := utils.GenerateJWT(user.ID, utils.JWTSecret)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *UserService) GetByID(id int) (models.User, error) {
	return s.Repositories.User.FindByID(id)
}
