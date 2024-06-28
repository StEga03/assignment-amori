package pgx

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"time"
)

func parseConfig(config DBConfig, options ...Option) (*pgxpool.Config, error) {
	pgxConfig, err := pgxpool.ParseConfig(config.DSN)
	if err != nil {
		return nil, err
	}

	if config.MaxIdleConns > 0 {
		pgxConfig.MinConns = int32(config.MaxIdleConns)
	}

	if config.MaxOpenConns > 0 {
		pgxConfig.MaxConns = int32(config.MaxOpenConns)
	}

	if config.ConnMaxLifetime > time.Second {
		pgxConfig.MaxConnLifetime = config.ConnMaxLifetime
	}

	if config.ConnMaxIdleTime > time.Second {
		pgxConfig.MaxConnIdleTime = config.ConnMaxIdleTime
	}

	for _, opt := range options {
		opt(pgxConfig)
	}

	return pgxConfig, nil
}

func WithTracer(tracer QueryTracer) Option {
	return func(config *pgxpool.Config) {
		config.ConnConfig.Tracer = tracer
	}
}

func WithAfterConnect(f func(context.Context, *pgx.Conn) error) Option {
	return func(config *pgxpool.Config) {
		config.AfterConnect = f
	}
}
