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

func TestHandler_CreateMessageInChannel(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockChannelUC := NewMockChannelUsecase(ctrl)

	w := httptest.NewRecorder()
	validator.New()

	tests := []struct {
		name    string
		request func() *http.Request
		mock    func()
		want    interface{}
		wantErr bool
	}{
		{
			name: "[fail] missing param",
			request: func() *http.Request {
				ctx := chi.NewRouteContext()

				req := httptest.NewRequest(http.MethodPost, "http://localhost:9000/api/v1/channels/{channelId}/messages", nil)
				req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, ctx))

				return req
			},
			mock:    func() {},
			want:    nil,
			wantErr: true,
		},
		{
			name: "[fail] parse uint",
			request: func() *http.Request {
				ctx := chi.NewRouteContext()
				ctx.URLParams.Add(constant.ParamChannelID, "abc123")

				req := httptest.NewRequest(http.MethodPost, "http://localhost:9000/api/v1/channels/{channelId}/messages", nil)
				req.Header.Set(constant.HTTPHeaderContentType, constant.HTTPContentTypeJSON)
				req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, ctx))

				return req
			},
			mock:    func() {},
			want:    nil,
			wantErr: true,
		},
		{
			name: "[fail] called create message in channel handler",
			request: func() *http.Request {
				param, _ := json.Marshal(entity.MessageUCRequest{
					Body: "How are you?",
				})
				body := bytes.NewBuffer(param)

				ctx := chi.NewRouteContext()
				ctx.URLParams.Add(constant.ParamChannelID, "520159024281288705")

				req := httptest.NewRequest(http.MethodPost, "http://localhost:9000/api/v1/channels/{channelId}/messages", body)
				req.Header.Set(constant.HTTPHeaderContentType, constant.HTTPContentTypeJSON)
				req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, ctx))

				return req
			},
			mock: func() {
				mockChannelUC.EXPECT().CreateMessageInChannel(gomock.Any(), gomock.Any()).Return(entity.MessageResponse{}, errors.New("error"))
			},
			want:    entity.MessageResponse{},
			wantErr: true,
		},
		{
			name: "[success] create message in channel handler",
			request: func() *http.Request {
				param, _ := json.Marshal(entity.MessageUCRequest{
					Body: "How are you?",
				})
				body := bytes.NewBuffer(param)

				ctx := chi.NewRouteContext()
				ctx.URLParams.Add(constant.ParamChannelID, "520159024281288705")

				req := httptest.NewRequest(http.MethodPost, "http://localhost:9000/api/v1/channels/{channelId}/messages", body)
				req.Header.Set(constant.HTTPHeaderContentType, constant.HTTPContentTypeJSON)
				req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, ctx))

				return req
			},
			mock: func() {
				mockChannelUC.EXPECT().CreateMessageInChannel(gomock.Any(), gomock.Any()).Return(entity.MessageResponse{}, nil)
			},
			want:    entity.MessageResponse{},
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

			got, err := h.CreateMessageInChannel(w, tt.request())
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateMessageInChannel() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateMessageInChannel() got = %v, want %v", got, tt.want)
			}
		})
	}
}
