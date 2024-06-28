package http

import (
	"errors"
	"io"
	"net/http"
	"os"
	"path"
	"path/filepath"

	"github.com/assignment-amori/pkg/errorwrapper"
)

func ProcessUpload(request *http.Request, folderPath string, newFileName string) (fileDetails FileDetails, err error) {
	return ProcessUploadFromParam(request, DefaultFileParameter, folderPath, newFileName)
}

func ProcessUploadFromParam(request *http.Request, fileParam string, folderPath string, newFileName string) (fileDetails FileDetails, err error) {
	if _, err := os.Stat(folderPath); errors.Is(err, os.ErrNotExist) {
		err := os.MkdirAll(folderPath, os.ModePerm)
		if err != nil {
			return fileDetails, err
		}
	}

	err = request.ParseMultipartForm(1024)
	if err != nil {
		return fileDetails, err
	}

	uploadedFile, handler, err := request.FormFile(fileParam)
	if err == http.ErrMissingFile {
		return fileDetails, errorwrapper.Wrap(err, errorwrapper.ErrIDHandlerBadRequest, errorwrapper.WithDevMessage("Invalid file"))
	} else if err != nil {
		return fileDetails, err
	}
	defer uploadedFile.Close()

	fileDetails = FileDetails{
		OriginalFileName: handler.Filename,
		FilePath:         path.Join(folderPath, newFileName),
		Extension:        filepath.Ext(handler.Filename),
		Size:             handler.Size,
		Header:           handler.Header,
	}

	targetFile, err := os.OpenFile(fileDetails.FilePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return fileDetails, err
	}
	defer targetFile.Close()

	if _, err := io.Copy(targetFile, uploadedFile); err != nil {
		return fileDetails, err
	}

	return
}
