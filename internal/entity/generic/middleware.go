package generic

import (
	"net/http"
)

type (
	// HTTPHandleFunc converting handler with return error to http.Handler.
	HTTPHandleFunc = func(w http.ResponseWriter, r *http.Request) (interface{}, error)

	// HTTPHandler http.Handler func.
	HTTPHandler = func(w http.ResponseWriter, r *http.Request)
)
