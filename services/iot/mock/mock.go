package mock

import (
	"context"

	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc"

	"github.com/SKF/go-enlight-sdk/services/iot"
	api "github.com/SKF/go-enlight-sdk/services/iot/iotgrpcapi"
)

type client struct {
	mock.Mock
}

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
func (mock *client) DeepPingWithContext(ctx context.Context) error {
	args := mock.Called(ctx)
	return args.Error(0)
}

func (mock *client) CreateTask(task api.InitialTaskDescription) (string, error) {
	args := mock.Called(task)
	return args.String(0), args.Error(1)
}
func (mock *client) CreateTaskWithContext(ctx context.Context, task api.InitialTaskDescription) (string, error) {
	args := mock.Called(ctx, task)
	return args.String(0), args.Error(1)
}

func (mock *client) DeleteTask(userID, taskID string) error {
	args := mock.Called(userID, taskID)
	return args.Error(0)
}
func (mock *client) DeleteTaskWithContext(ctx context.Context, userID, taskID string) error {
	args := mock.Called(ctx, userID, taskID)
	return args.Error(0)
}

func (mock *client) SetTaskCompleted(userID, taskID string) error {
	args := mock.Called(userID, taskID)
	return args.Error(0)
}
func (mock *client) SetTaskCompletedWithContext(ctx context.Context, userID, taskID string) error {
	args := mock.Called(ctx, userID, taskID)
	return args.Error(0)
}

func (mock *client) GetAllTasks(userID string) ([]api.TaskDescription, error) {
	args := mock.Called(userID)
	return args.Get(0).([]api.TaskDescription), args.Error(1)
}
func (mock *client) GetAllTasksWithContext(ctx context.Context, userID string) ([]api.TaskDescription, error) {
	args := mock.Called(ctx, userID)
	return args.Get(0).([]api.TaskDescription), args.Error(1)
}

func (mock *client) GetUncompletedTasks(userID string) ([]api.TaskDescription, error) {
	args := mock.Called(userID)
	return args.Get(0).([]api.TaskDescription), args.Error(1)
}
func (mock *client) GetUncompletedTasksWithContext(ctx context.Context, userID string) ([]api.TaskDescription, error) {
	args := mock.Called(ctx, userID)
	return args.Get(0).([]api.TaskDescription), args.Error(1)
}

func (mock *client) GetUncompletedTasksByHierarchy(nodeID string) ([]api.TaskDescription, error) {
	args := mock.Called(nodeID)
	return args.Get(0).([]api.TaskDescription), args.Error(1)
}
func (mock *client) GetUncompletedTasksByHierarchyWithContext(ctx context.Context, nodeID string) ([]api.TaskDescription, error) {
	args := mock.Called(ctx, nodeID)
	return args.Get(0).([]api.TaskDescription), args.Error(1)
}

func (mock *client) SetTaskStatus(taskID, userID string, status api.TaskStatus) (err error) {
	args := mock.Called(taskID, userID, status)
	return args.Error(0)
}
func (mock *client) SetTaskStatusWithContext(ctx context.Context, taskID, userID string, status api.TaskStatus) (err error) {
	args := mock.Called(ctx, taskID, userID, status)
	return args.Error(0)
}

func (mock *client) GetTaskStream(input api.GetTaskStreamInput, dc chan<- api.GetTaskStreamOutput) (err error) {
	args := mock.Called(input, dc)
	return args.Error(0)
}
func (mock *client) GetTaskStreamWithContext(ctx context.Context, input api.GetTaskStreamInput, dc chan<- api.GetTaskStreamOutput) (err error) {
	args := mock.Called(ctx, input, dc)
	return args.Error(0)
}

func (mock *client) IngestNodeData(nodeID string, nodeData api.NodeData) error {
	args := mock.Called(nodeID, nodeData)
	return args.Error(0)
}
func (mock *client) IngestNodeDataWithContext(ctx context.Context, nodeID string, nodeData api.NodeData) error {
	args := mock.Called(ctx, nodeID, nodeData)
	return args.Error(0)
}

func (mock *client) IngestNodeDataStream(c <-chan api.IngestNodeDataStreamInput) error {
	args := mock.Called(c)
	return args.Error(0)
}
func (mock *client) IngestNodeDataStreamWithContext(ctx context.Context, c <-chan api.IngestNodeDataStreamInput) error {
	args := mock.Called(ctx, c)
	return args.Error(0)
}

func (mock *client) GetLatestNodeData(input *api.GetLatestNodeDataInput) (api.NodeData, error) {
	args := mock.Called(input)
	return args.Get(0).(api.NodeData), args.Error(1)
}

func (mock *client) GetLatestNodeDataContext(ctx context.Context, input *api.GetLatestNodeDataInput) (nodeData api.NodeData, err error) {
	args := mock.Called(ctx, input)
	return args.Get(0).(api.NodeData), args.Error(1)
}

func (mock *client) GetNodeData(input api.GetNodeDataInput) ([]api.NodeData, error) {
	args := mock.Called(input)
	return args.Get(0).([]api.NodeData), args.Error(1)
}
func (mock *client) GetNodeDataWithContext(ctx context.Context, input api.GetNodeDataInput) ([]api.NodeData, error) {
	args := mock.Called(ctx, input)
	return args.Get(0).([]api.NodeData), args.Error(1)
}

func (mock *client) GetNodeDataStream(input api.GetNodeDataStreamInput, c chan<- api.GetNodeDataStreamOutput) error {
	args := mock.Called(input, c)
	return args.Error(0)
}
func (mock *client) GetNodeDataStreamWithContext(ctx context.Context, input api.GetNodeDataStreamInput, c chan<- api.GetNodeDataStreamOutput) error {
	args := mock.Called(ctx, input, c)
	return args.Error(0)
}
