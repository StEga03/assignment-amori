package user

import (
	"context"
	"strconv"

	"github.com/assignment-amori/internal/constant"
	"github.com/assignment-amori/internal/entity"
	"github.com/assignment-amori/pkg/errorwrapper"
	"github.com/golang-jwt/jwt/v4"
)

func (u *Usecase) GenerateValidJWT(ctx context.Context, userId uint64) (string, error) {
	var (
		result string
		err    error
	)

	user, err := u.userRepo.GetByID(ctx, userId)
	if err != nil {
		return result, err
	}

	result, err = u.userRepo.JWTGenerator(ctx, user)
	if err != nil {
		return result, err
	}

	return result, nil
}

func (u *Usecase) GetCurrentUser(ctx context.Context, token any) (entity.User, error) {
	var (
		result entity.User
		err    error
	)

	// Parsing data from claims and get the column as map.
	userIdRaw, ok := token.(jwt.MapClaims)[constant.UserColumnID]
	if !ok {
		return result, errorwrapper.New(errorwrapper.ErrParsing)
	}

	userIdStr, ok := userIdRaw.(string)
	if !ok {
		return result, errorwrapper.New(errorwrapper.ErrParsing)
	}

	userId, err := strconv.ParseUint(userIdStr, 10, 64)
	if err != nil {
		return result, errorwrapper.Wrap(err, errorwrapper.ErrParsing)
	}

	result, err = u.userRepo.GetByID(ctx, userId)
	if err != nil {
		return result, err
	}

	return result, nil
}
