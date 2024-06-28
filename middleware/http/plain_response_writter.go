package http

import (
	"net/http"

	"github.com/assignment-amori/internal/entity/generic"
	pkghttp "github.com/assignment-amori/pkg/http"
)

type plainRespMiddleware struct {
}

func (hm *MiddlewareModule) GetPlainResponseWriterMiddleware() *plainRespMiddleware {
	return &plainRespMiddleware{}
}

func (mw *plainRespMiddleware) Filter(h generic.HTTPHandleFunc) generic.HTTPHandleFunc {
	return func(w http.ResponseWriter, r *http.Request) (data interface{}, err error) {
		data, err = h(w, r)
		if err != nil {
			pkghttp.WriteError(w, r, err, nil)
			return data, err
		}

		return data, err
	}
}
