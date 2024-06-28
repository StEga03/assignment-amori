package openai

import (
	"context"

	"github.com/assignment-amori/internal/entity"
	"github.com/sashabaranov/go-openai"
)

func (r *Repository) CreateChatCompletion(ctx context.Context, param entity.ChatCompletionParams) (openai.ChatCompletionResponse, error) {
	return r.openaiService.CreateChatCompletion(ctx, param)
}
