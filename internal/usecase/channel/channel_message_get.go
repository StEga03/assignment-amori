package channel

import (
	"context"

	"github.com/assignment-amori/internal/constant"
	"github.com/assignment-amori/internal/entity"
)

func (u *Usecase) GetMessageInChannel(ctx context.Context, req entity.MessageUCRequest) ([]entity.MessageResponse, error) {
	var (
		result []entity.MessageResponse
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

	userMsgParam := entity.GetMessageParams{
		ChannelID: req.ChannelID,
		Limit:     constant.DefaultLimit,
		Offset:    constant.DefaultOffset,
	}
	userMessages, err := u.messageRepo.GetMessageByChannelID(ctx, userMsgParam)

	for _, userMessage := range userMessages {
		msgResp := entity.MessageResponse{
			ID:        userMessage.ID,
			ChannelID: userMessage.ChannelID,
			Type:      userMessage.SenderType,
			Body:      userMessage.Content,
			Timestamp: userMessage.CreatedAt,
		}
		result = append(result, msgResp)
	}

	return result, nil
}
