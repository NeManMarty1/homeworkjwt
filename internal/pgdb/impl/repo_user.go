package impl

import (
	"homeworkjwt/internal/postgres"
	"homeworkjwt/internal/models"
)

type UserRepo struct {
	pg *postgres.Postgres
}

func NewUserRepo(pg *postgres.Postgres) *UserRepo {
	return &UserRepo{
		pg: pg,
	}
}

func (r *UserRepo) Create(user models.User) models.User {
	
}
