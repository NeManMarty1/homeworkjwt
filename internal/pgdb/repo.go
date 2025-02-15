package pgdb

import (
	"homeworkjwt/internal/postgres"
	"homeworkjwt/internal/pgdb/impl"
)

type Repositories struct {
	impl.UserRepo
}

func NewRepositries(pg *postgres.Postgres) *Repositories {
	return &Repositories{
		User: impl.NewUserRepo(pg),
	}
}
