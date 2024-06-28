package channel

import (
	"context"
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/assignment-amori/internal/entity"
	"github.com/assignment-amori/internal/entity/generic"
	timeutils "github.com/assignment-amori/pkg/time_utils"
	"go.uber.org/mock/gomock"
)

func TestRepository_GetByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()
	dbMock := NewMockdatabaseResource(ctrl)
	sfMock := NewMocksonyFlakeResource(ctrl)

	// Mock TimeNow.
	timeutils.Now = func() time.Time {
		return time.Date(2023, 03, 12, 0, 0, 0, 0, time.UTC)
	}

	channelTableResult := channelTable{
		ID:        1,
		UserID:    1,
		Name:      "Test",
		CreatedAt: timeutils.Now(),
		UpdatedAt: timeutils.Now(),
	}

	type args struct {
		ctx context.Context
		id  uint64
	}
	tests := []struct {
		name    string
		args    args
		mock    func()
		want    entity.Channel
		wantErr bool
	}{
		{
			name: "[fail] get channel by id",
			args: args{
				ctx: ctx,
				id:  1,
			},
			mock: func() {
				dbMock.EXPECT().Select(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(false, errors.New("error"))
			},
			want:    entity.Channel{},
			wantErr: true,
		},
		{
			name: "[fail] empty channel",
			args: args{
				ctx: ctx,
				id:  1,
			},
			mock: func() {
				dbMock.EXPECT().Select(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).SetArg(1, channelTable{})
			},
			want:    entity.Channel{},
			wantErr: true,
		},
		{
			name: "[success] get channel by id",
			args: args{
				ctx: ctx,
				id:  1,
			},
			mock: func() {
				dbMock.EXPECT().Select(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).SetArg(1, channelTableResult)
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
			got, err := r.GetByID(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetByID() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRepository_GetByUserID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()
	dbMock := NewMockdatabaseResource(ctrl)
	sfMock := NewMocksonyFlakeResource(ctrl)

	// Mock TimeNow.
	timeutils.Now = func() time.Time {
		return time.Date(2023, 03, 12, 0, 0, 0, 0, time.UTC)
	}

	channelTables := []*channelTable{
		{
			ID:        1,
			UserID:    1,
			Name:      "Test",
			CreatedAt: timeutils.Now(),
			UpdatedAt: timeutils.Now(),
		},
	}

	type args struct {
		ctx    context.Context
		userId uint64
	}
	tests := []struct {
		name    string
		args    args
		mock    func()
		want    []entity.Channel
		wantErr bool
	}{
		{
			name: "[fail] get channel by user id",
			args: args{
				ctx:    ctx,
				userId: 1,
			},
			mock: func() {
				dbMock.EXPECT().Select(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(false, errors.New("error"))
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "[success] get channel by user id",
			args: args{
				ctx:    ctx,
				userId: 1,
			},
			mock: func() {
				dbMock.EXPECT().Select(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).SetArg(1, channelTables)
			},
			want: []entity.Channel{
				{
					ID:     1,
					UserID: 1,
					Name:   "Test",
					MetaInfo: generic.MetaInfo{
						CreatedAt: timeutils.Now(),
						UpdatedAt: timeutils.Now(),
					},
				},
			},
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
			got, err := r.GetByUserID(tt.args.ctx, tt.args.userId)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetByUserID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetByUserID() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRepository_GetByIDAndUserID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()
	dbMock := NewMockdatabaseResource(ctrl)
	sfMock := NewMocksonyFlakeResource(ctrl)

	// Mock TimeNow.
	timeutils.Now = func() time.Time {
		return time.Date(2023, 03, 12, 0, 0, 0, 0, time.UTC)
	}

	channelTableResult := channelTable{
		ID:        1,
		UserID:    1,
		Name:      "Test",
		CreatedAt: timeutils.Now(),
		UpdatedAt: timeutils.Now(),
	}

	type args struct {
		ctx    context.Context
		id     uint64
		userId uint64
	}
	tests := []struct {
		name    string
		args    args
		mock    func()
		want    entity.Channel
		wantErr bool
	}{
		{
			name: "[fail] get channel by id and user id",
			args: args{
				ctx:    ctx,
				id:     1,
				userId: 1,
			},
			mock: func() {
				dbMock.EXPECT().Select(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(false, errors.New("error"))
			},
			want:    entity.Channel{},
			wantErr: true,
		},
		{
			name: "[fail] empty channel by id and user id",
			args: args{
				ctx:    ctx,
				id:     1,
				userId: 1,
			},
			mock: func() {
				dbMock.EXPECT().Select(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).SetArg(1, channelTable{})
			},
			want:    entity.Channel{},
			wantErr: true,
		},
		{
			name: "[success] get channel by id and user id",
			args: args{
				ctx:    ctx,
				id:     1,
				userId: 1,
			},
			mock: func() {
				dbMock.EXPECT().Select(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).SetArg(1, channelTableResult)
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
			got, err := r.GetByIDAndUserID(tt.args.ctx, tt.args.id, tt.args.userId)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetByIDAndUserID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetByIDAndUserID() got = %v, want %v", got, tt.want)
			}
		})
	}
}
