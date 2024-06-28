package http

import (
	"bytes"
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/assignment-amori/internal/constant"
	"github.com/assignment-amori/internal/entity"
	"github.com/assignment-amori/pkg/json"
	"github.com/assignment-amori/pkg/validator"
	"github.com/go-chi/chi/v5"
	"go.uber.org/mock/gomock"
)

func TestHandler_CreateChannel(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockChannelUC := NewMockChannelUsecase(ctrl)

	w := httptest.NewRecorder()
	validator.New()

	validParam := entity.NewChannelUCRequest{
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

	tests := []struct {
		name    string
		request func() *http.Request
		mock    func()
		want    interface{}
		wantErr bool
	}{
		{
			name: "[fail] error cast and validate",
			request: func() *http.Request {
				ctx := chi.NewRouteContext()

				req := httptest.NewRequest(http.MethodPost, "http://localhost:9000/api/v1/channels", nil)
				req.Header.Set(constant.HTTPHeaderContentType, constant.HTTPContentTypeJSON)
				req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, ctx))

				return req
			},
			mock:    func() {},
			want:    nil,
			wantErr: true,
		},
		{
			name: "[fail] error when call create channel from usecase",
			request: func() *http.Request {
				param, _ := json.Marshal(validParam)
				body := bytes.NewBuffer(param)

				ctx := chi.NewRouteContext()

				req := httptest.NewRequest(http.MethodPost, "http://localhost:9000/api/v1/channels", body)
				req.Header.Set(constant.HTTPHeaderContentType, constant.HTTPContentTypeJSON)
				req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, ctx))

				return req
			},
			mock: func() {
				mockChannelUC.EXPECT().CreateChannel(gomock.Any(), gomock.Any()).Return(constant.DefaultUInt64, errors.New("error"))
			},
			want:    constant.DefaultUInt64,
			wantErr: true,
		},
		{
			name: "[success] create channel handler",
			request: func() *http.Request {
				param, _ := json.Marshal(validParam)
				body := bytes.NewBuffer(param)

				ctx := chi.NewRouteContext()

				req := httptest.NewRequest(http.MethodPost, "http://localhost:9000/api/v1/channels", body)
				req.Header.Set(constant.HTTPHeaderContentType, constant.HTTPContentTypeJSON)
				req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, ctx))

				return req
			},
			mock: func() {
				mockChannelUC.EXPECT().CreateChannel(gomock.Any(), gomock.Any()).Return(uint64(1), nil)
			},
			want:    uint64(1),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Handler{
				ChannelUC: mockChannelUC,
			}

			if tt.mock != nil {
				tt.mock()
			}

			got, err := h.CreateChannel(w, tt.request())
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateChannel() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateChannel() got = %v, want %v", got, tt.want)
			}
		})
	}
}
