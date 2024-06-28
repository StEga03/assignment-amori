package http

import (
	"net/http"

	"github.com/assignment-amori/internal/constant"
	"github.com/assignment-amori/pkg/errorwrapper"
	"github.com/assignment-amori/pkg/locale"
	pkgvalidator "github.com/assignment-amori/pkg/validator"
	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
)

// WriteError write error http with errorWrapper as a standard error response.
func WriteError(w http.ResponseWriter, r *http.Request, err error, data interface{}) {
	// Cast error to error wrapper.
	errWrapper := errorwrapper.CastToErrorWrapper(err)
	if errWrapper == nil {
		return
	}

	// Build response.
	response := &GlobalResponse{
		Code:      errWrapper.GetErrorCode(),
		Message:   errWrapper.UserMessage(r.Context()),
		Retryable: errWrapper.GetIsRetryable(),
	}
	vErrs := errWrapper.GetValidatorErrs()
	if vErrs != nil {
		response.Details = append(response.Details, appendValidatorErrDetail(r, vErrs)...)
	}

	// Assign data if exist.
	if data != nil {
		response.Data = data
	}

	// Assign header Request ID.
	w.Header().Set(constant.HTTPHeaderRequestID, constant.DefaultString)

	render.Status(r, errWrapper.GetHTTPStatus())
	render.DefaultResponder(w, r, response)
}

func appendValidatorErrDetail(r *http.Request, vErrs validator.ValidationErrors) []ErrorDetail {
	var details []ErrorDetail
	for _, vErr := range vErrs {
		details = append(details, ErrorDetail{
			Path: vErr.Field(),
			Info: pkgvalidator.BuildErrorMessage(r.Context(), vErr),
		})
	}
	return details
}

func WriteSuccess(w http.ResponseWriter, r *http.Request, statusCode int, data interface{}) {
	response := &GlobalResponse{
		Code:    errorwrapper.ErrCode(statusCode),
		Message: locale.TranslateString(r.Context(), "generic_success", nil, nil),
	}

	if data != nil {
		response.Data = data
	}

	// Assign header Request ID.
	w.Header().Set(constant.HTTPHeaderRequestID, constant.DefaultString)

	render.Status(r, statusCode)
	render.DefaultResponder(w, r, response)
}

// Success response is based on given struct
func WriteCustomSuccess(w http.ResponseWriter, r *http.Request, statusCode int, response interface{}) {
	// Assign header Request ID.
	w.Header().Set(constant.HTTPHeaderRequestID, constant.DefaultString)

	render.Status(r, statusCode)
	render.DefaultResponder(w, r, response)
}
