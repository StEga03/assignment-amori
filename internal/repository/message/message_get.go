package message

import (
	"context"

	"github.com/assignment-amori/internal/constant"
	"github.com/assignment-amori/internal/entity"
	"github.com/assignment-amori/pkg/errorwrapper"
)

func (r *Repository) GetMessageByID(ctx context.Context, param entity.GetMessageParams) (entity.Message, error) {
	var (
		result entity.Message
		err    error
	)

	resultTable := messageTable{}
	_, err = r.db.Select(ctx, &resultTable, querySelectMessageById, param.ID)
	if err != nil {
		return result, errorwrapper.Wrap(err, errorwrapper.ErrIDFailedGetFromRepoMessage)
	}

	result = resultTable.ToEntity()
	if result.ID == constant.DefaultUInt64 {
		return result, errorwrapper.New(errorwrapper.ErrIDMessageDataIsEmpty)
	}

	return result, nil
}

func (r *Repository) GetMessageByChannelID(ctx context.Context, param entity.GetMessageParams) ([]entity.Message, error) {
	var (
		result          []entity.Message
		resMessageTable []*messageTable
		err             error
	)

	_, err = r.db.Select(ctx, &resMessageTable, querySelectMessageByChannelId, param.ChannelID, param.Limit, param.Offset)
	if err != nil {
		return result, errorwrapper.Wrap(err, errorwrapper.ErrIDFailedGetFromRepoMessage)
	}

	for _, table := range resMessageTable {
		result = append(result, table.ToEntity())
	}

	return result, nil
}
