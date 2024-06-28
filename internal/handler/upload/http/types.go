package http

import (
	"net/http"

	"github.com/assignment-amori/internal/constant"
	"github.com/assignment-amori/internal/entity/generic"
	mwHTTP "github.com/assignment-amori/middleware/http"
)

type helperModule interface {
	HandleHTTP(handlerName constant.HandlerID, typ constant.MiddlewareID, h generic.HTTPHandleFunc, mw ...mwHTTP.MiddlewareExecutor) generic.HTTPHandler
	APIModule(module constant.Module) func(next http.Handler) http.Handler
}

type Handler struct{}
