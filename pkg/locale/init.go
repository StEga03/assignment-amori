package locale

import (
	"embed"
	"io/fs"

	"github.com/assignment-amori/pkg/json"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

//go:embed translation/*
var translation embed.FS

var langBundle *i18n.Bundle

// NewLocale represent initializer function of locale package.
func NewLocale() error {

	langBundle = i18n.NewBundle(language.AmericanEnglish)
	langBundle.RegisterUnmarshalFunc("json", json.Unmarshal)

	err := fs.WalkDir(translation, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() {
			return nil
		}

		_, err = langBundle.LoadMessageFileFS(translation, path)
		if err != nil {
			return err
		}
		return nil
	})

	return err
}
