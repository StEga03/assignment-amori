package channel

import (
	"context"
	"errors"
	"testing"

	"github.com/assignment-amori/internal/constant"
	"github.com/assignment-amori/internal/entity"
	"github.com/assignment-amori/pkg/testfiles"
	"go.uber.org/mock/gomock"
)

func TestUsecase_CreateChannel(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()
	mockConsistencyResource := NewMockconsistencyResource(ctrl)
	mockChannelResource := NewMockchannelResource(ctrl)
	mockMessageResource := NewMockmessageResource(ctrl)
	mockOpenaiResource := NewMockopenaiResource(ctrl)
	mockUserResource := NewMockuserResource(ctrl)
	mockSonyflakeResource := NewMocksonyFlakeResource(ctrl)

	req := entity.NewChannelUCRequest{
		Name: "Test",
		MessageSource: []entity.MessageSourceUCRequest{
			{
				Body:   "Hey, how was your day?",
				Sender: "Alice",
				SentAt: "2024-06-28T08:00:00",
			},
			{
				Body:   "It was good, thanks for asking! How about yours?",
				Sender: "Bob",
				SentAt: "2024-06-28T08:05:00",
			},
		},
	}

	type args struct {
		ctx context.Context
		req entity.NewChannelUCRequest
	}
	tests := []struct {
		name    string
		args    args
		mock    func()
		want    uint64
		wantErr bool
	}{
		{
			name: "[fail] error generate channel id",
			args: args{
				ctx: ctx,
				req: req,
			},
			mock: func() {
				mockUserResource.EXPECT().GetUserByContext(gomock.Any()).Return(entity.User{}, nil)
				mockSonyflakeResource.EXPECT().NextID().Return(constant.DefaultUInt64, errors.New("error"))
			},
			want:    constant.DefaultUInt64,
			wantErr: true,
		},
		{
			name: "[fail] error generate message input id",
			args: args{
				ctx: ctx,
				req: req,
			},
			mock: func() {
				mockUserResource.EXPECT().GetUserByContext(gomock.Any()).Return(entity.User{}, nil)
				mockSonyflakeResource.EXPECT().NextID().Return(uint64(123), nil)
				mockSonyflakeResource.EXPECT().NextID().Return(constant.DefaultUInt64, errors.New("error"))
			},
			want:    constant.DefaultUInt64,
			wantErr: true,
		},
		{
			name: "[fail] error run as unit",
			args: args{
				ctx: ctx,
				req: req,
			},
			mock: func() {
				mockUserResource.EXPECT().GetUserByContext(gomock.Any()).Return(entity.User{}, nil)
				mockSonyflakeResource.EXPECT().NextID().Return(uint64(123), nil)
				mockSonyflakeResource.EXPECT().NextID().Return(uint64(321), nil)
				mockConsistencyResource.EXPECT().RunAsUnit(gomock.Any(), gomock.Any()).Return(errors.New("error"))
			},
			want:    constant.DefaultUInt64,
			wantErr: true,
		},
		{
			name: "[fail] error create channel",
			args: args{
				ctx: ctx,
				req: req,
			},
			mock: func() {
				mockUserResource.EXPECT().GetUserByContext(gomock.Any()).Return(entity.User{}, nil)
				mockSonyflakeResource.EXPECT().NextID().Return(uint64(123), nil)
				mockSonyflakeResource.EXPECT().NextID().Return(uint64(321), nil)
				mockConsistencyResource.EXPECT().RunAsUnit(gomock.Any(), gomock.Any()).DoAndReturn(testfiles.DoRunAsUnit)
				mockChannelResource.EXPECT().CreateChannel(gomock.Any(), gomock.Any(), gomock.Any()).Return(constant.DefaultUInt64, errors.New("error"))
			},
			want:    constant.DefaultUInt64,
			wantErr: true,
		},
		{
			name: "[fail] error create message input",
			args: args{
				ctx: ctx,
				req: req,
			},
			mock: func() {
				mockUserResource.EXPECT().GetUserByContext(gomock.Any()).Return(entity.User{}, nil)
				mockSonyflakeResource.EXPECT().NextID().Return(uint64(123), nil)
				mockSonyflakeResource.EXPECT().NextID().Return(uint64(321), nil)
				mockConsistencyResource.EXPECT().RunAsUnit(gomock.Any(), gomock.Any()).DoAndReturn(testfiles.DoRunAsUnit)
				mockChannelResource.EXPECT().CreateChannel(gomock.Any(), gomock.Any(), gomock.Any()).Return(uint64(123), nil)
				mockMessageResource.EXPECT().CreateMessageInput(gomock.Any(), gomock.Any(), gomock.Any()).Return(constant.DefaultUInt64, errors.New("error"))
			},
			want:    constant.DefaultUInt64,
			wantErr: true,
		},
		{
			name: "[fail] error create message source",
			args: args{
				ctx: ctx,
				req: req,
			},
			mock: func() {
				mockUserResource.EXPECT().GetUserByContext(gomock.Any()).Return(entity.User{}, nil)
				mockSonyflakeResource.EXPECT().NextID().Return(uint64(123), nil)
				mockSonyflakeResource.EXPECT().NextID().Return(uint64(321), nil)
				mockConsistencyResource.EXPECT().RunAsUnit(gomock.Any(), gomock.Any()).DoAndReturn(testfiles.DoRunAsUnit)
				mockChannelResource.EXPECT().CreateChannel(gomock.Any(), gomock.Any(), gomock.Any()).Return(uint64(123), nil)
				mockMessageResource.EXPECT().CreateMessageInput(gomock.Any(), gomock.Any(), gomock.Any()).Return(uint64(321), nil)
				mockMessageResource.EXPECT().CreateMessageSource(gomock.Any(), gomock.Any(), gomock.Any()).Return(errors.New("error"))
			},
			want:    constant.DefaultUInt64,
			wantErr: true,
		},
		{
			name: "[success] create channel",
			args: args{
				ctx: ctx,
				req: req,
			},
			mock: func() {
				mockUserResource.EXPECT().GetUserByContext(gomock.Any()).Return(entity.User{}, nil)
				mockSonyflakeResource.EXPECT().NextID().Return(uint64(123), nil)
				mockSonyflakeResource.EXPECT().NextID().Return(uint64(321), nil)
				mockConsistencyResource.EXPECT().RunAsUnit(gomock.Any(), gomock.Any()).DoAndReturn(testfiles.DoRunAsUnit)
				mockChannelResource.EXPECT().CreateChannel(gomock.Any(), gomock.Any(), gomock.Any()).Return(uint64(123), nil)
				mockMessageResource.EXPECT().CreateMessageInput(gomock.Any(), gomock.Any(), gomock.Any()).Return(uint64(321), nil)
				mockMessageResource.EXPECT().CreateMessageSource(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
			},
			want:    uint64(123),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &Usecase{
				consistency: mockConsistencyResource,
				channelRepo: mockChannelResource,
				messageRepo: mockMessageResource,
				openaiRepo:  mockOpenaiResource,
				userRepo:    mockUserResource,
				sf:          mockSonyflakeResource,
			}

			if tt.mock != nil {
				tt.mock()
			}

			got, err := u.CreateChannel(ctx, tt.args.req)
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
