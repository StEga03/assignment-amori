package channel

import (
	"context"

	"github.com/assignment-amori/internal/entity"
)

func (u *Usecase) GetMessageInChannel(ctx context.Context, req entity.MessageUCRequest) ([]entity.MessageResponse, error) {
	var (
		result []entity.MessageResponse
		err    error
	)

	_, err = u.channelRepo.GetByIDAndUserID(ctx, req.ChannelID, 10)
	if err != nil {
		return result, err
	}

	userMsgParam := entity.GetMessageParams{
		ChannelID: req.ChannelID,
		Limit:     10,
		Offset:    0,
	}
	userMessages, err := u.messageRepo.GetMessageByChannelID(ctx, userMsgParam)

	for _, userMessage := range userMessages {
		msgResp := entity.MessageResponse{
			ID:        userMessage.ID,
			ChannelID: userMessage.ChannelID,
			Body:      userMessage.Content,
			Timestamp: userMessage.CreatedAt,
		}
		result = append(result, msgResp)
	}

	return result, nil
}
