package openai

import (
	"context"

	"github.com/assignment-amori/internal/entity"
	"github.com/sashabaranov/go-openai"
)

type OpenaiResources interface {
	CreateChatCompletion(ctx context.Context, param entity.ChatCompletionParams) (openai.ChatCompletionResponse, error)
}

type Repository struct {
	openaiService OpenaiResources
}
