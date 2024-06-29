package http

import (
	"github.com/assignment-amori/internal/constant"
	"github.com/go-chi/chi/v5"
)

func (u *Handler) Routes(hm helperModule) chi.Router {
	r := chi.NewRouter()
	r.Use(hm.APIModule(constant.ModuleUsers))

	r.Get("/{userId}/token/generator", hm.HandleHTTP(
		constant.HTTPHandlerIDUserTokenGeneratorGet,
		constant.HTTPDefaultResponseWriter,
		u.GenerateJWT,
	))

	r.Get("/current", hm.HandleHTTP(
		constant.HTTPHandlerIDUserCurrentGet,
		constant.HTTPDefaultResponseWriter,
		u.GetCurrentUser,
		hm.GetJWTAuthMiddleware(),
	))

	return r
}
