// Code generated by MockGen. DO NOT EDIT.
// Source: broker/broker.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"
	time "time"

	protobuf "github.com/daffarg/distributed-cascading-cb/protobuf"
	gomock "github.com/golang/mock/gomock"
)

// MockMessageBroker is a mock of MessageBroker interface.
type MockMessageBroker struct {
	ctrl     *gomock.Controller
	recorder *MockMessageBrokerMockRecorder
}

// MockMessageBrokerMockRecorder is the mock recorder for MockMessageBroker.
type MockMessageBrokerMockRecorder struct {
	mock *MockMessageBroker
}

// NewMockMessageBroker creates a new mock instance.
func NewMockMessageBroker(ctrl *gomock.Controller) *MockMessageBroker {
	mock := &MockMessageBroker{ctrl: ctrl}
	mock.recorder = &MockMessageBrokerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMessageBroker) EXPECT() *MockMessageBrokerMockRecorder {
	return m.recorder
}

// Publish mocks base method.
func (m *MockMessageBroker) Publish(ctx context.Context, topic string, message *protobuf.Status) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Publish", ctx, topic, message)
	ret0, _ := ret[0].(error)
	return ret0
}

// Publish indicates an expected call of Publish.
func (mr *MockMessageBrokerMockRecorder) Publish(ctx, topic, message interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Publish", reflect.TypeOf((*MockMessageBroker)(nil).Publish), ctx, topic, message)
}

// Subscribe mocks base method.
func (m *MockMessageBroker) Subscribe(ctx context.Context, topic string) (*protobuf.Status, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Subscribe", ctx, topic)
	ret0, _ := ret[0].(*protobuf.Status)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Subscribe indicates an expected call of Subscribe.
func (mr *MockMessageBrokerMockRecorder) Subscribe(ctx, topic interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Subscribe", reflect.TypeOf((*MockMessageBroker)(nil).Subscribe), ctx, topic)
}

// SubscribeAsync mocks base method.
func (m *MockMessageBroker) SubscribeAsync(ctx context.Context, topic string, handler func(context.Context, string, string, time.Duration) error) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SubscribeAsync", ctx, topic, handler)
}

// SubscribeAsync indicates an expected call of SubscribeAsync.
func (mr *MockMessageBrokerMockRecorder) SubscribeAsync(ctx, topic, handler interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubscribeAsync", reflect.TypeOf((*MockMessageBroker)(nil).SubscribeAsync), ctx, topic, handler)
}