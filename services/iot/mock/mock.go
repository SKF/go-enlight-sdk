package mock

import (
	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc"

	"github.com/SKF/go-enlight-sdk/services/iot"
	api "github.com/SKF/go-enlight-sdk/services/iot/iot_grpc_api"
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

func (mock *client) CreateTask(task api.InitialTaskDescription) (string, error) {
	args := mock.Called(task)
	return args.String(0), args.Error(1)
}
func (mock *client) DeleteTask(userID, taskID string) error {
	args := mock.Called(userID, taskID)
	return args.Error(0)
}
func (mock *client) SetTaskCompleted(userID, taskID string) error {
	args := mock.Called(userID, taskID)
	return args.Error(0)
}
func (mock *client) GetAllTasks(userID string) ([]api.TaskDescription, error) {
	args := mock.Called(userID)
	return args.Get(0).([]api.TaskDescription), args.Error(1)
}
func (mock *client) GetUncompletedTasks(userID string) ([]api.TaskDescription, error) {
	args := mock.Called(userID)
	return args.Get(0).([]api.TaskDescription), args.Error(1)
}
