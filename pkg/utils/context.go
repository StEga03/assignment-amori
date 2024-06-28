package utils

import (
	"context"
	"fmt"

	"github.com/assignment-amori/internal/constant"
)

// Generate usecase key using group level.
func GenerateUsecaseKey(ctx context.Context, cases string) (context.Context, string) {
	level, ok := ctx.Value(constant.ContextKeyAPISubModule).(string)
	if !ok {
		level = constant.DefaultString
	}

	key := fmt.Sprintf(cases, level)

	ctx = context.WithValue(ctx, constant.ContextKeyUsecaseKey, key)
	return ctx, key
}

// Get usecase key stored within context.
func GetUsecaseKey(ctx context.Context) string {
	key, ok := ctx.Value(constant.ContextKeyUsecaseKey).(string)
	if !ok {
		return constant.DefaultString
	}
	return key
}

// Get module key stored within context.
func GetModuleKey(ctx context.Context) constant.Module {
	key, ok := ctx.Value(constant.ContextKeyAPIModule).(constant.Module)
	if !ok {
		return constant.Module(constant.DefaultString)
	}

	return key
}

// Get module key stored within context.
func GetSubModuleKey(ctx context.Context) string {
	key, ok := ctx.Value(constant.ContextKeyAPISubModule).(string)
	if !ok {
		return constant.DefaultString
	}
	return key
}
