package mock

import (
	"context"

	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc"

	"github.com/SKF/go-enlight-sdk/services/pas"
	"github.com/SKF/go-enlight-sdk/services/pas/pasapi"
)

type client struct {
	mock.Mock
}

func Create() *client {
	return new(client)
}

var _ pas.PointAlarmStatusClient = &client{}

func (mock *client) Dial(host, port string, opts ...grpc.DialOption) error {
	args := mock.Called(host, port, opts)
	return args.Error(0)
}

func (mock *client) Close() {
	mock.Called()
	return
}

func (mock *client) DeepPing() error {
	args := mock.Called()
	return args.Error(0)
}
func (mock *client) DeepPingWithContext(ctx context.Context) error {
	args := mock.Called(ctx)
	return args.Error(0)
}

func (mock *client) SetPointThreshold(input pasapi.SetPointThresholdInput) error {
	args := mock.Called(input)
	return args.Error(0)
}
func (mock *client) SetPointThresholdWithContext(ctx context.Context, input pasapi.SetPointThresholdInput) error {
	args := mock.Called(ctx, input)
	return args.Error(0)
}

func (mock *client) GetPointThreshold(nodeID string) ([]pasapi.AlarmStatusInterval, error) {
	args := mock.Called(nodeID)
	return args.Get(0).([]pasapi.AlarmStatusInterval), args.Error(1)
}
func (mock *client) GetPointThresholdWithContext(ctx context.Context, nodeID string) (intervals []pasapi.AlarmStatusInterval, err error) {
	args := mock.Called(ctx, nodeID)
	return args.Get(0).([]pasapi.AlarmStatusInterval), args.Error(1)
}

func (mock *client) SetPointStatus(input pasapi.SetPointStatusInput) error {
	args := mock.Called(input)
	return args.Error(0)
}
func (mock *client) SetPointStatusWithContext(ctx context.Context, input pasapi.SetPointStatusInput) error {
	args := mock.Called(ctx, input)
	return args.Error(0)
}

func (mock *client) GetPointStatus(input pasapi.GetPointStatusInput) (pasapi.AlarmStatus, error) {
	args := mock.Called(input)
	return args.Get(0).(pasapi.AlarmStatus), args.Error(1)
}
func (mock *client) GetPointStatusWithContext(ctx context.Context, input pasapi.GetPointStatusInput) (pasapi.AlarmStatus, error) {
	args := mock.Called(ctx, input)
	return args.Get(0).(pasapi.AlarmStatus), args.Error(1)
}

func (mock *client) GetPointStatusStream(dc chan<- pasapi.GetPointStatusStreamOutput) error {
	args := mock.Called(dc)
	return args.Error(0)
}
func (mock *client) GetPointStatusStreamWithContext(ctx context.Context, dc chan<- pasapi.GetPointStatusStreamOutput) error {
	args := mock.Called(ctx, dc)
	return args.Error(0)
}
