package pgx

import (
	"context"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/wcamarao/pmx"
)

func (tx *Tx) Insert(ctx context.Context, entity interface{}) (pgconn.CommandTag, error) {
	return pmx.Insert(ctx, tx, entity)
}

func (tx *Tx) Update(ctx context.Context, entity interface{}, options *UpdateOptions) (pgconn.CommandTag, error) {
	return pmx.Update(ctx, tx, entity, (*pmx.UpdateOptions)(options))
}

func (t *Tx) Beginx(ctx context.Context) (*Tx, error) {
	tx, err := t.Begin(ctx)
	if err != nil {
		return nil, err
	}

	return &Tx{tx}, nil
}
