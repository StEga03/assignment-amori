package channel

import (
	"context"

	"github.com/assignment-amori/internal/entity"
	"github.com/assignment-amori/pkg/consistency"
	"github.com/sashabaranov/go-openai"
)

//go:generate mockgen -package=channel -source=types.go -destination=channel_mock_test.go
type consistencyResource interface {
	RunAsUnit(ctx context.Context, action func(celTemp *consistency.ConsistencyElement) error) error
}

type channelResource interface {
	CreateChannel(ctx context.Context, param entity.NewChannelParams, cel *consistency.ConsistencyElement) (uint64, error)
	GetByID(ctx context.Context, id uint64) (entity.Channel, error)
	GetByUserID(ctx context.Context, userId string) ([]entity.Channel, error)
	GetByIDAndUserID(ctx context.Context, id, userId uint64) (entity.Channel, error)
}

type messageResource interface {
	// Message.
	CreateMessage(ctx context.Context, param entity.NewMessageParams, cel *consistency.ConsistencyElement) (uint64, error)
	CreateMessageBulk(ctx context.Context, param []entity.NewMessageParams, cel *consistency.ConsistencyElement) error
	GetMessageByID(ctx context.Context, param entity.GetMessageParams) (entity.Message, error)
	GetMessageByChannelID(ctx context.Context, param entity.GetMessageParams) ([]entity.Message, error)

	// Message Input.
	CreateMessageInput(ctx context.Context, param entity.NewMessageInputParams, cel *consistency.ConsistencyElement) (uint64, error)
	GetMessageInputByID(ctx context.Context, param entity.GetMessageInputParams) (entity.MessageInput, error)
	GetMessageInputByChannelID(ctx context.Context, param entity.GetMessageParams) ([]entity.MessageInput, error)

	// Message Source.
	CreateMessageSource(ctx context.Context, param []entity.NewMessageSourceParams, cel *consistency.ConsistencyElement) error
	GetMessageSourceByID(ctx context.Context, param entity.GetMessageSourceParams) (entity.MessageSource, error)
	GetMessageSourceByMessageInputID(ctx context.Context, param entity.GetMessageSourceParams) ([]entity.MessageSource, error)
}

type openaiResource interface {
	CreateChatCompletion(ctx context.Context, param entity.ChatCompletionParams) (openai.ChatCompletionResponse, error)
}

type userResource interface {
	CreateMessage(ctx context.Context, param entity.NewUserParams, cel *consistency.ConsistencyElement) (uint64, error)
	GetByID(ctx context.Context, id uint64) (entity.User, error)
}

type sonyFlakeResource interface {
	NextID() (uint64, error)
}

type Usecase struct {
	consistency consistencyResource
	channelRepo channelResource
	messageRepo messageResource
	openaiRepo  openaiResource
	userRepo    userResource
	sf          sonyFlakeResource
}
