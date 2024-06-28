package message

import (
	"context"

	"github.com/assignment-amori/internal/entity"
	"github.com/assignment-amori/pkg/consistency"
	"github.com/assignment-amori/pkg/errorwrapper"
	"github.com/assignment-amori/pkg/sql/pgx"
)

func (r *Repository) CreateMessage(ctx context.Context, param entity.NewMessageParams, cel *consistency.ConsistencyElement) (uint64, error) {
	var (
		messageId uint64
		tx        *pgx.Tx
		err       error
	)

	if cel != nil {
		tx = cel.Txn
	}

	messageId, err = r.sf.NextID()
	if err != nil {
		return messageId, errorwrapper.Wrap(err, errorwrapper.ErrIDFailedToGenerateID)
	}

	err = r.db.ExecuteInTx(ctx, tx, func(tx *pgx.Tx) error {
		// Insert Message.
		message := messageTable{
			ID:          messageId,
			ChannelID:   param.ChannelID,
			SenderType:  param.SenderType,
			SenderID:    param.SenderID,
			ContentType: param.ContentType,
			Content:     param.Content,
			CreatedAt:   param.CreatedAt,
			UpdatedAt:   param.CreatedAt,
		}

		err := r.insertMessage(ctx, tx, message)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return messageId, errorwrapper.Wrap(err, errorwrapper.ErrIDFailedCreateFromRepoMessage)
	}

	return messageId, nil
}

func (r *Repository) CreateMessageBulk(ctx context.Context, param []entity.NewMessageParams, cel *consistency.ConsistencyElement) error {
	var (
		tx  *pgx.Tx
		err error
	)

	if cel != nil {
		tx = cel.Txn
	}

	err = r.db.ExecuteInTx(ctx, tx, func(tx *pgx.Tx) error {
		// Insert Message.
		for _, msg := range param {
			messageId, err := r.sf.NextID()
			if err != nil {
				return errorwrapper.Wrap(err, errorwrapper.ErrIDFailedToGenerateID)
			}

			message := messageTable{
				ID:          messageId,
				ChannelID:   msg.ChannelID,
				SenderType:  msg.SenderType,
				SenderID:    msg.SenderID,
				ContentType: msg.ContentType,
				Content:     msg.Content,
				CreatedAt:   msg.CreatedAt,
				UpdatedAt:   msg.CreatedAt,
			}

			err = r.insertMessage(ctx, tx, message)
			if err != nil {
				return err
			}
		}

		return nil
	})
	if err != nil {
		return errorwrapper.Wrap(err, errorwrapper.ErrIDFailedCreateFromRepoMessage)
	}

	return nil
}

/*
================================================ UNEXPORTED FUNCTIONS =====================================================
*/

func (r *Repository) insertMessage(ctx context.Context, tx txResource, msg messageTable) (err error) {
	_, err = tx.Insert(ctx, &msg)

	return errorwrapper.Wrap(err, errorwrapper.ErrIDInsertDB,
		errorwrapper.WithMetaData(
			errorwrapper.MetaKV{
				"stmt": insertMessage,
			}))
}
