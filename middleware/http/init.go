package http

import "github.com/assignment-amori/internal/entity"

func New(
	appConfig entity.AppConfig,
) *MiddlewareModule {
	return &MiddlewareModule{
		appConfig: appConfig,
	}
}
