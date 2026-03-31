package postgres

import (
	"booking-system/services/auth/internal/config"
	"context"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Storage struct {
	pool *pgxpool.Pool
}

func New(ctx context.Context, cfg config.Postgres) (*Storage, error) {
	const op = "storage.postgres.New"

	connString := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresDB,
	)

	conn, err := pgxpool.New(ctx, connString)
	if err != nil {
		return nil, fmt.Errorf("%s  unable to connect to database: %w", op, err)
	}

	if err := conn.Ping(ctx); err != nil {
		return nil, fmt.Errorf("%s database ping failed: %w",op, err)
	}

	m, err := migrate.New("file://db/migrations",
		fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
			cfg.PostgresUser,
			cfg.PostgresPassword,
			cfg.PostgresHost,
			cfg.PostgresPort,
			cfg.PostgresDB,
		))
	if err != nil {
		return nil, fmt.Errorf("%s  migration init failed: %w", op, err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return nil, fmt.Errorf("%s  migration up failed: %w", op, err)
	}

	return &Storage{pool: conn}, nil
}



