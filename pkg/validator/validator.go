package validator

import (
	"context"
	"sync"

	"github.com/assignment-amori/pkg/locale"
	"github.com/go-playground/validator/v10"
)

var (
	validate *validator.Validate
	vOnce    sync.Once
)

// New represent initial function for validator.
func New() error {
	var err error
	vOnce.Do(func() {
		validate = validator.New()
		
	})
	return err
}

// ValidateStruct represent a function to validate struct using validator and return the error using laas err format.
func ValidateStruct(ctx context.Context, s interface{}) error {
	return validate.StructCtx(ctx, s)
}

func ValidateVar(ctx context.Context, s interface{}, tag string) error {
	return validate.VarCtx(ctx, s, tag)
}

func BuildErrorMessage(ctx context.Context, e validator.FieldError) string {
	// Build user message based on validator tag.
	switch e.Tag() {
	case "gt":
		return locale.TranslateString(ctx, "validator_min", map[string]interface{}{
			"Field": e.Field(),
			"Param": e.Param(),
		}, nil)
	case "lt":
		return locale.TranslateString(ctx, "validator_max", map[string]interface{}{
			"Field": e.Field(),
			"Param": e.Param(),
		}, nil)
	case "required":
		return locale.TranslateString(ctx, "validator_required", map[string]interface{}{
			"Field": e.Field(),
		}, nil)
	case "latitude":
		return locale.TranslateString(ctx, "validator_latitude", map[string]interface{}{
			"Field": e.Field(),
		}, nil)
	case "longitude":
		return locale.TranslateString(ctx, "validator_longitude", map[string]interface{}{
			"Field": e.Field(),
		}, nil)
	case "alpha":
		return locale.TranslateString(ctx, "validator_alpha", map[string]interface{}{
			"Field": e.Field(),
		}, nil)
	case "numeric":
		return locale.TranslateString(ctx, "validator_numeric", map[string]interface{}{
			"Field": e.Field(),
		}, nil)
	default:
		return locale.TranslateString(ctx, "validator_default", map[string]interface{}{
			"Field": e.Field(),
		}, nil)
	}
}
