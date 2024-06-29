package http

import (
	"context"
	"net/http"

	"github.com/assignment-amori/internal/constant"
	"github.com/assignment-amori/internal/entity/generic"
	mwHTTP "github.com/assignment-amori/middleware/http"
	pkgHttp "github.com/assignment-amori/pkg/http"
	"github.com/assignment-amori/pkg/whatsapp"
)

//go:generate mockgen -package=http -source=types.go -destination=file_http_mock_test.go
type FileUC interface {
	WhatsappParser(ctx context.Context, fileDetails pkgHttp.FileDetails) ([]whatsapp.Message, error)
}

type helperModule interface {
	HandleHTTP(handlerName constant.HandlerID, typ constant.MiddlewareID, h generic.HTTPHandleFunc, mw ...mwHTTP.MiddlewareExecutor) generic.HTTPHandler
	APIModule(module constant.Module) func(next http.Handler) http.Handler
	GetJWTAuthMiddleware() *mwHTTP.JWTMiddleware
}

type Handler struct {
	fileUC FileUC
}
