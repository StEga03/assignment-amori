package message

import (
	"context"

	"github.com/assignment-amori/internal/entity"
	"github.com/assignment-amori/pkg/consistency"
	"github.com/assignment-amori/pkg/errorwrapper"
	"github.com/assignment-amori/pkg/sql/pgx"
	timeutils "github.com/assignment-amori/pkg/time_utils"
)

func (r *Repository) CreateMessageInput(ctx context.Context, param entity.NewMessageInputParams, cel *consistency.ConsistencyElement) (uint64, error) {
	var (
		tx  *pgx.Tx
		err error

		now = timeutils.Now()
	)

	if cel != nil {
		tx = cel.Txn
	}

	err = r.db.ExecuteInTx(ctx, tx, func(tx *pgx.Tx) error {
		// Insert Message Input.
		messageInput := messageInputTable{
			ID:              param.ID,
			ChannelID:       param.ChannelID,
			Source:          param.Source,
			Sender:          param.Sender,
			Receiver:        param.Receiver,
			ReceiverPronoun: param.ReceiverPronoun,
			CreatedAt:       now,
			UpdatedAt:       now,
		}

		err := r.insertMessageInput(ctx, tx, messageInput)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return param.ID, errorwrapper.Wrap(err, errorwrapper.ErrIDFailedCreateFromRepoMessageInput)
	}

	return param.ID, nil
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
