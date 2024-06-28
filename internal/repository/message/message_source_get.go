package message

import (
	"context"

	"github.com/assignment-amori/internal/constant"
	"github.com/assignment-amori/internal/entity"
	"github.com/assignment-amori/pkg/errorwrapper"
)

func (r *Repository) GetMessageSourceByID(ctx context.Context, param entity.GetMessageSourceParams) (entity.MessageSource, error) {
	var (
		result entity.MessageSource
		err    error
	)

	resultTable := messageSourceTable{}
	_, err = r.db.Select(ctx, &resultTable, querySelectMessageSourceById, param.ID)
	if err != nil {
		return result, errorwrapper.Wrap(err, errorwrapper.ErrIDFailedGetFromRepoMessageSource)
	}

	result = resultTable.ToEntity()
	if result.ID == constant.DefaultUInt64 {
		return result, errorwrapper.New(errorwrapper.ErrIDMessageSourceDataIsEmpty)
	}

	return result, nil
}

func (r *Repository) GetMessageSourceByMessageInputID(ctx context.Context, param entity.GetMessageSourceParams) ([]entity.MessageSource, error) {
	var (
		result          []entity.MessageSource
		resMessageTable []*messageSourceTable
		err             error
	)

	_, err = r.db.Select(ctx, &resMessageTable, querySelectMessageSourceByMessageInputId, param.MessageInputID, param.Limit, param.Offset)
	if err != nil {
		return result, errorwrapper.Wrap(err, errorwrapper.ErrIDFailedGetFromRepoMessageSource)
	}

	for _, table := range resMessageTable {
		result = append(result, table.ToEntity())
	}

	return result, nil
}
