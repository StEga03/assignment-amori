package locale

import "testing"

func TestNewLocale(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{name: "success", wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := NewLocale(); (err != nil) != tt.wantErr {
				t.Errorf("NewLocale() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
