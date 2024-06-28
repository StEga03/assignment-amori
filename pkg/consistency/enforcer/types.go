package enforcer

import (
	"context"

	"github.com/assignment-amori/pkg/sql/pgx"
)

type Module struct {
	db dbResource
}

type dbResource interface {
	Beginx(ctx context.Context) (*pgx.Tx, error)
}

type txResource interface {
	Commit(ctx context.Context) error
	Rollback(ctx context.Context) error
}
