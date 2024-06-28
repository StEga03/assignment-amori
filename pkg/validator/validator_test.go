package validator

import (
	"context"
	"testing"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name: "success",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := New(); (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidateStruct(t *testing.T) {
	New()
	type args struct {
		ctx context.Context
		s   interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				ctx: context.TODO(),
				s: struct {
					Test string `validate:"required"`
				}{
					Test: "ABC",
				},
			},
			wantErr: false,
		},
		{
			name: "failed",
			args: args{
				ctx: context.TODO(),
				s: struct {
					Test string `validate:"required"`
				}{},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidateStruct(tt.args.ctx, tt.args.s); (err != nil) != tt.wantErr {
				t.Errorf("ValidateStruct() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
