package entity

import (
	"time"

	"github.com/assignment-amori/internal/entity/generic"
)

type Message struct {
	ID          uint64 `json:"id"`
	ChannelID   uint64 `json:"channelId"`
	SenderType  string `json:"senderType"`
	SenderID    uint64 `json:"senderId"`
	ContentType string `json:"contentType"`
	Content     string `json:"content"`
	generic.MetaInfo
}

type MessageInput struct {
	ID              uint64 `json:"id"`
	ChannelID       uint64 `json:"channelId"`
	Source          string `json:"source"`
	Sender          string `json:"sender"`
	Receiver        string `json:"receiver"`
	ReceiverPronoun string `json:"receiverPronoun"`
	generic.MetaInfo
}

type MessageSource struct {
	ID             uint64    `json:"id"`
	MessageInputID uint64    `json:"messageInputId"`
	Sender         string    `json:"sender"`
	ContentType    string    `json:"contentType"`
	Content        string    `json:"content"`
	SentAt         time.Time `json:"sentAt"`
	generic.MetaInfo
}

type MessageUCRequest struct {
	ChannelID uint64 `json:"channelId" validate:"required"`
	Body      string `json:"body"`
}

type MessageSourceUCRequest struct {
	Body   string `json:"body"`
	Sender string `json:"sender"`
	SentAt string `json:"sentAt"`
}

type MessageResponse struct {
	ID        uint64    `json:"id"`
	ChannelID uint64    `json:"channelId"`
	Body      string    `json:"body"`
	Timestamp time.Time `json:"timestamp"`
}

type NewMessageParams struct {
	ID          uint64 `json:"id"`
	ChannelID   uint64 `json:"channelId"`
	SenderType  string `json:"senderType"`
	SenderID    uint64 `json:"senderId"`
	ContentType string `json:"contentType"`
	Content     string `json:"content"`
	generic.MetaInfo
}

type NewMessageInputParams struct {
	ID              uint64 `json:"id"`
	ChannelID       uint64 `json:"channelId"`
	Source          string `json:"source"`
	Sender          string `json:"sender"`
	Receiver        string `json:"receiver"`
	ReceiverPronoun string `json:"receiverPronoun"`
}

type NewMessageSourceParams struct {
	MessageInputID uint64    `json:"messageInputId"`
	Sender         string    `json:"sender"`
	ContentType    string    `json:"contentType"`
	Content        string    `json:"content"`
	SentAt         time.Time `json:"sentAt"`
}

type GetMessageParams struct {
	ID        uint64 `json:"id"`
	ChannelID uint64 `json:"channelId"`
	Limit     int    `json:"limit"`
	Offset    int    `json:"offset"`
}

type GetMessageInputParams struct {
	ID        uint64 `json:"id"`
	ChannelID uint64 `json:"channelId"`
	Limit     int    `json:"limit"`
	Offset    int    `json:"offset"`
}

type GetMessageSourceParams struct {
	ID             uint64 `json:"id"`
	MessageInputID uint64 `json:"messageInputId"`
	Limit          int    `json:"limit"`
	Offset         int    `json:"offset"`
}
