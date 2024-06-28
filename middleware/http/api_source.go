package http

import (
	"context"
	"net/http"

	"github.com/assignment-amori/internal/constant"
)

func (m *MiddlewareModule) APISource(source string) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			r = r.WithContext(context.WithValue(r.Context(), constant.ContextKeyAPISource, source))
			next.ServeHTTP(w, r)
		})
	}
}
