package database

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PoolConfig struct {
	MaxConns        int32         // Maximum number of connections in the pool
	MinConns        int32         // Minimum number of connections (reserve)
	MaxConnLifetime time.Duration // Maximum connection lifetime
	MaxConnIdleTime time.Duration // Maximum connection idle time before closing
}

func ConnectPGX(ctx context.Context, connectURL string, cfg *PoolConfig) (*pgxpool.Pool, error) {
	// Get the pool config from dsn
	poolConfig, err := pgxpool.ParseConfig(connectURL)
	if err != nil {
		return nil, fmt.Errorf("failed to parse DSN: %w", err)
	}

	poolConfig.MaxConns = cfg.MaxConns
	poolConfig.MinConns = cfg.MinConns
	poolConfig.MaxConnLifetime = cfg.MaxConnLifetime
	poolConfig.MaxConnIdleTime = cfg.MaxConnIdleTime

	pool, err := pgxpool.NewWithConfig(ctx, poolConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to create pgx pool: %w", err)
	}

	if err := pool.Ping(ctx); err != nil {
		pool.Close()
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return pool, nil
}
