// Based on Kero package timeutil.
package timeutils

import (
	"reflect"
	"testing"
	"time"

	"github.com/assignment-amori/internal/constant"
	"github.com/stretchr/testify/assert"
)

func TestNewTimeLocation(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "success",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := NewTimeLocation(); (err != nil) != tt.wantErr {
				t.Errorf("NewTimeLocation() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetTimeLocation(t *testing.T) {
	err := NewTimeLocation()
	if err != nil {
		t.Errorf("error init NewTimeLocation: %v", err)
	}
	tests := []struct {
		name string
		want *TimeLocations
	}{
		{
			name: "success",
			want: tl,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetTimeLocation(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetTimeLocation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConvertUTCtoWIB(t *testing.T) {
	validTime, _ := time.Parse(constant.CustomRFC3339, "2021-01-20T04:00:01+07:00")
	invalidTime, _ := time.Parse(constant.CustomRFC3339, "2021-01-00")
	tests := []struct {
		name    string
		wantErr bool
		time    time.Time
	}{
		{
			name:    "valid time",
			wantErr: false,
			time:    validTime,
		},
		{
			name:    "invalid time",
			wantErr: false,
			time:    invalidTime,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			NewTimeLocation()
			got, err := ConvertUTCtoWIB(tt.time)
			if tt.wantErr {
				assert.Nil(t, got)
				assert.NotNil(t, err)
				return
			}
			assert.Nil(t, err)
			assert.NotNil(t, got)
		})
	}
}
