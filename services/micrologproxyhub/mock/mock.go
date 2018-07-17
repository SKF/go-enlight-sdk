package mock

import (
	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc"

	"github.com/SKF/go-enlight-sdk/services/iot"
	api "github.com/SKF/go-enlight-sdk/services/iot/iotgrpcapi"
	"github.com/SKF/go-utility/uuid"
)

type client struct {
	mock.Mock
	iot.IoTClient
}

// Create returns an empty mock
func Create() *client {
	return new(client)
}

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

func (mock *client) SetTaskStatus(taskID, userID uuid.UUID, status api.TaskStatus) error {
	args := mock.Called(taskID, userID, status)
	return args.Error(0)
}

func (mock *client) GetTasksStream(dc chan<- api.GetTasksStreamOutput) error {
	args := mock.Called(dc)
	return args.Error(0)
}
