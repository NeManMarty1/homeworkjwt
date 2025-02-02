package repository

import (
	"errors"
	"homeworkjwt/internal/models"
)

type UserRepository struct {
	users  []models.User
	nextID int
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		users:  []models.User{},
		nextID: 1,
	}
}

func (r *UserRepository) Create(user models.User) models.User {
	user.ID = r.nextID
	r.nextID++
	r.users = append(r.users, user)
	return user
}

func (r *UserRepository) FindByEmail(email string) (models.User, error) {
	for _, user := range r.users {
		if user.Email == email {
			return user, nil
		}
	}
	return models.User{}, errors.New("user not found")
}

func (r *UserRepository) FindByID(id int) (models.User, error) {
	for _, user := range r.users {
		if user.ID == id {
			return user, nil
		}
	}
	return models.User{}, errors.New("user not found")
}
