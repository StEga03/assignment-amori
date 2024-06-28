package errorwrapper

import (
	"embed"
	"io/fs"
	"net/http"
	"strconv"

	"github.com/assignment-amori/internal/constant"
	"github.com/assignment-amori/pkg/json"
)

//go:embed error-list/*
var errList embed.FS

var holder = &errWrapperHolder{}

func NewErrorWrapper(config Config) error {
	structNormalize, err := constructErrorMap(errList)
	if err != nil {
		return err
	}

	holder = &errWrapperHolder{
		Data:   structNormalize,
		Config: config,
	}
	return nil
}

func constructErrorMap(data embed.FS) (map[ErrID]wrapperError, error) {
	errMap := make(map[ErrID]wrapperError)

	err := fs.WalkDir(data, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() {
			return nil
		}

		file, err := data.ReadFile(path)
		if err != nil {
			return err
		}

		var wrapper wrapperCentralizeError
		// Note that this will overwrite the previous value of the key
		// TODO: Make a test which will check if there's any duplicate key
		err = json.Unmarshal(file, &wrapper)
		if err != nil {
			return err
		}

		for k, v := range wrapper.ErrorList {
			finalCode := wrapper.ErrorDetail.CodePrefix + strconv.Itoa(int(v.Code))
			finalCodeInt, err := strconv.Atoi(finalCode)
			if err != nil {
				return err
			}

			v.Code = ErrCode(finalCodeInt)
			errMap[k] = v
		}

		return nil
	})

	return errMap, err
}

// private func
func constructToMap(data []byte) (map[ErrID]wrapperError, error) {
	var result map[ErrID]wrapperError
	var err = json.Unmarshal(data, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (e *errWrapperHolder) WithErrorList(data []byte) (err error) {
	structNormalize, err := constructToMap(data)
	if err != nil {
		return err
	}

	holder.Data = structNormalize

	return err
}

func (e *errWrapperHolder) getCode(msgCode ErrID) ErrCode {
	if errDetail, ok := e.Data[msgCode]; ok {
		return errDetail.Code
	}

	return ErrCodeUnclassified
}

func (e *errWrapperHolder) getHTTPStatus(msgCode ErrID) int {
	if errDetail, ok := e.Data[msgCode]; ok {
		return errDetail.HTTPStatus
	}

	return http.StatusInternalServerError
}

func (e *errWrapperHolder) getUserMsgTranslationID(msgCode ErrID) string {
	if errDetail, ok := e.Data[msgCode]; ok {
		return errDetail.UserMessageTranslationID
	}

	return defaultUserMessageTranslationID
}

func (e *errWrapperHolder) getRetryable(msgCode ErrID) bool {
	if errDetail, ok := e.Data[msgCode]; ok {
		return errDetail.Retryable
	}

	return constant.BoolFalse
}

func (e *errWrapperHolder) getDefaultUserMsgParam(msgCode ErrID) _UserMessageParams {
	if errDetail, ok := e.Data[msgCode]; ok {
		return errDetail.UserMessageDefaultParams
	}

	return _UserMessageParams{}
}
