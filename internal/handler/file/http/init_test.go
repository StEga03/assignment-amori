package http

import (
	"reflect"
	"testing"

	"go.uber.org/mock/gomock"
)

func TestNew(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockChannelUC := NewMockFileUC(ctrl)

	type args struct {
		fileUC FileUC
	}
	tests := []struct {
		name string
		args args
		want *Handler
	}{
		{
			name: "Success",
			args: args{
				fileUC: mockChannelUC,
			},
			want: &Handler{
				fileUC: mockChannelUC,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.fileUC); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}
