package errorwrapper

import (
	"context"
	"fmt"
	"regexp"

	"github.com/assignment-amori/internal/constant"
	"github.com/assignment-amori/pkg/locale"
	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
)

// New creates new error wrapper with options.
func New(id ErrID, options ...Option) error {
	return wrapErr(id, 1, options...)
}

func wrapErr(id ErrID, caller int, options ...Option) *errWrapper {
	s := &errWrapper{
		id:         id,
		devMessage: msgUnclassifiedErr,
		metadata:   MetaKV{},
		msgParams:  _UserMessageParams{},
	}
	s.stack = callers(caller)

	// Put the default user message params at the end to ensure
	// the defaults don't override the added one.
	defaultUserMsgParams := holder.getDefaultUserMsgParam(id)
	if len(defaultUserMsgParams) != constant.DefaultInt {
		for key, val := range defaultUserMsgParams {
			opt := WithUserMsgParams(UserMessageParams{
				UserMsgKey(key): val,
			})

			options = append(options, opt)
		}
	}

	for _, opt := range options {
		opt(s)
	}

	return s
}

// Wrap error wrapper to add attributes.
// Return nil if no error.
func Wrap(err error, id ErrID, options ...Option) error {
	if err == nil {
		return nil
	}

	// Append devMessage with current error.
	options = append(options, WithDevMessage(err.Error()))

	// Append With validation error if it's coming from validator v10.
	vErr, ok := err.(validator.ValidationErrors)
	if ok {
		options = append(options, WithValidatorError(vErr))
	}

	// If not error wrapper, then return as is.
	errCast, ok := err.(*errWrapper)
	if !ok || errCast == nil {
		return wrapErr(id, 1, options...)
	}

	errWrap := errCast.DeepCopy()

	// Clear current dev message.
	errWrap.devMessage = ""

	for _, opt := range options {
		opt(errWrap)
	}

	return errWrap
}

// GetMetaData get meta data from error
// return nil if not error wrapper
func (la *errWrapper) GetMetaData() MetaKV {
	return la.getMetadata()
}

// Error get complete error.
func (la *errWrapper) Error() string {
	return fmt.Sprintf("error_id: %s, http_status: %d, error_code: %d, dev_message: %s",
		la.id,
		la.GetHTTPStatus(),
		la.GetErrorCode(),
		la.devMessage,
	)
}

// GetErrorLine get error line from stack
func (la *errWrapper) GetErrorLine() string {
	var fun, file, stack string
	var line int

	if la.stack != nil {
		fun, file, line = la.getFuncFileLine()
		stack = fmt.Sprintf("[%s]%s:%d", fun, file, line)
	}
	return trimRootPath(stack)
}

// Code get error code.
func (la *errWrapper) Code() ErrCode {
	return holder.getCode(la.id)
}

// UserMessage Get error message.
func (la *errWrapper) UserMessage(ctx context.Context) string {
	pluralCount, ok := la.msgParams[string(PluralCount)]
	if !ok {
		// default plural count is nil
		pluralCount = nil
	}

	// Try to translate the message params.
	for key, value := range la.msgParams {
		if isParamTranslateable(UserMsgKey(key)) {
			msgId := fmt.Sprintf("%s", value)
			la.msgParams[key] = locale.TranslateString(ctx, msgId, nil, nil)
		}
	}

	return locale.TranslateString(ctx, holder.getUserMsgTranslationID(la.id), la.msgParams, pluralCount)
}

// GetHTTPStatus get status http from error id.
func (la *errWrapper) GetHTTPStatus() int {
	return holder.getHTTPStatus(la.id)
}

// GetErrorCode get error code in string.
func (la *errWrapper) GetErrorCode() ErrCode {
	return holder.getCode(la.id)
}

// GetErrorID get error code.
func (la *errWrapper) GetErrorID() ErrID {
	return la.id
}

// GetIsRetryable get error retryable in bool.
func (la *errWrapper) GetIsRetryable() bool {
	return holder.getRetryable(la.id)
}

func (la *errWrapper) GetValidatorErrs() validator.ValidationErrors {
	return la.validatorErr
}

// StackTrace get error stack trace.
func (la *errWrapper) StackTrace() errors.StackTrace {
	f := make([]errors.Frame, len(*la.stack))
	for i := 0; i < len(f); i++ {
		f[i] = errors.Frame((*la.stack)[i])
	}
	return f
}

func (la *errWrapper) DeepCopy() *errWrapper {
	metadata := make(MetaKV)
	msgParams := make(_UserMessageParams)
	for key, val := range la.metadata {
		metadata[key] = val
	}
	for key, val := range la.msgParams {
		msgParams[key] = val
	}
	copyLa := errWrapper{
		id:         la.id,
		devMessage: la.devMessage,
		metadata:   metadata,
		msgParams:  msgParams,
		stack:      la.stack,
	}
	return &copyLa
}

// CastToErrorWrapper cast error to error wrapper.
func CastToErrorWrapper(err error) *errWrapper {
	if err == nil {
		return nil
	}

	eWrapper, ok := err.(*errWrapper)
	if !ok || eWrapper == nil {
		return Wrap(err, ErrIDUnclassified).(*errWrapper)
	}

	return eWrapper
}

func NewMeta() MetaKV {
	return MetaKV{}
}

func (meta MetaKV) Set(key MetaKey, value interface{}) {
	meta[key] = value
}

// Regex.
var (
	regexErrorMessage = regexp.MustCompile("[^a-zA-Z0-9 _-]+")
)

// RemoveUniqueErrorMsg get error for allowed character only.
func RemoveUniqueErrorMsg(msg string) string {
	return regexErrorMessage.ReplaceAllString(msg, "")
}

func isParamTranslateable(param UserMsgKey) bool {
	return param == DataOperator ||
		param == SingleOperator ||
		param == DoubleOperator ||
		param == ComparatorUnit ||
		param == Reason
}
