package openai

import (
	"context"

	"github.com/assignment-amori/internal/entity"
	"github.com/sashabaranov/go-openai"
)

func (m *Module) CreateChatCompletion(ctx context.Context, param entity.ChatCompletionParams) (openai.ChatCompletionResponse, error) {
	var (
		resp openai.ChatCompletionResponse
		err  error
	)

	req := openai.ChatCompletionRequest{
		Model:    openai.GPT4o,
		Messages: param.Messages,
	}

	resp, err = m.Client.CreateChatCompletion(ctx, req)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
