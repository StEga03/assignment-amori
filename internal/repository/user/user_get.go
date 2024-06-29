package user

import (
	"context"

	"github.com/assignment-amori/internal/constant"
	"github.com/assignment-amori/internal/entity"
	"github.com/assignment-amori/pkg/errorwrapper"
	"github.com/golang-jwt/jwt/v4"
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

func (r *Repository) GetUserByContext(ctx context.Context) (entity.User, error) {
	var result entity.User

	token := ctx.Value(constant.ContextUser)

	// Parsing data from claims and get the column as map.
	userIdRaw, ok := token.(jwt.MapClaims)[constant.UserColumnID]
	if !ok {
		return result, errorwrapper.New(errorwrapper.ErrParsing)
	}

	userIdFloat, ok := userIdRaw.(float64)
	if !ok {
		return result, errorwrapper.New(errorwrapper.ErrParsing)
	}

	return r.GetByID(ctx, uint64(userIdFloat))
}
