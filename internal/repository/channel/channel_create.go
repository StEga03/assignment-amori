package channel

import (
	"context"

	"github.com/assignment-amori/internal/entity"
	"github.com/assignment-amori/pkg/consistency"
	"github.com/assignment-amori/pkg/errorwrapper"
	"github.com/assignment-amori/pkg/sql/pgx"
	timeutils "github.com/assignment-amori/pkg/time_utils"
)

func (r *Repository) CreateChannel(ctx context.Context, param entity.NewChannelParams, cel *consistency.ConsistencyElement) (uint64, error) {
	var (
		tx  *pgx.Tx
		err error

		now = timeutils.Now()
	)

	if cel != nil {
		tx = cel.Txn
	}

	err = r.db.ExecuteInTx(ctx, tx, func(tx *pgx.Tx) error {
		// Insert Channel.
		channel := channelTable{
			ID:        param.ID,
			UserID:    param.UserID,
			Name:      param.Name,
			CreatedAt: now,
			UpdatedAt: now,
		}

		err := r.insertChannel(ctx, tx, channel)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return param.ID, errorwrapper.Wrap(err, errorwrapper.ErrIDFailedCreateFromRepoChannel)
	}

	return param.ID, nil
}

/*
================================================ UNEXPORTED FUNCTIONS =====================================================
*/

func (r *Repository) insertChannel(ctx context.Context, tx txResource, channel channelTable) (err error) {
	_, err = tx.Insert(ctx, &channel)

	return errorwrapper.Wrap(err, errorwrapper.ErrIDInsertDB,
		errorwrapper.WithMetaData(
			errorwrapper.MetaKV{
				"stmt": insertChannel,
			}))
}
