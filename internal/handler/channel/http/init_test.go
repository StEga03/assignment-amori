package http

import (
	"reflect"
	"testing"

	"go.uber.org/mock/gomock"
)

func TestNew(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockChannelUC := NewMockChannelUsecase(ctrl)

	type args struct {
		ChannelUC ChannelUsecase
	}
	tests := []struct {
		name string
		args args
		want *Handler
	}{
		{
			name: "Success",
			args: args{
				ChannelUC: mockChannelUC,
			},
			want: &Handler{
				ChannelUC: mockChannelUC,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.ChannelUC); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}
