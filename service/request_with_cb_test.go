package service

import (
	"context"
	"errors"
	"github.com/daffarg/distributed-cascading-cb/circuitbreaker"
	"github.com/daffarg/distributed-cascading-cb/config"
	"github.com/daffarg/distributed-cascading-cb/mock"
	"github.com/daffarg/distributed-cascading-cb/protobuf"
	"github.com/daffarg/distributed-cascading-cb/util"
	"github.com/go-kit/log"
	logkit "github.com/go-kit/log"
	"github.com/go-playground/validator/v10"
	"github.com/golang/mock/gomock"
	"github.com/jarcoal/httpmock"
	"go.opentelemetry.io/otel/trace"
	"net/http"
	"os"
	"reflect"
	"testing"
	"time"
)

func Test_service_requestWithCircuitBreaker(t *testing.T) {
	reqValidator := validator.New()

	var svcLog logkit.Logger
	{
		svcLog = logkit.NewJSONLogger(os.Stdout)
		svcLog = logkit.With(svcLog, util.LogTimestamp, logkit.TimestampFormat(time.Now, time.RFC3339), util.LogPath, logkit.DefaultCaller)
	}

	type fields struct {
		log          log.Logger
		validator    *validator.Validate
		breakers     map[string]*circuitbreaker.CircuitBreaker
		httpClient   *http.Client
		tracer       trace.Tracer
		config       *config.Config
		subscribeMap map[string]bool
	}
	type args struct {
		ctx       context.Context
		req       *request
		mockFunc  func(ctrl *gomock.Controller, mockRepository *mock.MockRepository, mockBroker *mock.MockMessageBroker)
		deferFunc func()
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Response
		wantErr bool
	}{
		{
			name: "Failed_parse_url",
			fields: fields{
				log:          svcLog,
				validator:    reqValidator,
				breakers:     make(map[string]*circuitbreaker.CircuitBreaker),
				subscribeMap: make(map[string]bool),
				config:       config.NewConfig(),
				httpClient:   &http.Client{},
			},
			args: args{
				ctx: context.Background(),
				req: &request{
					Method:            "GET",
					URL:               "\n\thttp://localhost:8081/hello",
					Header:            map[string]string{},
					Body:              []byte{},
					RequiringEndpoint: "http://localhost:8080/hello",
					RequiringMethod:   "GET",
				},
				mockFunc: func(ctrl *gomock.Controller, mockRepository *mock.MockRepository, mockBroker *mock.MockMessageBroker) {
				},
			},
			wantErr: true,
			want:    &Response{},
		},
		{
			name: "New_service_no_status_found",
			fields: fields{
				log:          svcLog,
				validator:    reqValidator,
				breakers:     make(map[string]*circuitbreaker.CircuitBreaker),
				subscribeMap: make(map[string]bool),
				config:       config.NewConfig(),
				httpClient:   &http.Client{},
			},
			args: args{
				ctx: context.Background(),
				req: &request{
					Method:            "GET",
					URL:               "http://localhost:8081/hello",
					Header:            map[string]string{},
					Body:              []byte{},
					RequiringEndpoint: "http://localhost:8080/hello",
					RequiringMethod:   "GET",
				},
				mockFunc: func(ctrl *gomock.Controller, mockRepository *mock.MockRepository, mockBroker *mock.MockMessageBroker) {
					httpmock.Activate()

					httpmock.RegisterResponder("GET", "http://localhost:8081/hello",
						httpmock.NewStringResponder(200, `{"status":"ok"}`))

					mockBroker.EXPECT().Subscribe(gomock.Any(), gomock.Any()).Return(nil, util.ErrUpdatedStatusNotFound)
					mockRepository.EXPECT().AddMembersIntoSet(gomock.Any(), gomock.Any(), gomock.Any()).Return(int64(1), nil).AnyTimes()
					mockRepository.EXPECT().AddMembersIntoSet(gomock.Any(), gomock.Any(), gomock.Any()).Return(int64(1), nil).AnyTimes()
					mockBroker.EXPECT().SubscribeAsync(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes()
					mockRepository.EXPECT().Get(gomock.Any(), gomock.Any()).Return("", util.ErrKeyNotFound)
				},
			},
			wantErr: false,
			want: &Response{
				Body:                      []byte(`{"status":"ok"}`),
				Status:                    "200",
				StatusCode:                http.StatusOK,
				Proto:                     "",
				ProtoMajor:                0,
				ProtoMinor:                0,
				Header:                    make(map[string]string),
				ContentLength:             -1,
				IsFromAlternativeEndpoint: false,
			},
		},
		{
			name: "New_service_subscribe_failed",
			fields: fields{
				log:          svcLog,
				validator:    reqValidator,
				breakers:     make(map[string]*circuitbreaker.CircuitBreaker),
				subscribeMap: make(map[string]bool),
				config:       config.NewConfig(),
				httpClient:   &http.Client{},
			},
			args: args{
				ctx: context.Background(),
				req: &request{
					Method:            "GET",
					URL:               "http://localhost:8081/hello",
					Header:            map[string]string{},
					Body:              []byte{},
					RequiringEndpoint: "http://localhost:8080/hello",
					RequiringMethod:   "GET",
				},
				mockFunc: func(ctrl *gomock.Controller, mockRepository *mock.MockRepository, mockBroker *mock.MockMessageBroker) {
					httpmock.Activate()

					httpmock.RegisterResponder("GET", "http://localhost:8081/hello",
						httpmock.NewStringResponder(200, `{"status":"ok"}`))

					mockBroker.EXPECT().Subscribe(gomock.Any(), gomock.Any()).Return(nil, errors.New("failed to subscribe"))
					mockRepository.EXPECT().AddMembersIntoSet(gomock.Any(), gomock.Any(), gomock.Any()).Return(int64(1), nil).AnyTimes()
					mockRepository.EXPECT().AddMembersIntoSet(gomock.Any(), gomock.Any(), gomock.Any()).Return(int64(1), nil).AnyTimes()
					mockBroker.EXPECT().SubscribeAsync(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes()
					mockRepository.EXPECT().Get(gomock.Any(), gomock.Any()).Return("", util.ErrKeyNotFound)
				},
			},
			wantErr: false,
			want: &Response{
				Body:                      []byte(`{"status":"ok"}`),
				Status:                    "200",
				StatusCode:                http.StatusOK,
				Proto:                     "",
				ProtoMajor:                0,
				ProtoMinor:                0,
				Header:                    make(map[string]string),
				ContentLength:             -1,
				IsFromAlternativeEndpoint: false,
			},
		},
		{
			name: "New_service_status_open",
			fields: fields{
				log:          svcLog,
				validator:    reqValidator,
				breakers:     make(map[string]*circuitbreaker.CircuitBreaker),
				subscribeMap: make(map[string]bool),
				config:       config.NewConfig(),
				httpClient:   &http.Client{},
			},
			args: args{
				ctx: context.Background(),
				req: &request{
					Method:            "GET",
					URL:               "http://localhost:8081/hello",
					Header:            map[string]string{},
					Body:              []byte{},
					RequiringEndpoint: "http://localhost:8080/hello",
					RequiringMethod:   "GET",
				},
				mockFunc: func(ctrl *gomock.Controller, mockRepository *mock.MockRepository, mockBroker *mock.MockMessageBroker) {
					httpmock.Activate()

					httpmock.RegisterResponder("GET", "http://localhost:8081/hello",
						httpmock.NewStringResponder(200, `{"status":"ok"}`))

					mockBroker.EXPECT().Subscribe(gomock.Any(), gomock.Any()).Return(
						&protobuf.Status{
							Status:    circuitbreaker.StateOpen.String(),
							Endpoint:  "GET:localhost:8081/hello",
							Timeout:   10000,
							Timestamp: time.Now().Add(-1 * time.Minute).Format(time.RFC3339),
						},
						nil,
					)
					mockRepository.EXPECT().AddMembersIntoSet(gomock.Any(), gomock.Any(), gomock.Any()).Return(int64(1), nil).AnyTimes()
					mockRepository.EXPECT().AddMembersIntoSet(gomock.Any(), gomock.Any(), gomock.Any()).Return(int64(1), nil).AnyTimes()
					mockRepository.EXPECT().SetWithExp(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(nil).AnyTimes()
					mockBroker.EXPECT().SubscribeAsync(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes()
				},
			},
			wantErr: true,
			want:    &Response{},
		},
		{
			name: "Status_close",
			fields: fields{
				log:       svcLog,
				validator: reqValidator,
				breakers:  make(map[string]*circuitbreaker.CircuitBreaker),
				subscribeMap: map[string]bool{
					"GET:localhost:8081/hello": true,
				},
				config:     config.NewConfig(),
				httpClient: &http.Client{},
			},
			args: args{
				ctx: context.Background(),
				req: &request{
					Method:            "GET",
					URL:               "http://localhost:8081/hello",
					Header:            map[string]string{},
					Body:              []byte{},
					RequiringEndpoint: "http://localhost:8080/hello",
					RequiringMethod:   "GET",
				},
				mockFunc: func(ctrl *gomock.Controller, mockRepository *mock.MockRepository, mockBroker *mock.MockMessageBroker) {
					httpmock.Activate()

					httpmock.RegisterResponder("GET", "http://localhost:8081/hello",
						httpmock.NewStringResponder(200, `{"status":"ok"}`))

					mockRepository.EXPECT().AddMembersIntoSet(gomock.Any(), gomock.Any(), gomock.Any()).Return(int64(0), nil).AnyTimes()
					mockRepository.EXPECT().AddMembersIntoSet(gomock.Any(), gomock.Any(), gomock.Any()).Return(int64(0), nil).AnyTimes()
					mockRepository.EXPECT().Get(gomock.Any(), gomock.Any()).Return("", util.ErrKeyNotFound)
				},
			},
			wantErr: false,
			want: &Response{
				Body:                      []byte(`{"status":"ok"}`),
				Status:                    "200",
				StatusCode:                http.StatusOK,
				Proto:                     "",
				ProtoMajor:                0,
				ProtoMinor:                0,
				Header:                    make(map[string]string),
				ContentLength:             -1,
				IsFromAlternativeEndpoint: false,
			},
		},
		{
			name: "Failed_to_get_status",
			fields: fields{
				log:       svcLog,
				validator: reqValidator,
				breakers:  make(map[string]*circuitbreaker.CircuitBreaker),
				subscribeMap: map[string]bool{
					"GET:localhost:8081/hello": true,
				},
				config:     config.NewConfig(),
				httpClient: &http.Client{},
			},
			args: args{
				ctx: context.Background(),
				req: &request{
					Method:            "GET",
					URL:               "http://localhost:8081/hello",
					Header:            map[string]string{},
					Body:              []byte{},
					RequiringEndpoint: "http://localhost:8080/hello",
					RequiringMethod:   "GET",
				},
				mockFunc: func(ctrl *gomock.Controller, mockRepository *mock.MockRepository, mockBroker *mock.MockMessageBroker) {
					httpmock.Activate()

					httpmock.RegisterResponder("GET", "http://localhost:8081/hello",
						httpmock.NewStringResponder(200, `{"status":"ok"}`))

					mockRepository.EXPECT().AddMembersIntoSet(gomock.Any(), gomock.Any(), gomock.Any()).Return(int64(0), nil).AnyTimes()
					mockRepository.EXPECT().AddMembersIntoSet(gomock.Any(), gomock.Any(), gomock.Any()).Return(int64(0), nil).AnyTimes()
					mockRepository.EXPECT().Get(gomock.Any(), gomock.Any()).Return("", errors.New("failed to get status"))
				},
			},
			wantErr: false,
			want: &Response{
				Body:                      []byte(`{"status":"ok"}`),
				Status:                    "200",
				StatusCode:                http.StatusOK,
				Proto:                     "",
				ProtoMajor:                0,
				ProtoMinor:                0,
				Header:                    make(map[string]string),
				ContentLength:             -1,
				IsFromAlternativeEndpoint: false,
			},
		},
		{
			name: "Status_open_no_alt_and_not_exception",
			fields: fields{
				log:       svcLog,
				validator: reqValidator,
				breakers:  make(map[string]*circuitbreaker.CircuitBreaker),
				subscribeMap: map[string]bool{
					"GET:localhost:8081/hello": true,
				},
				config:     config.NewConfig(),
				httpClient: &http.Client{},
			},
			args: args{
				ctx: context.Background(),
				req: &request{
					Method:            "GET",
					URL:               "http://localhost:8081/hello",
					Header:            map[string]string{},
					Body:              []byte{},
					RequiringEndpoint: "http://localhost:8080/hello",
					RequiringMethod:   "GET",
				},
				mockFunc: func(ctrl *gomock.Controller, mockRepository *mock.MockRepository, mockBroker *mock.MockMessageBroker) {
					mockRepository.EXPECT().AddMembersIntoSet(gomock.Any(), gomock.Any(), gomock.Any()).Return(int64(0), nil).AnyTimes()
					mockRepository.EXPECT().AddMembersIntoSet(gomock.Any(), gomock.Any(), gomock.Any()).Return(int64(0), nil).AnyTimes()
					mockRepository.EXPECT().Get(gomock.Any(), gomock.Any()).Return("open", nil)
				},
			},
			wantErr: true,
			want:    &Response{},
		},
		{
			name: "Status_open_available_alt",
			fields: fields{
				log:       svcLog,
				validator: reqValidator,
				breakers:  make(map[string]*circuitbreaker.CircuitBreaker),
				subscribeMap: map[string]bool{
					"GET:localhost:8081/hello": true,
				},
				config: &config.Config{
					AlternativeEndpoints: map[string]config.AlternativeEndpoint{
						"GET:localhost:8081/hello": {
							Endpoint: "http://localhost:8081/hello",
							Method:   "GET",
							Alternatives: []config.Endpoint{
								{
									Endpoint: "http://localhost:8082/hello",
									Method:   "GET",
								},
							},
						},
					},
				},
				httpClient: &http.Client{},
			},
			args: args{
				ctx: context.Background(),
				req: &request{
					Method:            "GET",
					URL:               "http://localhost:8081/hello",
					Header:            map[string]string{},
					Body:              []byte{},
					RequiringEndpoint: "http://localhost:8080/hello",
					RequiringMethod:   "GET",
				},
				mockFunc: func(ctrl *gomock.Controller, mockRepository *mock.MockRepository, mockBroker *mock.MockMessageBroker) {
					httpmock.Activate()

					httpmock.RegisterResponder("GET", "http://localhost:8082/hello",
						httpmock.NewStringResponder(200, `{"status":"ok"}`))

					mockRepository.EXPECT().AddMembersIntoSet(gomock.Any(), gomock.Any(), gomock.Any()).Return(int64(0), nil).AnyTimes()
					mockRepository.EXPECT().AddMembersIntoSet(gomock.Any(), gomock.Any(), gomock.Any()).Return(int64(0), nil).AnyTimes()
					mockRepository.EXPECT().Get(gomock.Any(), gomock.Any()).Return("open", nil)
					mockRepository.EXPECT().Get(gomock.Any(), gomock.Any()).Return("", util.ErrKeyNotFound).AnyTimes()
				},
			},
			wantErr: false,
			want: &Response{
				Body:                      []byte(`{"status":"ok"}`),
				Status:                    "200",
				StatusCode:                http.StatusOK,
				Proto:                     "",
				ProtoMajor:                0,
				ProtoMinor:                0,
				Header:                    make(map[string]string),
				ContentLength:             -1,
				IsFromAlternativeEndpoint: true,
			},
		},
		{
			name: "Status_open_exception",
			fields: fields{
				log:       svcLog,
				validator: reqValidator,
				breakers:  make(map[string]*circuitbreaker.CircuitBreaker),
				subscribeMap: map[string]bool{
					"GET:localhost:8081/hello": true,
				},
				config: &config.Config{
					Exceptions: map[string]config.Endpoint{
						"GET:localhost:8081/hello": {
							Endpoint: "http://localhost:8081/hello",
							Method:   "GET",
						},
					},
				},
				httpClient: &http.Client{},
			},
			args: args{
				ctx: context.Background(),
				req: &request{
					Method:            "GET",
					URL:               "http://localhost:8081/hello",
					Header:            map[string]string{},
					Body:              []byte{},
					RequiringEndpoint: "http://localhost:8080/hello",
					RequiringMethod:   "GET",
				},
				mockFunc: func(ctrl *gomock.Controller, mockRepository *mock.MockRepository, mockBroker *mock.MockMessageBroker) {
					httpmock.Activate()

					httpmock.RegisterResponder("GET", "http://localhost:8081/hello",
						httpmock.NewStringResponder(200, `{"status":"ok"}`))

					mockRepository.EXPECT().AddMembersIntoSet(gomock.Any(), gomock.Any(), gomock.Any()).Return(int64(0), nil).AnyTimes()
					mockRepository.EXPECT().AddMembersIntoSet(gomock.Any(), gomock.Any(), gomock.Any()).Return(int64(0), nil).AnyTimes()
					mockRepository.EXPECT().Get(gomock.Any(), gomock.Any()).Return("open", nil)
				},
			},
			wantErr: false,
			want: &Response{
				Body:                      []byte(`{"status":"ok"}`),
				Status:                    "200",
				StatusCode:                http.StatusOK,
				Proto:                     "",
				ProtoMajor:                0,
				ProtoMinor:                0,
				Header:                    make(map[string]string),
				ContentLength:             -1,
				IsFromAlternativeEndpoint: false,
			},
		},
		{
			name: "Status_closed_failed_execute",
			fields: fields{
				log:       svcLog,
				validator: reqValidator,
				breakers:  make(map[string]*circuitbreaker.CircuitBreaker),
				subscribeMap: map[string]bool{
					"GET:localhost:8081/hello": true,
				},
				config: &config.Config{
					Exceptions: map[string]config.Endpoint{
						"GET:localhost:8081/hello": {
							Endpoint: "http://localhost:8081/hello",
							Method:   "GET",
						},
					},
				},
				httpClient: &http.Client{},
			},
			args: args{
				ctx: context.Background(),
				req: &request{
					Method:            "GET",
					URL:               "http://localhost:8081/hello",
					Header:            map[string]string{},
					Body:              []byte{},
					RequiringEndpoint: "http://localhost:8080/hello",
					RequiringMethod:   "GET",
				},
				mockFunc: func(ctrl *gomock.Controller, mockRepository *mock.MockRepository, mockBroker *mock.MockMessageBroker) {
					httpmock.Activate()

					httpmock.RegisterResponder("GET", "http://localhost:8081/hello",
						httpmock.NewStringResponder(500, `{"status":"failed"}`))

					mockRepository.EXPECT().AddMembersIntoSet(gomock.Any(), gomock.Any(), gomock.Any()).Return(int64(0), nil).AnyTimes()
					mockRepository.EXPECT().AddMembersIntoSet(gomock.Any(), gomock.Any(), gomock.Any()).Return(int64(0), nil).AnyTimes()
					mockRepository.EXPECT().Get(gomock.Any(), gomock.Any()).Return("", util.ErrKeyNotFound)
				},
			},
			wantErr: true,
			want:    &Response{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockRepository := mock.NewMockRepository(ctrl)
			mockBroker := mock.NewMockMessageBroker(ctrl)

			s := &service{
				log:          tt.fields.log,
				validator:    tt.fields.validator,
				repository:   mockRepository,
				broker:       mockBroker,
				breakers:     tt.fields.breakers,
				httpClient:   tt.fields.httpClient,
				tracer:       tt.fields.tracer,
				config:       tt.fields.config,
				subscribeMap: tt.fields.subscribeMap,
			}
			tt.args.mockFunc(ctrl, mockRepository, mockBroker)
			got, err := s.requestWithCircuitBreaker(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("requestWithCircuitBreaker() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("requestWithCircuitBreaker() got = %v, want %v", got, tt.want)
			}
		})
	}
}
