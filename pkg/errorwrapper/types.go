package errorwrapper

import "github.com/go-playground/validator/v10"

type Config struct {
	RepoRoot string `yaml:"repo_root" json:"repo_root"`
}

type (
	wrapperCentralizeError struct {
		ErrorDetail wrapperErrorDetail     `json:"detail"`
		ErrorList   map[ErrID]wrapperError `json:"errors"`
	}

	wrapperErrorDetail struct {
		CodePrefix string `json:"code_prefix"`
	}

	wrapperError struct {
		HTTPStatus               int                `json:"http_status"`
		Code                     ErrCode            `json:"code"`
		Retryable                bool               `json:"retryable"`
		UserMessageTranslationID string             `json:"user_message_translation_id"`
		UserMessageDefaultParams _UserMessageParams `json:"user_message_default_param"`
	}

	errWrapperHolder struct {
		Data   map[ErrID]wrapperError
		Config Config
	}

	// errWrapper error attributes.
	// httpStatus will be used for http status response on handler.
	// code: error code for developer.
	// userMessage: error message show to user.
	// devMessage: exclusive error for developer.
	// metadata: additional error.
	// stack: determine stack error line.
	errWrapper struct {
		id           ErrID
		devMessage   string
		msgParams    _UserMessageParams
		metadata     MetaKV
		stack        *stack
		validatorErr validator.ValidationErrors
		error
	}
)

// ErrCode code alias.
type (
	ErrCode int32
	ErrID   string
)

// Metadata types
type (
	MetaKey string
	MetaKV  map[MetaKey]interface{}
)

// User message param types.
type (
	UserMsgKey        string
	UserMessageParams map[UserMsgKey]interface{}

	// Due to the nature of text/template, the internal map key should be string.
	// Otherwise, it will error can't evaluate field in type.
	_UserMessageParams map[string]interface{}
)

type Option func(*errWrapper)
