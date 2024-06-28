package channel

import (
	"context"
	"testing"
	"time"

	"github.com/assignment-amori/internal/entity"
	"github.com/assignment-amori/pkg/consistency"
	timeutils "github.com/assignment-amori/pkg/time_utils"
	"go.uber.org/mock/gomock"
)

func TestRepository_CreateChannel(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()
	dbMock := NewMockdatabaseResource(ctrl)
	sfMock := NewMocksonyFlakeResource(ctrl)
	celMock := consistency.ConsistencyElement{}

	// Mock TimeNow.
	timeutils.Now = func() time.Time {
		return time.Date(2023, 03, 12, 0, 0, 0, 0, time.UTC)
	}

	channelParam := entity.NewChannelParams{
		ID:     1,
		UserID: 1,
		Name:   "Test",
	}

	type args struct {
		ctx   context.Context
		param entity.NewChannelParams
		cel   *consistency.ConsistencyElement
	}
	tests := []struct {
		name    string
		args    args
		mock    func()
		want    uint64
		wantErr bool
	}{
		{
			name: "[success] create channel",
			args: args{
				ctx:   ctx,
				param: channelParam,
				cel:   &celMock,
			},
			mock: func() {
				dbMock.EXPECT().ExecuteInTx(ctx, gomock.Any(), gomock.Any()).Return(nil)
			},
			want:    1,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.mock != nil {
				tt.mock()
			}

			r := &Repository{
				db: dbMock,
				sf: sfMock,
			}
			got, err := r.CreateChannel(tt.args.ctx, tt.args.param, tt.args.cel)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateChannel() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CreateChannel() got = %v, want %v", got, tt.want)
			}
		})
	}
}
