package entity

import "github.com/sashabaranov/go-openai"

type ChatCompletionParams struct {
	Messages []openai.ChatCompletionMessage `json:"messages"`
}
