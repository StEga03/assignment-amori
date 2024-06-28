package user

import (
	"context"

	"github.com/assignment-amori/internal/entity"
)

type userResource interface {
	GetByID(ctx context.Context, id uint64) (entity.User, error)
	
	JWTGenerator(ctx context.Context, user entity.User) (string, error)
}

type Usecase struct {
	userRepo userResource
}
