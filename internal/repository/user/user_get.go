package user

import (
	"context"

	"github.com/assignment-amori/internal/constant"
	"github.com/assignment-amori/internal/entity"
	"github.com/assignment-amori/pkg/errorwrapper"
)

func (r *Repository) GetByID(ctx context.Context, id uint64) (entity.User, error) {
	var (
		result entity.User
		err    error
	)

	resultTable := userTable{}
	_, err = r.db.Select(ctx, &resultTable, querySelectUserById, id)
	if err != nil {
		return result, errorwrapper.Wrap(err, errorwrapper.ErrIDFailedGetFromRepoUser)
	}

	result = resultTable.ToEntity()
	if result.ID == constant.DefaultUInt64 {
		return result, errorwrapper.New(errorwrapper.ErrIDUserDataIsEmpty)
	}

	return result, nil
}
