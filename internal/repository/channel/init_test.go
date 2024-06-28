package channel

import (
	"reflect"
	"testing"

	"go.uber.org/mock/gomock"
)

func TestNew(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	dbMock := NewMockdatabaseResource(ctrl)
	sfMock := NewMocksonyFlakeResource(ctrl)

	type args struct {
		db databaseResource
		sf sonyFlakeResource
	}
	tests := []struct {
		name string
		args args
		want *Repository
	}{
		{
			name: "[success] init",
			args: args{
				db: dbMock,
				sf: sfMock,
			},
			want: &Repository{
				db: dbMock,
				sf: sfMock,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.db, tt.args.sf); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}
