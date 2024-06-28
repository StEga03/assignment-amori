package openai

import (
	"github.com/assignment-amori/internal/entity"
	"github.com/sashabaranov/go-openai"
)

func New(openaiConfig entity.OpenAIConfig) *Module {
	return &Module{
		Client: openai.NewClient(openaiConfig.APIKey),
	}
}
