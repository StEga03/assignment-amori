package utils

import (
	"testing"
	"time"

	"github.com/assignment-amori/internal/constant"
)

func Test_ParseStringToTime(t *testing.T) {
	timeStr := "2021-01-23T15:04:05"
	timeParsed, _ := time.Parse(constant.CustomDateFormat, timeStr)
	type args struct {
		input string
	}
	tests := []struct {
		name    string
		args    args
		want    time.Time
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				input: "2021-01-23T15:04:05",
			},
			want:    timeParsed,
			wantErr: false,
		},
		{
			name: "empty input",
			args: args{
				input: "",
			},
			want:    time.Time{},
			wantErr: true,
		},
		{
			name: "invalid input",
			args: args{
				input: "abc-def-ghi",
			},
			want:    time.Time{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, err := ParseStringToTime(tt.args.input, constant.CustomDateFormat); (err != nil) != tt.wantErr {
				t.Errorf("util.ParseStringToTime() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
