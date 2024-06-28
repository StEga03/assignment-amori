package locale

import (
	"context"
	"log"

	"github.com/assignment-amori/internal/constant"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

// TranslateString represent method for do translation based on the locale
// the function will read the locale via context and the default lang is en-us.
func TranslateString(ctx context.Context, messageID string, args interface{}, pluralCount interface{}) string {
	lang, ok := ctx.Value(constant.ContextKeyLanguage).(string)
	if !ok {
		lang = constant.LangENUS
	}

	localizer := i18n.NewLocalizer(langBundle, lang)
	result, err := localizer.Localize(&i18n.LocalizeConfig{
		MessageID:    messageID,
		TemplateData: args,

		// Used for differentiating between singular and plural messages if needed.
		PluralCount: pluralCount,

		// Used for fallback if the messageID is not found (esp. user message params that doesn't have translation).
		DefaultMessage: &i18n.Message{
			ID:    messageID,
			Other: messageID,
		},
	})
	if err != nil {
		log.Fatalf("failed to localize message: %v", err)
		return result
	}

	return result
}
