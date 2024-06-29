package user

import (
	"context"

	"github.com/assignment-amori/internal/entity"
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

func (u *Usecase) GetCurrentUser(ctx context.Context) (entity.User, error) {
	var (
		result entity.User
		err    error
	)

	result, err = u.userRepo.GetUserByContext(ctx)
	if err != nil {
		return result, err
	}

	return result, nil
}
