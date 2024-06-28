package http

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/assignment-amori/pkg/validator"
	"go.uber.org/mock/gomock"
)

func TestHandler_UploadMessageSource(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockFileUC := NewMockFileUC(ctrl)

	w := httptest.NewRecorder()
	validator.New()

	tests := []struct {
		name    string
		request func() *http.Request
		mock    func()
		want    interface{}
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Handler{
				fileUC: mockFileUC,
			}

			if tt.mock != nil {
				tt.mock()
			}

			got, err := h.UploadMessageSource(w, tt.request())
			if (err != nil) != tt.wantErr {
				t.Errorf("UploadMessageSource() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UploadMessageSource() got = %v, want %v", got, tt.want)
			}
		})
	}
}
