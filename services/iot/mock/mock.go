package mock

import (
	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc"

	"github.com/SKF/go-enlight-sdk/services/iot"
	api "github.com/SKF/go-enlight-sdk/services/iot/iotgrpcapi"
)

type client struct {
	mock.Mock
}

// Create returns an empty mock
func Create() *client {
	return new(client)
}

var _ iot.IoTClient = &client{}

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

func (mock *client) GetUncompletedTasksByHierarchy(nodeID string) ([]api.TaskDescription, error) {
	args := mock.Called(nodeID)
	return args.Get(0).([]api.TaskDescription), args.Error(1)
}

func (mock *client) SetTaskStatus(taskID, userID string, status api.TaskStatus) (err error) {
	args := mock.Called(taskID, userID, status)
	return args.Error(0)
}
func (mock *client) GetTaskStream(input api.GetTaskStreamInput, dc chan<- api.GetTaskStreamOutput) (err error) {
	args := mock.Called(input, dc)
	return args.Error(0)
}

func (mock *client) IngestNodeData(nodeID string, nodeData api.NodeData) error {
	args := mock.Called(nodeID, nodeData)
	return args.Error(0)
}
func (mock *client) GetNodeData(input api.GetNodeDataInput) ([]api.NodeData, error) {
	args := mock.Called(input)
	return nil, args.Error(0)
}
func (mock *client) GetNodeDataStream(input api.GetNodeDataStreamInput, c chan<- api.GetNodeDataStreamOutput) error {
	args := mock.Called(input, c)
	return args.Error(0)
}
