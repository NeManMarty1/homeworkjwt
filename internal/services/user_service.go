package services

import (
	"errors"
	"homeworkjwt/internal/models"
	"homeworkjwt/internal/pgdb"
	"homeworkjwt/internal/utils"
)

var (
	ErrUserAlreadyExists  = errors.New("user already exists")
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrUserNotFound       = errors.New("user not found")
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
		return models.User{}, ErrUserAlreadyExists
	}

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return models.User{}, err
	}
	user.Password = hashedPassword

	createdUser, err := s.Repositories.User.Create(user)
	if err != nil {
		return models.User{}, err
	}

	return createdUser, nil
}

func (s *UserService) Login(email, password string) (string, error) {
	user, err := s.Repositories.User.FindByEmail(email)
	if err != nil {
		return "", ErrInvalidCredentials
	}

	if !utils.CheckPassword(password, user.Password) {
		return "", ErrInvalidCredentials
	}

	token, err := utils.GenerateJWT(user.ID, utils.JWTSecret)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *UserService) GetByID(id int) (models.User, error) {
	user, err := s.Repositories.User.FindByID(id)
	if err != nil {
		return models.User{}, ErrUserNotFound
	}
	return user, nil
}
