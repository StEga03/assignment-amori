package message

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

type messageTable struct {
	ID          uint64    `db:"id" table:"messages"`
	ChannelID   uint64    `db:"channel_id"`
	SenderType  string    `db:"sender_type"`
	SenderID    uint64    `db:"sender_id"`
	ContentType string    `db:"content_type"`
	Content     string    `db:"content"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}

type messageInputTable struct {
	ID              uint64    `db:"id" table:"message_inputs"`
	ChannelID       uint64    `db:"channel_id"`
	Source          string    `db:"source"`
	Sender          string    `db:"sender"`
	Receiver        string    `db:"receiver"`
	ReceiverPronoun string    `db:"receiver_pronoun"`
	CreatedAt       time.Time `db:"created_at"`
	UpdatedAt       time.Time `db:"updated_at"`
}

type messageSourceTable struct {
	ID             uint64    `db:"id" table:"messages_sources"`
	MessageInputID uint64    `db:"message_input_id"`
	Sender         string    `db:"sender"`
	ContentType    string    `db:"content_type"`
	Content        string    `db:"content"`
	SentAt         time.Time `db:"sent_at"`
	CreatedAt      time.Time `db:"created_at"`
	UpdatedAt      time.Time `db:"updated_at"`
}
