package channel

import (
	"github.com/assignment-amori/internal/entity"
	"github.com/assignment-amori/internal/entity/generic"
)

func (c *channelTable) ToEntity() entity.Channel {
	return entity.Channel{
		ID:     c.ID,
		UserID: c.UserID,
		Name:   c.Name,
		MetaInfo: generic.MetaInfo{
			CreatedAt: c.CreatedAt,
			UpdatedAt: c.UpdatedAt,
		},
	}
}
