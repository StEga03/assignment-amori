package channel

import (
	"context"
	"time"

	"github.com/assignment-amori/pkg/sql/pgx"
	pgxmodule "github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type databaseResource interface {
	Insert(ctx context.Context, entity interface{}) (pgconn.CommandTag, error)
	Update(ctx context.Context, entity interface{}, options *pgx.UpdateOptions) (pgconn.CommandTag, error)
	Select(ctx context.Context, dest interface{}, sql string, args ...interface{}) (bool, error)
	ExecuteInTx(ctx context.Context, txConsistency *pgx.Tx, fn func(tx *pgx.Tx) error) (err error)
	QueryRow(ctx context.Context, sql string, args ...any) pgxmodule.Row
	Query(ctx context.Context, sql string, args ...any) (pgxmodule.Rows, error)
}

type txResource interface {
	Insert(ctx context.Context, entity interface{}) (pgconn.CommandTag, error)
	Update(ctx context.Context, entity interface{}, options *pgx.UpdateOptions) (pgconn.CommandTag, error)
	QueryRow(ctx context.Context, sql string, args ...any) pgxmodule.Row
	Query(ctx context.Context, sql string, args ...any) (pgxmodule.Rows, error)
	Beginx(ctx context.Context) (*pgx.Tx, error)
}

type sonyFlakeResource interface {
	NextID() (uint64, error)
}

type Repository struct {
	db databaseResource
	sf sonyFlakeResource
}

type channelTable struct {
	ID        uint64    `db:"id" table:"channels"`
	UserID    uint64    `db:"user_id"`
	Name      string    `db:"name"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
