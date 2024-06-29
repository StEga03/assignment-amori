package http

import (
	"net/http"
)

func (h *Handler) GetCurrentUser(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	ctx := r.Context()

	return h.UserUC.GetCurrentUser(ctx)
}
