package mock

import (
	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc"

	"github.com/SKF/go-enlight-sdk/services/iot"
	"github.com/SKF/go-enlight-sdk/services/iot/iotgrpcapi"
)

type client struct {
	*mock.Mock
}

// Create returns an empty mock
func Create() iot.IoTClient {
	return &client{
		Mock: &mock.Mock{},
	}
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

func (mock *client) CreateTask(task iotgrpcapi.InitialTaskDescription) (string, error) {
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
func (mock *client) GetAllTasks(userID string) ([]iotgrpcapi.TaskDescription, error) {
	args := mock.Called(userID)
	return args.Get(0).([]iotgrpcapi.TaskDescription), args.Error(1)
}
func (mock *client) GetUncompletedTasks(userID string) ([]iotgrpcapi.TaskDescription, error) {
	args := mock.Called(userID)
	return args.Get(0).([]iotgrpcapi.TaskDescription), args.Error(1)
}
