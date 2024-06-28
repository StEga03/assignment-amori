package http

import (
	"net/http"
	"path/filepath"

	pkghttp "github.com/assignment-amori/pkg/http"
	"github.com/assignment-amori/pkg/utils"
	uuidPkg "github.com/assignment-amori/pkg/uuid"
)

func (h *Handler) UploadMessageSource(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	ctx := r.Context()

	// Also add current date as folder to avoid collision with the UUID
	fileLocation := utils.GenerateUniqueFilePath("resources/file", uuidPkg.New())
	folderLocation := filepath.Dir(fileLocation)
	fileName := filepath.Base(fileLocation)

	// Read uploaded file, and store it to local file system.
	fileDetails, err := pkghttp.ProcessUpload(r, folderLocation, fileName)
	if err != nil {
		return nil, err
	}

	return h.fileUC.WhatsappParser(ctx, fileDetails)
}
