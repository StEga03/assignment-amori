package http

import (
	"net/http"

	"github.com/assignment-amori/internal/entity"
	httphelper "github.com/assignment-amori/internal/handler/helper/http"
)

func (h *Handler) CreateChannel(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	ctx := r.Context()

	req := entity.NewChannelUCRequest{}
	if err := httphelper.CastAndValidate(r, &req); err != nil {
		return nil, err
	}

	return h.ChannelUC.CreateChannel(ctx, req)
}
