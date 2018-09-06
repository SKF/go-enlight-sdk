// Package mock provides a mock for the PointAlarmStatusClient
package mock

import (
	"context"

	proto_pas "github.com/SKF/proto/pas"
	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc"

	"github.com/SKF/go-enlight-sdk/services/pas"
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

func (mock *client) SetPointAlarmThreshold(input proto_pas.SetPointAlarmThresholdInput) error {
	args := mock.Called(input)
	return args.Error(0)
}
func (mock *client) SetPointAlarmThresholdWithContext(ctx context.Context, input proto_pas.SetPointAlarmThresholdInput) error {
	args := mock.Called(ctx, input)
	return args.Error(0)
}

func (mock *client) GetPointAlarmThreshold(nodeID string) (proto_pas.GetPointAlarmThresholdOutput, error) {
	args := mock.Called(nodeID)
	return args.Get(0).(proto_pas.GetPointAlarmThresholdOutput), args.Error(1)
}
func (mock *client) GetPointAlarmThresholdWithContext(ctx context.Context, nodeID string) (proto_pas.GetPointAlarmThresholdOutput, error) {
	args := mock.Called(ctx, nodeID)
	return args.Get(0).(proto_pas.GetPointAlarmThresholdOutput), args.Error(1)
}

func (mock *client) SetPointAlarmStatus(input proto_pas.SetPointAlarmStatusInput) error {
	args := mock.Called(input)
	return args.Error(0)
}
func (mock *client) SetPointAlarmStatusWithContext(ctx context.Context, input proto_pas.SetPointAlarmStatusInput) error {
	args := mock.Called(ctx, input)
	return args.Error(0)
}

func (mock *client) GetPointAlarmStatus(input proto_pas.GetPointAlarmStatusInput) (proto_pas.AlarmStatus, error) {
	args := mock.Called(input)
	return args.Get(0).(proto_pas.AlarmStatus), args.Error(1)
}
func (mock *client) GetPointAlarmStatusWithContext(ctx context.Context, input proto_pas.GetPointAlarmStatusInput) (proto_pas.AlarmStatus, error) {
	args := mock.Called(ctx, input)
	return args.Get(0).(proto_pas.AlarmStatus), args.Error(1)
}

func (mock *client) GetPointAlarmStatusStream(dc chan<- proto_pas.GetPointAlarmStatusStreamOutput) error {
	args := mock.Called(dc)
	return args.Error(0)
}
func (mock *client) GetPointAlarmStatusStreamWithContext(ctx context.Context, dc chan<- proto_pas.GetPointAlarmStatusStreamOutput) error {
	args := mock.Called(ctx, dc)
	return args.Error(0)
}
