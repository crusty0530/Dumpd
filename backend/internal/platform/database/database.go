package database

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

// Connect to the database from the url and verifies it with a ping.
func Connect(ctx context.Context, dbURL string) (*pgxpool.Pool, error) {
	pool, err := pgxpool.New(ctx, dbURL)

	if err != nil {
		return nil, fmt.Errorf("creating the pool: %w", err)
	}

	err = pool.Ping(ctx)

	if err != nil {
		pool.Close()
		return nil, fmt.Errorf("pinging the pool: %w", err)
	}

	return pool, nil
}
