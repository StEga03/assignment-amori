package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTimeValidator(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantErr bool
	}{
		{"success", "16-03-2022 22:52", false},
		{"success", "2019-10-12T07:20:50Z", false},
		{"success time format field cnote_pod_date JNE", "02 MAR 2022  08:26", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := TimeValidator(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("TimeValidator() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.NotNil(t, got)
		})
	}
}

func TestBetween(t *testing.T) {
	tests := []struct {
		name  string
		value string
		a     string
		b     string
		want  string
	}{
		{"t1", "foo", "f", "o", ""},
		{"t2", "foo", "x", "o", ""},
		{"t3", "foo", "f", "x", ""},
		{"t4", "fox", "f", "x", "o"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Between(tt.value, tt.a, tt.b); got != tt.want {
				t.Errorf("Between() = %v, want %v", got, tt.want)
			}
		})
	}
}
