package http

import (
	"net/http"
	"strconv"

	"github.com/assignment-amori/internal/constant"
	"github.com/assignment-amori/internal/entity"
	httphelper "github.com/assignment-amori/internal/handler/helper/http"
	"github.com/assignment-amori/pkg/errorwrapper"
	"github.com/assignment-amori/pkg/validator"
	"github.com/go-chi/chi/v5"
)

func (h *Handler) CreateMessageInChannel(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	ctx := r.Context()

	channelIdRaw := chi.URLParam(r, constant.ParamChannelID)
	if err := validator.ValidateVar(ctx, channelIdRaw, "required"); err != nil {
		return nil, errorwrapper.Wrap(err, errorwrapper.ErrIDValidationNotPassed)
	}

	channelId, err := strconv.ParseUint(channelIdRaw, 10, 64)
	if err != nil {
		return nil, errorwrapper.Wrap(err, errorwrapper.ErrParsing)
	}

	req := entity.MessageUCRequest{}
	req.ChannelID = channelId
	if err := httphelper.CastAndValidate(r, &req); err != nil {
		return nil, err
	}

	return h.ChannelUC.CreateMessageInChannel(ctx, req)
}
