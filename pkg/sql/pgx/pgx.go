package pgx

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/wcamarao/pmx"
)

// New is creating new connection to postgres using PGX driver.
func New(ctx context.Context, config DBConfig, options ...Option) (db *DB, err error) {
	dbConfig, err := parseConfig(config, options...)
	if err != nil {
		return nil, err
	}

	dbPool, err := pgxpool.NewWithConfig(ctx, dbConfig)
	if err != nil {
		return nil, err
	}

	return &DB{
		Master: dbPool,
		master: dbPool,
	}, nil
}

func (db *DB) Insert(ctx context.Context, entity interface{}) (pgconn.CommandTag, error) {
	return pmx.Insert(ctx, db.master, entity)
}

func (db *DB) Update(ctx context.Context, entity interface{}, options *UpdateOptions) (pgconn.CommandTag, error) {
	return pmx.Update(ctx, db.master, entity, (*pmx.UpdateOptions)(options))
}

func (db *DB) Select(ctx context.Context, dest interface{}, sql string, args ...interface{}) (bool, error) {
	err := pmx.Select(ctx, db.master, dest, sql, args...)
	if errors.Is(err, pgx.ErrNoRows) {
		return false, nil
	}

	return err == nil, err
}

// Close closes all connections in the pool and rejects future Acquire calls. Blocks until all connections are returned
// to pool and closed.
func (db *DB) Close() {
	db.master.Close()
}

func (db *DB) Beginx(ctx context.Context) (*Tx, error) {
	tx, err := db.master.Begin(ctx)
	if err != nil {
		return nil, err
	}
	return &Tx{tx}, nil
}

// ExecuteInTx runs fn inside tx which should already have begun.
func (db *DB) ExecuteInTx(ctx context.Context, txConsistency *Tx, fn func(tx *Tx) error) (err error) {
	if txConsistency != nil {
		err = fn(txConsistency)
		return
	}

	tx, err := db.Beginx(ctx)
	if err != nil {
		return err
	}

	err = fn(tx)
	if err == nil {
		err = tx.Commit(ctx)
	} else {
		_ = tx.Rollback(ctx)
	}
	return
}
