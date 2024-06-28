package http

import (
	"github.com/assignment-amori/internal/constant"
	"github.com/go-chi/chi/v5"
)

func (u *Handler) Routes(hm helperModule) chi.Router {
	r := chi.NewRouter()
	r.Use(hm.APIModule(constant.ModuleUploads))

	r.Post("/messages/sources/{sourceType}", hm.HandleHTTP(
		constant.HTTPHandlerIDUploadMessageSource,
		constant.HTTPDefaultResponseWriter,
		u.UploadMessageSource,
	))

	return r
}
