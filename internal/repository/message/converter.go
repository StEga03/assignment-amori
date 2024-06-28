package message

import (
	"github.com/assignment-amori/internal/entity"
	"github.com/assignment-amori/internal/entity/generic"
)

func (m *messageTable) ToEntity() entity.Message {
	return entity.Message{
		ID:          m.ID,
		ChannelID:   m.ChannelID,
		SenderType:  m.SenderType,
		SenderID:    m.SenderID,
		ContentType: m.ContentType,
		Content:     m.Content,
		MetaInfo: generic.MetaInfo{
			CreatedAt: m.CreatedAt,
			UpdatedAt: m.UpdatedAt,
		},
	}
}

func (m *messageInputTable) ToEntity() entity.MessageInput {
	return entity.MessageInput{
		ID:              m.ID,
		ChannelID:       m.ChannelID,
		Source:          m.Source,
		Sender:          m.Sender,
		Receiver:        m.Receiver,
		ReceiverPronoun: m.ReceiverPronoun,
		MetaInfo: generic.MetaInfo{
			CreatedAt: m.CreatedAt,
			UpdatedAt: m.UpdatedAt,
		},
	}
}

func (m *messageSourceTable) ToEntity() entity.MessageSource {
	return entity.MessageSource{
		ID:             m.ID,
		MessageInputID: m.MessageInputID,
		Sender:         m.Sender,
		ContentType:    m.ContentType,
		Content:        m.Content,
		SentAt:         m.SentAt,
		MetaInfo: generic.MetaInfo{
			CreatedAt: m.CreatedAt,
			UpdatedAt: m.UpdatedAt,
		},
	}
}
