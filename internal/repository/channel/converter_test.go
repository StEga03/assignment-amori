package channel

import (
	"reflect"
	"testing"
	"time"

	"github.com/assignment-amori/internal/entity"
	"github.com/assignment-amori/internal/entity/generic"
	timeutils "github.com/assignment-amori/pkg/time_utils"
)

func Test_channelTable_ToEntity(t *testing.T) {
	// Mock TimeNow.
	timeutils.Now = func() time.Time {
		return time.Date(2023, 03, 12, 0, 0, 0, 0, time.UTC)
	}

	type fields struct {
		ID        uint64
		UserID    uint64
		Name      string
		CreatedAt time.Time
		UpdatedAt time.Time
	}
	tests := []struct {
		name   string
		fields fields
		want   entity.Channel
	}{
		{
			name: "[success] convert to entity",
			fields: fields{
				ID:        1,
				UserID:    1,
				Name:      "Test",
				CreatedAt: timeutils.Now(),
				UpdatedAt: timeutils.Now(),
			},
			want: entity.Channel{
				ID:     1,
				UserID: 1,
				Name:   "Test",
				MetaInfo: generic.MetaInfo{
					CreatedAt: timeutils.Now(),
					UpdatedAt: timeutils.Now(),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &channelTable{
				ID:        tt.fields.ID,
				UserID:    tt.fields.UserID,
				Name:      tt.fields.Name,
				CreatedAt: tt.fields.CreatedAt,
				UpdatedAt: tt.fields.UpdatedAt,
			}
			if got := c.ToEntity(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToEntity() = %v, want %v", got, tt.want)
			}
		})
	}
}
