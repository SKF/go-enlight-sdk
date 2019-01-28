// Package mock provides a mock for the PointAlarmStatusClient
package mock

import (
	"context"

	pas_api "github.com/SKF/proto/pas"
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
func (mock *client) DialWithContext(ctx context.Context, host, port string, opts ...grpc.DialOption) error {
	args := mock.Called(ctx, host, port, opts)
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

func (mock *client) SetPointAlarmThreshold(input pas_api.SetPointAlarmThresholdInput) error {
	args := mock.Called(input)
	return args.Error(0)
}
func (mock *client) SetPointAlarmThresholdWithContext(ctx context.Context, input pas_api.SetPointAlarmThresholdInput) error {
	args := mock.Called(ctx, input)
	return args.Error(0)
}

func (mock *client) GetPointAlarmThreshold(nodeID string) (pas_api.GetPointAlarmThresholdOutput, error) {
	args := mock.Called(nodeID)
	return args.Get(0).(pas_api.GetPointAlarmThresholdOutput), args.Error(1)
}
func (mock *client) GetPointAlarmThresholdWithContext(ctx context.Context, nodeID string) (pas_api.GetPointAlarmThresholdOutput, error) {
	args := mock.Called(ctx, nodeID)
	return args.Get(0).(pas_api.GetPointAlarmThresholdOutput), args.Error(1)
}

func (mock *client) SetPointAlarmStatus(input pas_api.SetPointAlarmStatusInput) error {
	args := mock.Called(input)
	return args.Error(0)
}
func (mock *client) SetPointAlarmStatusWithContext(ctx context.Context, input pas_api.SetPointAlarmStatusInput) error {
	args := mock.Called(ctx, input)
	return args.Error(0)
}

func (mock *client) GetPointAlarmStatus(input pas_api.GetPointAlarmStatusInput) (pas_api.AlarmStatus, error) {
	args := mock.Called(input)
	return args.Get(0).(pas_api.AlarmStatus), args.Error(1)
}
func (mock *client) GetPointAlarmStatusWithContext(ctx context.Context, input pas_api.GetPointAlarmStatusInput) (pas_api.AlarmStatus, error) {
	args := mock.Called(ctx, input)
	return args.Get(0).(pas_api.AlarmStatus), args.Error(1)
}

func (mock *client) GetPointAlarmStatusEventLog(seqID string) (events pas_api.GetPointAlarmStatusEventLogOutput, err error) {
	args := mock.Called(seqID)
	return args.Get(0).(pas_api.GetPointAlarmStatusEventLogOutput), args.Error(1)
}

func (mock *client) GetPointAlarmStatusEventLogWithContext(ctx context.Context, seqID string) (events pas_api.GetPointAlarmStatusEventLogOutput, err error) {
	args := mock.Called(ctx, seqID)
	return args.Get(0).(pas_api.GetPointAlarmStatusEventLogOutput), args.Error(1)
}

func (mock *client) GetPointAlarmStatusStream(dc chan<- pas_api.GetPointAlarmStatusStreamOutput) error {
	args := mock.Called(dc)
	return args.Error(0)
}
func (mock *client) GetPointAlarmStatusStreamWithContext(ctx context.Context, dc chan<- pas_api.GetPointAlarmStatusStreamOutput) error {
	args := mock.Called(ctx, dc)
	return args.Error(0)
}

// CalculateAndSetPointAlarmStatus calculates and sets new PAS based on input data
func (mock *client) CalculateAndSetPointAlarmStatus(input pas_api.CalculateAndSetPointAlarmStatusInput) error {
	args := mock.Called(input)
	return args.Error(0)
}

// CalculateAndSetPointAlarmStatusWithContext calculates and sets new PAS based on input data
func (mock *client) CalculateAndSetPointAlarmStatusWithContext(ctx context.Context, input pas_api.CalculateAndSetPointAlarmStatusInput) error {
	args := mock.Called(ctx, input)
	return args.Error(0)
}
