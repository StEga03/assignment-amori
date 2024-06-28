package http

import (
	"context"
	"net/http"

	"github.com/assignment-amori/internal/constant"
	"github.com/assignment-amori/internal/entity/generic"
	"github.com/google/uuid"
)

func (hm *MiddlewareModule) HandleHTTP(
	handlerName constant.HandlerID,
	typ constant.MiddlewareID,
	h generic.HTTPHandleFunc,
	mw ...MiddlewareExecutor,
) generic.HTTPHandler {
	switch typ {
	case constant.HTTPDefaultResponseWriter:
		mw = setupHTTPMiddlewarePacket(hm, handlerName, mw...)
	case constant.HTTPPlainResponseWriter:
		mw = setupHTTPMiddlewarePlainPacket(hm, handlerName, mw...)
	}

	for _, m := range mw {
		h = m.Filter(h)
	}

	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		// Get request id from header first.
		requestID := r.Header.Get(constant.HTTPHeaderRequestID)

		// Create new request id if it not exists from header.
		if requestID == constant.DefaultString {
			requestID = uuid.New().String()
		}

		// Assign Language.
		ctx = context.WithValue(ctx, constant.ContextKeyLanguage, r.Header.Get(constant.HTTPHeaderAcceptLanguage))
		r = r.WithContext(ctx)

		w.Header().Set(constant.HTTPHeaderRequestID, requestID)
		h(w, r)
	}
}
