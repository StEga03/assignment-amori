package file

import (
	"context"
	"os"

	"github.com/assignment-amori/pkg/errorwrapper"
	"github.com/assignment-amori/pkg/http"
	"github.com/assignment-amori/pkg/whatsapp"
)

func (u *Usecase) WhatsappParser(ctx context.Context, fileDetails http.FileDetails) ([]whatsapp.Message, error) {
	var (
		result []whatsapp.Message
		err    error
	)

	data, err := os.ReadFile(fileDetails.FilePath)
	if err != nil {
		return result, errorwrapper.Wrap(err, errorwrapper.ErrIDProcessFileFailedToRead)
	}

	content := string(data)

	result, err = whatsapp.ParseString(content)
	if err != nil {
		return result, errorwrapper.Wrap(err, errorwrapper.ErrParsing)
	}

	return result, nil
}
