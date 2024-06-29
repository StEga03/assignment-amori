package channel

import (
	"context"
	"fmt"

	"github.com/assignment-amori/internal/constant"
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

	user, err := u.userRepo.GetUserByContext(ctx)
	if err != nil {
		return result, err
	}

	_, err = u.channelRepo.GetByIDAndUserID(ctx, req.ChannelID, user.ID)
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
		SenderType:  openai.ChatMessageRoleUser,
		SenderID:    user.ID,
		ContentType: constant.ContentTypeText,
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
		Type:      openai.ChatMessageRoleUser,
		Body:      req.Body,
		Timestamp: now,
	}

	err = u.getAIResponse(ctx, req)
	if err != nil {
		return result, err
	}

	return result, nil
}

func (u *Usecase) getAIResponse(ctx context.Context, req entity.MessageUCRequest) error {
	userMsgParam := entity.GetMessageParams{
		ChannelID: req.ChannelID,
		Limit:     constant.DefaultLimit,
		Offset:    constant.DefaultOffset,
	}
	messageInput, err := u.messageRepo.GetMessageInputByChannelID(ctx, userMsgParam)
	if err != nil {
		return err
	}

	userMsgSourceParam := entity.GetMessageSourceParams{
		MessageInputID: messageInput[0].ID,
		Limit:          50,
		Offset:         0,
	}
	messageSources, err := u.messageRepo.GetMessageSourceByMessageInputID(ctx, userMsgSourceParam)
	if err != nil {
		return err
	}

	var messages []openai.ChatCompletionMessage
	for i := len(messageSources) - 1; i >= 0; i-- {
		message := openai.ChatCompletionMessage{
			Role:    openai.ChatMessageRoleUser,
			Content: fmt.Sprintf("[%s] %s: %s", messageSources[i].SentAt, messageSources[i].Sender, messageSources[i].Content),
		}
		messages = append(messages, message)
	}

	userMsgParam = entity.GetMessageParams{
		ChannelID: req.ChannelID,
		Limit:     constant.DefaultLimit,
		Offset:    constant.DefaultOffset,
	}
	userMessages, err := u.messageRepo.GetMessageByChannelID(ctx, userMsgParam)

	for i := len(userMessages) - 1; i >= 0; i-- {
		message := openai.ChatCompletionMessage{
			Role:    userMessages[i].SenderType,
			Content: userMessages[i].Content,
		}
		messages = append(messages, message)
	}

	message := openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleAssistant,
		Content: constant.MessageContextPrompt,
	}
	messages = append(messages, message)
	
	param := entity.ChatCompletionParams{
		Messages: messages,
	}
	response, err := u.openaiRepo.CreateChatCompletion(ctx, param)
	if err != nil {
		return err
	}

	messageId, err := u.sf.NextID()
	if err != nil {
		return errorwrapper.Wrap(err, errorwrapper.ErrIDFailedToGenerateID)
	}

	now := timeutils.Now()
	responseParam := entity.NewMessageParams{
		ID:          messageId,
		ChannelID:   req.ChannelID,
		SenderType:  openai.ChatMessageRoleAssistant,
		SenderID:    1, // Hardcoded Assistant's UserID.
		ContentType: constant.ContentTypeText,
		Content:     response.Choices[0].Message.Content,
		MetaInfo: generic.MetaInfo{
			CreatedAt: now,
			UpdatedAt: now,
		},
	}
	_, err = u.messageRepo.CreateMessage(ctx, responseParam, nil)
	if err != nil {
		return err
	}

	return nil
}
