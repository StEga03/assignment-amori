package http

import (
	"github.com/assignment-amori/internal/constant"
	"github.com/assignment-amori/internal/entity/generic"
)

type MiddlewareExecutor interface {
	Filter(generic.HTTPHandleFunc) generic.HTTPHandleFunc
}

func setupHTTPMiddlewarePacket(hm *MiddlewareModule, handlerName constant.HandlerID, m ...MiddlewareExecutor) []MiddlewareExecutor {
	m = append(m, hm.GetResponseWriterMiddleware())
	return m
}

// Middleware plain for return response data without any interfere.
// Custom response type on the handler side.
func setupHTTPMiddlewarePlainPacket(hm *MiddlewareModule, handlerName constant.HandlerID, m ...MiddlewareExecutor) []MiddlewareExecutor {
	m = append(m, hm.GetPlainResponseWriterMiddleware())
	return m
}
