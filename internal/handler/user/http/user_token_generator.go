package http

import (
	"net/http"
	"strconv"

	"github.com/assignment-amori/internal/constant"
	"github.com/assignment-amori/pkg/errorwrapper"
	"github.com/assignment-amori/pkg/validator"
	"github.com/go-chi/chi/v5"
)

func (h *Handler) GenerateJWT(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	ctx := r.Context()

	userIdRaw := chi.URLParam(r, constant.ParamUserID)
	if err := validator.ValidateVar(ctx, userIdRaw, "required"); err != nil {
		return nil, errorwrapper.Wrap(err, errorwrapper.ErrIDValidationNotPassed)
	}

	userId, err := strconv.ParseUint(userIdRaw, 10, 64)
	if err != nil {
		return nil, errorwrapper.Wrap(err, errorwrapper.ErrParsing)
	}

	return h.UserUC.GenerateValidJWT(ctx, userId)
}
