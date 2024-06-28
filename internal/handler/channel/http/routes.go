package http

import (
	"github.com/assignment-amori/internal/constant"
	"github.com/go-chi/chi/v5"
)

func (u *Handler) Routes(hm helperModule) chi.Router {
	r := chi.NewRouter()
	r.Use(hm.APIModule(constant.ModuleChannels))

	r.Post("/", hm.HandleHTTP(
		constant.HTTPHandlerIDChannelCreate,
		constant.HTTPDefaultResponseWriter,
		u.CreateChannel,
	))

	r.Post("/{channelId}/messages", hm.HandleHTTP(
		constant.HTTPHandlerIDChannelMessageCreate,
		constant.HTTPDefaultResponseWriter,
		u.CreateMessageInChannel,
	))

	r.Get("/{channelId}/messages", hm.HandleHTTP(
		constant.HTTPHandlerIDChannelMessageGet,
		constant.HTTPDefaultResponseWriter,
		u.GetMessageInChannelByID,
	))

	return r
}
