package errorwrapper

import (
	"fmt"

	"github.com/assignment-amori/internal/constant"
	"github.com/go-playground/validator/v10"
)

// WithDevMessage add dev message
func WithDevMessage(m string) Option {
	return func(err *errWrapper) {
		if err.devMessage == "" || err.devMessage == msgUnclassifiedErr {
			err.devMessage = m
		} else {
			err.devMessage = fmt.Sprintf("%s: %s", err.devMessage, m)
		}
	}
}

// WithValidatorError represent a function for appending error from golang validator v10.
func WithValidatorError(vErrs validator.ValidationErrors) Option {
	return func(err *errWrapper) {
		// Immediately return if validator errors is nil.
		if vErrs == nil {
			return
		}

		// Make sure if there's nothing error.
		if len(vErrs) < 1 {
			return
		}

		var validatorData []map[string]interface{}
		for _, vErr := range vErrs {
			validatorData = append(validatorData, map[string]interface{}{
				"field": vErr.Field(),
				"value": vErr.Value(),
				"tag":   vErr.Tag(),
				"error": vErr.Error(),
			})
		}

		err.metadata[metaKeyValidator] = validatorData
		err.id = ErrIDValidationNotPassed
		err.validatorErr = vErrs
	}
}

// WithMetaData add metadata.
// don't use default key: error_code", "user_message", "error_line"
func WithMetaData(metadata MetaKV) Option {
	return func(err *errWrapper) {
		for key, value := range metadata {
			if key != MetaKey(constant.DefaultString) {
				err.metadata[key] = value
			}
		}
	}
}

// WithMetaData add metadata.
// don't use default key: error_code", "user_message", "error_line"
func WithMetadataValue(key MetaKey, value interface{}) Option {
	return func(err *errWrapper) {
		if err.metadata == nil {
			err.metadata = MetaKV{}
		}

		if key != MetaKey(constant.DefaultString) {
			err.metadata[key] = value
		}
	}
}

// WithUserMsgParams add params for User Message.
func WithUserMsgParams(params UserMessageParams) Option {
	return func(err *errWrapper) {
		for key, value := range params {
			if key != UserMsgKey(constant.DefaultString) {
				// If msgParam already exist, don't replace it
				// to ensure user message uses the deepest WithUserMsgParams
				_, exist := err.msgParams[string(key)]
				if !exist {
					err.msgParams[string(key)] = value
				}
			}
		}
	}
}
