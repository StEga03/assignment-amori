package channel

import (
	"context"

	"github.com/assignment-amori/internal/entity"
	"github.com/assignment-amori/internal/entity/generic"
	"github.com/assignment-amori/pkg/errorwrapper"
	timeutils "github.com/assignment-amori/pkg/time_utils"
	"github.com/sashabaranov/go-openai"
)

func (u *Usecase) CreateMessageInChannel(ctx context.Context, req entity.MessageUCRequest) (entity.MessageResponse, error) {
	var (
		result entity.MessageResponse
		err    error
	)

	_, err = u.channelRepo.GetByIDAndUserID(ctx, req.ChannelID, 10)
	if err != nil {
		return result, err
	}

	messageId, err := u.sf.NextID()
	if err != nil {
		return result, errorwrapper.Wrap(err, errorwrapper.ErrIDFailedToGenerateID)
	}

	now := timeutils.Now()
	newMessage := entity.NewMessageParams{
		ID:          messageId,
		ChannelID:   req.ChannelID,
		SenderType:  "user",
		SenderID:    10,
		ContentType: "text",
		Content:     req.Body,
		MetaInfo: generic.MetaInfo{
			CreatedAt: now,
			UpdatedAt: now,
		},
	}
	_, err = u.messageRepo.CreateMessage(ctx, newMessage, nil)
	if err != nil {
		return result, err
	}

	result = entity.MessageResponse{
		ID:        messageId,
		ChannelID: req.ChannelID,
		Body:      req.Body,
		Timestamp: now,
	}

	userMsgParam := entity.GetMessageParams{
		ChannelID: req.ChannelID,
		Limit:     10,
		Offset:    0,
	}
	userMessages, err := u.messageRepo.GetMessageByChannelID(ctx, userMsgParam)

	var messages []openai.ChatCompletionMessage
	for i := len(userMessages) - 1; i >= 0; i-- {
		message := openai.ChatCompletionMessage{
			Role:    userMessages[i].SenderType,
			Content: userMessages[i].Content,
		}
		messages = append(messages, message)
	}

	messages = []openai.ChatCompletionMessage{
		{
			Role:    openai.ChatMessageRoleUser,
			Content: req.Body,
		},
	}

	param := entity.ChatCompletionParams{
		Messages: messages,
	}
	response, err := u.openaiRepo.CreateChatCompletion(ctx, param)
	if err != nil {
		return result, err
	}

	messageId, err = u.sf.NextID()
	if err != nil {
		return result, errorwrapper.Wrap(err, errorwrapper.ErrIDFailedToGenerateID)
	}

	now = timeutils.Now()
	responseParam := entity.NewMessageParams{
		ID:          messageId,
		ChannelID:   req.ChannelID,
		SenderType:  "assistant",
		SenderID:    1,
		ContentType: "text",
		Content:     response.Choices[0].Message.Content,
		MetaInfo: generic.MetaInfo{
			CreatedAt: now,
			UpdatedAt: now,
		},
	}
	_, err = u.messageRepo.CreateMessage(ctx, responseParam, nil)
	if err != nil {
		return result, err
	}

	return result, nil
}
