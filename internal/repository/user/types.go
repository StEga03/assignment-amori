package user

import (
	"context"
	"database/sql"
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

type userTable struct {
	ID                 uint64         `db:"id" table:"users"`
	FirstName          string         `db:"first_name"`
	LastName           sql.NullString `db:"last_name"`
	BirthDate          sql.NullTime   `db:"birth_date"`
	Gender             sql.NullString `db:"gender"`
	GenderInterest     sql.NullString `db:"gender_interest"`
	PhoneNumber        sql.NullString `db:"phone_number"`
	RelationshipStatus sql.NullString `db:"relationship_status"`
	RelationshipGoal   sql.NullString `db:"relationship_goal"`
	CreatedAt          time.Time      `db:"created_at"`
	UpdatedAt          time.Time      `db:"updated_at"`
}
