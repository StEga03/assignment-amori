package channel

import (
	"context"

	"github.com/assignment-amori/internal/constant"
	"github.com/assignment-amori/internal/entity"
	"github.com/assignment-amori/pkg/consistency"
	"github.com/assignment-amori/pkg/errorwrapper"
	"github.com/assignment-amori/pkg/utils"
)

func (u *Usecase) CreateChannel(ctx context.Context, req entity.NewChannelUCRequest) (uint64, error) {
	var (
		result uint64
		err    error
	)

	channelId, err := u.sf.NextID()
	if err != nil {
		return result, errorwrapper.Wrap(err, errorwrapper.ErrIDFailedToGenerateID)
	}

	messageInputId, err := u.sf.NextID()
	if err != nil {
		return result, errorwrapper.Wrap(err, errorwrapper.ErrIDFailedToGenerateID)
	}

	newChannel := entity.NewChannelParams{
		ID:     channelId,
		UserID: 10,
		Name:   req.Name,
	}

	newMessageInput := entity.NewMessageInputParams{
		ID:              messageInputId,
		ChannelID:       channelId,
		Source:          "whatsapp",
		Sender:          "Alice",
		Receiver:        "Bob",
		ReceiverPronoun: "Him",
	}

	var newMessageSources []entity.NewMessageSourceParams
	for _, msgSource := range req.MessageSource {
		sentAt, err := utils.ParseStringToTime(msgSource.SentAt, constant.CustomDateFormat)
		if err != nil {
			return result, err
		}

		newMessageSource := entity.NewMessageSourceParams{
			MessageInputID: messageInputId,
			Sender:         msgSource.Sender,
			ContentType:    "text",
			Content:        msgSource.Body,
			SentAt:         sentAt,
		}
		newMessageSources = append(newMessageSources, newMessageSource)
	}

	err = u.consistency.RunAsUnit(ctx, func(celTemp *consistency.ConsistencyElement) error {
		_, err = u.channelRepo.CreateChannel(ctx, newChannel, celTemp)
		if err != nil {
			return errorwrapper.Wrap(err, errorwrapper.ErrIDInsertDB)
		}

		_, err = u.messageRepo.CreateMessageInput(ctx, newMessageInput, celTemp)
		if err != nil {
			return errorwrapper.Wrap(err, errorwrapper.ErrIDInsertDB)
		}

		err = u.messageRepo.CreateMessageSource(ctx, newMessageSources, celTemp)
		if err != nil {
			return errorwrapper.Wrap(err, errorwrapper.ErrIDInsertDB)
		}

		return nil
	})
	if err != nil {
		return result, err
	}

	return channelId, nil
}
