package http

import "github.com/assignment-amori/pkg/errorwrapper"

type GlobalResponse struct {
	Code      errorwrapper.ErrCode `json:"code"`
	Message   string               `json:"message"`
	Retryable bool                 `json:"retryable"`
	Details   []ErrorDetail        `json:"details,omitempty"`
	Data      interface{}          `json:"data,omitempty"`
}

type ErrorDetail struct {
	Path string `json:"path"`
	Info string `json:"info"`
}

type FileDetails struct {
	OriginalFileName string
	FilePath         string
	Extension        string
	Size             int64
	Header           map[string][]string
}
