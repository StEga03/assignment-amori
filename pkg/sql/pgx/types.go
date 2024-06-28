package pgx

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"time"
)

type Option func(*pgxpool.Config)

type DBConfig struct {
	DSN string `json:"dsn" yaml:"dsn"`

	MaxOpenConns    int           `json:"max_open_conns" yaml:"max_open_conns"`
	MaxIdleConns    int           `json:"max_idle_conns" yaml:"max_idle_conns"`
	ConnMaxLifetime time.Duration `json:"conn_max_lifetime" yaml:"conn_max"`
	ConnMaxIdleTime time.Duration `json:"conn_max_idle_time" yaml:"conn_max_idle_time"`
}

type Config struct {
	Master   DBConfig `json:"master" yaml:"master"`
	Follower DBConfig `json:"follower" yaml:"follower"`
}

type DB struct {
	Master
	master *pgxpool.Pool
}

type UpdateOptions struct {
	Set []string
	By  []string
}

type txInterface interface {
	Begin(ctx context.Context) (pgx.Tx, error)

	Commit(ctx context.Context) error

	Rollback(ctx context.Context) error

	Exec(ctx context.Context, sql string, arguments ...any) (commandTag pgconn.CommandTag, err error)
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row

	Conn() *pgx.Conn
}

type QueryTracer interface {
	TraceQueryStart(ctx context.Context, conn *pgx.Conn, data pgx.TraceQueryStartData) context.Context
	TraceQueryEnd(ctx context.Context, conn *pgx.Conn, data pgx.TraceQueryEndData)
}

type Tx struct {
	txInterface
}

type Master interface {
	Exec(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
}
