package channel

import (
	"context"

	"github.com/assignment-amori/internal/constant"
	"github.com/assignment-amori/internal/entity"
	"github.com/assignment-amori/pkg/errorwrapper"
)

func (r *Repository) GetByID(ctx context.Context, id uint64) (entity.Channel, error) {
	var (
		result entity.Channel
		err    error
	)

	resultTable := channelTable{}
	_, err = r.db.Select(ctx, &resultTable, querySelectChannelByID, id)
	if err != nil {
		return result, errorwrapper.Wrap(err, errorwrapper.ErrIDFailedGetFromRepoChannel)
	}

	result = resultTable.ToEntity()
	if result.ID == constant.DefaultUInt64 {
		return result, errorwrapper.New(errorwrapper.ErrIDChannelDataIsEmpty)
	}

	return result, nil
}

func (r *Repository) GetByUserID(ctx context.Context, userId string) ([]entity.Channel, error) {
	var (
		result          []entity.Channel
		resChannelTable []*channelTable
		err             error
	)

	_, err = r.db.Select(ctx, &resChannelTable, querySelectChannelByUserID, userId)
	if err != nil {
		return result, errorwrapper.Wrap(err, errorwrapper.ErrIDFailedGetFromRepoChannel)
	}

	for _, table := range resChannelTable {
		result = append(result, table.ToEntity())
	}

	return result, nil
}
