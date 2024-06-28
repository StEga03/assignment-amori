package http

import (
	"net/http"

	"github.com/assignment-amori/pkg/errorwrapper"
	"github.com/assignment-amori/pkg/validator"
	"github.com/go-chi/render"
)

func CastAndValidate(r *http.Request, target interface{}) error {
	err := render.DefaultDecoder(r, target)
	if err != nil {
		return errorwrapper.Wrap(err, errorwrapper.ErrIDUnmarshal, errorwrapper.WithUserMsgParams(
			errorwrapper.UserMessageParams{
				errorwrapper.DataName: "request",
			},
		))
	}

	return errorwrapper.Wrap(
		validator.ValidateStruct(r.Context(), target),
		errorwrapper.ErrIDValidationNotPassed)
}
