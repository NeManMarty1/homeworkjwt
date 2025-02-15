package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	// "github.com/jackc/pgx/v5"

	"homeworkjwt/internal/config"
)

type Postgres struct {
	// maxPoolSize int
	Pool *pgxpool.Pool
}

func New(cfg *config.Config) *Postgres {
	url := fmt.Sprintf("postgres://%s:%s@%s/%s&sslmode=disable",
	 	cfg.Postgres.User,
		cfg.Postgres.Password,
		cfg.Postgres.HostPort,
		cfg.Postgres.Database,
	)

	pg := &Postgres{}

	poolConfig, _ := pgxpool.ParseConfig(url)

	pg.Pool, _ = pgxpool.NewWithConfig(context.Background(), poolConfig)

	return pg	
}
