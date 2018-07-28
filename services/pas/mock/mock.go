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

func (mock *client) SetPointAlarmThresholds(input pasapi.SetPointAlarmThresholdsInput) error {
	args := mock.Called(input)
	return args.Error(0)
}
func (mock *client) SetPointAlarmThresholdsWithContext(ctx context.Context, input pasapi.SetPointAlarmThresholdsInput) error {
	args := mock.Called(ctx, input)
	return args.Error(0)
}

func (mock *client) GetPointAlarmThresholds(nodeID string) ([]pasapi.AlarmStatusInterval, error) {
	args := mock.Called(nodeID)
	return args.Get(0).([]pasapi.AlarmStatusInterval), args.Error(1)
}
func (mock *client) GetPointAlarmThresholdsWithContext(ctx context.Context, nodeID string) (intervals []pasapi.AlarmStatusInterval, err error) {
	args := mock.Called(ctx, nodeID)
	return args.Get(0).([]pasapi.AlarmStatusInterval), args.Error(1)
}

func (mock *client) SetPointAlarmStatus(input pasapi.SetPointAlarmStatusInput) error {
	args := mock.Called(input)
	return args.Error(0)
}
func (mock *client) SetPointAlarmStatusWithContext(ctx context.Context, input pasapi.SetPointAlarmStatusInput) error {
	args := mock.Called(ctx, input)
	return args.Error(0)
}

func (mock *client) GetPointAlarmStatus(input pasapi.GetPointAlarmStatusInput) (pasapi.AlarmStatus, error) {
	args := mock.Called(input)
	return args.Get(0).(pasapi.AlarmStatus), args.Error(1)
}
func (mock *client) GetPointAlarmStatusWithContext(ctx context.Context, input pasapi.GetPointAlarmStatusInput) (pasapi.AlarmStatus, error) {
	args := mock.Called(ctx, input)
	return args.Get(0).(pasapi.AlarmStatus), args.Error(1)
}

func (mock *client) GetPointAlarmStatusStream(dc chan<- pasapi.GetPointAlarmStatusStreamOutput) error {
	args := mock.Called(dc)
	return args.Error(0)
}
func (mock *client) GetPointAlarmStatusStreamWithContext(ctx context.Context, dc chan<- pasapi.GetPointAlarmStatusStreamOutput) error {
	args := mock.Called(ctx, dc)
	return args.Error(0)
}
