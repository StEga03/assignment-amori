package http

import (
	"encoding/json"
	"net/http"

	"github.com/assignment-amori/internal/constant"
	"github.com/assignment-amori/internal/entity/generic"
	pkghttp "github.com/assignment-amori/pkg/http"
)

// GenericMiddleware Deprecated.
func GenericMiddleware(next func(http.ResponseWriter, *http.Request) (interface{}, error)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			err          error
			jsonResponse []byte
		)

		// Set the Content-Type header to indicate JSON response.
		w.Header().Set(constant.HTTPHeaderContentType, constant.HTTPContentTypeJSON)

		// Call the handler and capture its response and error.
		response, err := next(w, r)
		mapData := generic.Response[interface{}]{
			Data:    response,
			Message: "Successfully executed",
			Success: err == nil,
		}

		if err != nil {
			// Catch the errors.
			mapData.Data = nil
			mapData.Message = err.Error()
			// Safe operation, without breaking changes or chat_flow.
			jsonResponse, err = json.Marshal(mapData) //nolint:all

			// Write the JSON response.
			w.WriteHeader(http.StatusBadRequest)
			w.Write(jsonResponse)
			return
		}

		// Marshal the response data to JSON.
		jsonResponse, err = json.Marshal(mapData)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Write the JSON response.
		w.WriteHeader(http.StatusOK)
		w.Write(jsonResponse) //nolint:all
	}
}

type respMiddleware struct{}

func (hm *MiddlewareModule) GetResponseWriterMiddleware() *respMiddleware {
	return &respMiddleware{}
}

func (mw *respMiddleware) Filter(h generic.HTTPHandleFunc) generic.HTTPHandleFunc {
	return func(w http.ResponseWriter, r *http.Request) (data interface{}, err error) {
		data, err = h(w, r)
		if err != nil {
			pkghttp.WriteError(w, r, err, nil)
			return data, err
		}

		pkghttp.WriteSuccess(w, r, http.StatusOK, data)
		return data, err
	}
}
