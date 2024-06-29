package http

import (
	"github.com/assignment-amori/internal/constant"
	"github.com/go-chi/chi/v5"
)

func (u *Handler) Routes(hm helperModule) chi.Router {
	r := chi.NewRouter()
	r.Use(hm.APIModule(constant.ModuleFiles))

	r.Post("/sources/{sourceType}", hm.HandleHTTP(
		constant.HTTPHandlerIDFilesMessageSource,
		constant.HTTPDefaultResponseWriter,
		u.UploadMessageSource,
		hm.GetJWTAuthMiddleware(),
	))

	return r
}
