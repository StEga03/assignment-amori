package http

import (
	"context"
	"net/http"

	"github.com/assignment-amori/internal/constant"
	"github.com/assignment-amori/internal/entity"
	"github.com/assignment-amori/internal/entity/generic"
	mwHTTP "github.com/assignment-amori/middleware/http"
)

//go:generate mockgen -package=http -source=types.go -destination=channel_http_mock_test.go
type ChannelUsecase interface {
	CreateChannel(ctx context.Context, req entity.NewChannelUCRequest) (uint64, error)
	CreateMessageInChannel(ctx context.Context, req entity.MessageUCRequest) (entity.MessageResponse, error)
	GetMessageInChannel(ctx context.Context, req entity.MessageUCRequest) ([]entity.MessageResponse, error)
}

type helperModule interface {
	HandleHTTP(handlerName constant.HandlerID, typ constant.MiddlewareID, h generic.HTTPHandleFunc, mw ...mwHTTP.MiddlewareExecutor) generic.HTTPHandler
	APIModule(module constant.Module) func(next http.Handler) http.Handler
}

type Handler struct {
	ChannelUC ChannelUsecase
}
