package impl

import (
	"context"

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

func (r *UserRepo) Create(user models.User) (models.User, error) {
	ctx := context.Background()

	query := `INSERT INTO users (name, email, password) VALUES ($1, $2, $3) RETURNING id`

	var id int
	err := r.pg.Pool.QueryRow(ctx, query, user.Name, user.Email, user.Password).Scan(&id)
	if err != nil {
		return models.User{}, err
	}

	user.ID = id
	return user, nil
}

func (r *UserRepo) FindByEmail(email string) (models.User, error) {
	ctx := context.Background()

	query := `SELECT id, name, email, password FROM users WHERE email = $1`

	var user models.User
	err := r.pg.Pool.QueryRow(ctx, query, email).Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (r *UserRepo) FindByID(id int) (models.User, error) {
	ctx := context.Background()

	query := `SELECT id, name, email, password FROM users WHERE id = $1`

	var user models.User
	err := r.pg.Pool.QueryRow(ctx, query, id).Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}
