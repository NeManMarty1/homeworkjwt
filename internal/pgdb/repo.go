package pgdb

import (
	"homeworkjwt/internal/pgdb/impl"
	"homeworkjwt/internal/postgres"
)

type Repositories struct {
	User *impl.UserRepo
}

func NewRepositries(pg *postgres.Postgres) *Repositories {
	return &Repositories{
		User: impl.NewUserRepo(pg),
	}
}
