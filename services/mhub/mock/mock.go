package mock

import (
	"context"

	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc"

	"github.com/SKF/go-enlight-sdk/services/mhub"
	"github.com/SKF/go-enlight-sdk/services/mhub/mhubapi"
	"github.com/SKF/go-utility/uuid"
)

type client struct {
	mock.Mock
}

func Create() *client {
	return new(client)
}

var _ mhub.MicrologProxyHubClient = &client{}

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

func (mock *client) SetTaskStatus(taskID, userID uuid.UUID, status mhubapi.TaskStatus) error {
	args := mock.Called(taskID, userID, status)
	return args.Error(0)
}

func (mock *client) SetTaskStatusWithContext(ctx context.Context, taskID, userID uuid.UUID, status mhubapi.TaskStatus) error {
	args := mock.Called(ctx, taskID, userID, status)
	return args.Error(0)
}

func (mock *client) AvailableDSKFStream(dc chan<- mhubapi.AvailableDSKFStreamOutput) error {
	args := mock.Called(dc)
	return args.Error(0)
}

func (mock *client) AvailableDSKFStreamWithContext(ctx context.Context, dc chan<- mhubapi.AvailableDSKFStreamOutput) error {
	args := mock.Called(ctx, dc)
	return args.Error(0)
}
