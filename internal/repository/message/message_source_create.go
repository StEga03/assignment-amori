package message

import (
	"context"

	"github.com/assignment-amori/internal/entity"
	"github.com/assignment-amori/pkg/consistency"
	"github.com/assignment-amori/pkg/errorwrapper"
	"github.com/assignment-amori/pkg/sql/pgx"
	timeutils "github.com/assignment-amori/pkg/time_utils"
)

func (r *Repository) CreateMessageSource(ctx context.Context, param []entity.NewMessageSourceParams, cel *consistency.ConsistencyElement) error {
	var (
		tx  *pgx.Tx
		err error

		now = timeutils.Now()
	)

	if cel != nil {
		tx = cel.Txn
	}

	err = r.db.ExecuteInTx(ctx, tx, func(tx *pgx.Tx) error {
		// Insert Message Source.
		for _, req := range param {
			id, err := r.sf.NextID()
			if err != nil {
				return errorwrapper.Wrap(err, errorwrapper.ErrIDFailedToGenerateID)
			}

			messageSource := messageSourceTable{
				ID:             id,
				MessageInputID: req.MessageInputID,
				Sender:         req.Sender,
				ContentType:    req.ContentType,
				Content:        req.Content,
				SentAt:         req.SentAt,
				CreatedAt:      now,
				UpdatedAt:      now,
			}

			err = r.insertMessageSource(ctx, tx, messageSource)
			if err != nil {
				return err
			}
		}

		return nil
	})
	if err != nil {
		return errorwrapper.Wrap(err, errorwrapper.ErrIDFailedCreateFromRepoMessageSource)
	}

	return nil
}

/*
================================================ UNEXPORTED FUNCTIONS =====================================================
*/

func (r *Repository) insertMessageSource(ctx context.Context, tx txResource, msgSource messageSourceTable) (err error) {
	_, err = tx.Insert(ctx, &msgSource)

	return errorwrapper.Wrap(err, errorwrapper.ErrIDInsertDB,
		errorwrapper.WithMetaData(
			errorwrapper.MetaKV{
				"stmt": insertMessageSource,
			}))
}
