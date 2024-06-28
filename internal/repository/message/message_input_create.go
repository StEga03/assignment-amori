package message

import (
	"context"

	"github.com/assignment-amori/internal/entity"
	"github.com/assignment-amori/pkg/consistency"
	"github.com/assignment-amori/pkg/errorwrapper"
	"github.com/assignment-amori/pkg/sql/pgx"
)

func (r *Repository) CreateMessageInput(ctx context.Context, param entity.NewMessageInputParams, cel *consistency.ConsistencyElement) (uint64, error) {
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
		// Insert Message Input.
		messageInput := messageInputTable{
			ID:              messageInputId,
			ChannelID:       param.ChannelID,
			Source:          param.Source,
			Sender:          param.Sender,
			Receiver:        param.Receiver,
			ReceiverPronoun: param.ReceiverPronoun,
			CreatedAt:       param.CreatedAt,
			UpdatedAt:       param.CreatedAt,
		}

		err := r.insertMessageInput(ctx, tx, messageInput)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return messageInputId, errorwrapper.Wrap(err, errorwrapper.ErrIDFailedCreateFromRepoMessageInput)
	}

	return messageInputId, nil
}

/*
================================================ UNEXPORTED FUNCTIONS =====================================================
*/

func (r *Repository) insertMessageInput(ctx context.Context, tx txResource, msgInput messageInputTable) (err error) {
	_, err = tx.Insert(ctx, &msgInput)

	return errorwrapper.Wrap(err, errorwrapper.ErrIDInsertDB,
		errorwrapper.WithMetaData(
			errorwrapper.MetaKV{
				"stmt": insertMessageInput,
			}))
}
