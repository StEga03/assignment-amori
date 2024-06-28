package message

import (
	"context"

	"github.com/assignment-amori/internal/constant"
	"github.com/assignment-amori/internal/entity"
	"github.com/assignment-amori/pkg/errorwrapper"
)

func (r *Repository) GetMessageInputByID(ctx context.Context, param entity.GetMessageInputParams) (entity.MessageInput, error) {
	var (
		result entity.MessageInput
		err    error
	)

	resultTable := messageInputTable{}
	_, err = r.db.Select(ctx, &resultTable, querySelectMessageInputById, param.ID)
	if err != nil {
		return result, errorwrapper.Wrap(err, errorwrapper.ErrIDFailedGetFromRepoMessageInput)
	}

	result = resultTable.ToEntity()
	if result.ID == constant.DefaultUInt64 {
		return result, errorwrapper.New(errorwrapper.ErrIDMessageInputDataIsEmpty)
	}

	return result, nil
}

func (r *Repository) GetMessageInputByChannelID(ctx context.Context, param entity.GetMessageParams) ([]entity.MessageInput, error) {
	var (
		result          []entity.MessageInput
		resMessageTable []*messageInputTable
		err             error
	)

	_, err = r.db.Select(ctx, &resMessageTable, querySelectMessageInputByChannelId, param.ChannelID, param.Limit, param.Offset)
	if err != nil {
		return result, errorwrapper.Wrap(err, errorwrapper.ErrIDFailedGetFromRepoMessageInput)
	}

	for _, table := range resMessageTable {
		result = append(result, table.ToEntity())
	}

	return result, nil
}
