package message

import (
	"context"

	"github.com/assignment-amori/internal/entity"
	"github.com/assignment-amori/pkg/consistency"
	"github.com/assignment-amori/pkg/errorwrapper"
	"github.com/assignment-amori/pkg/sql/pgx"
)

func (r *Repository) CreateMessageSource(ctx context.Context, param entity.NewMessageSourceParams, cel *consistency.ConsistencyElement) (uint64, error) {
	var (
		messageInputId uint64
		tx             *pgx.Tx
		err            error
	)

	if cel != nil {
		tx = cel.Txn
	}

	messageInputId, err = r.sf.NextID()
	if err != nil {
		return messageInputId, errorwrapper.Wrap(err, errorwrapper.ErrIDFailedToGenerateID)
	}

	err = r.db.ExecuteInTx(ctx, tx, func(tx *pgx.Tx) error {
		// Insert Message Source.
		messageSource := messageSourceTable{
			ID:             messageInputId,
			MessageInputID: param.MessageInputID,
			Sender:         param.Sender,
			ContentType:    param.ContentType,
			Content:        param.Content,
			SentAt:         param.SentAt,
			CreatedAt:      param.CreatedAt,
			UpdatedAt:      param.CreatedAt,
		}

		err := r.insertMessageSource(ctx, tx, messageSource)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return messageInputId, errorwrapper.Wrap(err, errorwrapper.ErrIDFailedCreateFromRepoMessageSource)
	}

	return messageInputId, nil
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
