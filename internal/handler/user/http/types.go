package http

import (
	"context"
	"net/http"

	"github.com/assignment-amori/internal/constant"
	"github.com/assignment-amori/internal/entity"
	"github.com/assignment-amori/internal/entity/generic"
	mwHTTP "github.com/assignment-amori/middleware/http"
)

type UserUsecase interface {
	GenerateValidJWT(ctx context.Context, userId uint64) (string, error)
	GetCurrentUser(ctx context.Context) (entity.User, error)
}

type helperModule interface {
	HandleHTTP(handlerName constant.HandlerID, typ constant.MiddlewareID, h generic.HTTPHandleFunc, mw ...mwHTTP.MiddlewareExecutor) generic.HTTPHandler
	APIModule(module constant.Module) func(next http.Handler) http.Handler
	GetJWTAuthMiddleware() *mwHTTP.JWTMiddleware
}

type Handler struct {
	UserUC UserUsecase
}
