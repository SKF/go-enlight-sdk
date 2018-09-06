package mock

import (
	"context"

	proto_iot "github.com/SKF/proto/iot"
	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc"

	"github.com/SKF/go-enlight-sdk/services/iot"
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

func (mock *client) CreateTask(task proto_iot.InitialTaskDescription) (string, error) {
	args := mock.Called(task)
	return args.String(0), args.Error(1)
}
func (mock *client) CreateTaskWithContext(ctx context.Context, task proto_iot.InitialTaskDescription) (string, error) {
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

func (mock *client) GetAllTasks(userID string) ([]proto_iot.TaskDescription, error) {
	args := mock.Called(userID)
	return args.Get(0).([]proto_iot.TaskDescription), args.Error(1)
}
func (mock *client) GetAllTasksWithContext(ctx context.Context, userID string) ([]proto_iot.TaskDescription, error) {
	args := mock.Called(ctx, userID)
	return args.Get(0).([]proto_iot.TaskDescription), args.Error(1)
}

func (mock *client) GetUncompletedTasks(userID string) ([]proto_iot.TaskDescription, error) {
	args := mock.Called(userID)
	return args.Get(0).([]proto_iot.TaskDescription), args.Error(1)
}
func (mock *client) GetUncompletedTasksWithContext(ctx context.Context, userID string) ([]proto_iot.TaskDescription, error) {
	args := mock.Called(ctx, userID)
	return args.Get(0).([]proto_iot.TaskDescription), args.Error(1)
}

func (mock *client) GetUncompletedTasksByHierarchy(nodeID string) ([]proto_iot.TaskDescription, error) {
	args := mock.Called(nodeID)
	return args.Get(0).([]proto_iot.TaskDescription), args.Error(1)
}
func (mock *client) GetUncompletedTasksByHierarchyWithContext(ctx context.Context, nodeID string) ([]proto_iot.TaskDescription, error) {
	args := mock.Called(ctx, nodeID)
	return args.Get(0).([]proto_iot.TaskDescription), args.Error(1)
}

func (mock *client) SetTaskStatus(input proto_iot.SetTaskStatusInput) error {
	args := mock.Called(input)
	return args.Error(0)
}
func (mock *client) SetTaskStatusWithContext(ctx context.Context, input proto_iot.SetTaskStatusInput) error {
	args := mock.Called(ctx, input)
	return args.Error(0)
}

func (mock *client) GetTaskStream(input proto_iot.GetTaskStreamInput, dc chan<- proto_iot.GetTaskStreamOutput) error {
	args := mock.Called(input, dc)
	return args.Error(0)
}
func (mock *client) GetTaskStreamWithContext(ctx context.Context, input proto_iot.GetTaskStreamInput, dc chan<- proto_iot.GetTaskStreamOutput) error {
	args := mock.Called(ctx, input, dc)
	return args.Error(0)
}

func (mock *client) IngestNodeData(input proto_iot.IngestNodeDataInput) error {
	args := mock.Called(input)
	return args.Error(0)
}
func (mock *client) IngestNodeDataWithContext(ctx context.Context, input proto_iot.IngestNodeDataInput) error {
	args := mock.Called(ctx, input)
	return args.Error(0)
}

func (mock *client) IngestNodeDataStream(c <-chan proto_iot.IngestNodeDataStreamInput) error {
	args := mock.Called(c)
	return args.Error(0)
}
func (mock *client) IngestNodeDataStreamWithContext(ctx context.Context, c <-chan proto_iot.IngestNodeDataStreamInput) error {
	args := mock.Called(ctx, c)
	return args.Error(0)
}

func (mock *client) GetLatestNodeData(input proto_iot.GetLatestNodeDataInput) (*proto_iot.NodeData, error) {
	args := mock.Called(input)
	return args.Get(0).(*proto_iot.NodeData), args.Error(1)
}

func (mock *client) GetLatestNodeDataWithContext(ctx context.Context, input proto_iot.GetLatestNodeDataInput) (*proto_iot.NodeData, error) {
	args := mock.Called(ctx, input)
	return args.Get(0).(*proto_iot.NodeData), args.Error(1)
}

func (mock *client) GetNodeData(input proto_iot.GetNodeDataInput) ([]proto_iot.NodeData, error) {
	args := mock.Called(input)
	return args.Get(0).([]proto_iot.NodeData), args.Error(1)
}
func (mock *client) GetNodeDataWithContext(ctx context.Context, input proto_iot.GetNodeDataInput) ([]proto_iot.NodeData, error) {
	args := mock.Called(ctx, input)
	return args.Get(0).([]proto_iot.NodeData), args.Error(1)
}

func (mock *client) GetNodeDataStream(input proto_iot.GetNodeDataStreamInput, c chan<- proto_iot.GetNodeDataStreamOutput) error {
	args := mock.Called(input, c)
	return args.Error(0)
}
func (mock *client) GetNodeDataStreamWithContext(ctx context.Context, input proto_iot.GetNodeDataStreamInput, c chan<- proto_iot.GetNodeDataStreamOutput) error {
	args := mock.Called(ctx, input, c)
	return args.Error(0)
}

func (mock *client) GetMedia(input proto_iot.GetMediaInput) (proto_iot.Media, error) {
	args := mock.Called(input)
	return args.Get(0).(proto_iot.Media), args.Error(1)
}
func (mock *client) GetMediaWithContext(ctx context.Context, input proto_iot.GetMediaInput) (media proto_iot.Media, err error) {
	args := mock.Called(ctx, input)
	return args.Get(0).(proto_iot.Media), args.Error(1)
}

func (mock *client) GetTasksByStatus(input proto_iot.GetTasksByStatusInput) ([]*proto_iot.TaskDescription, error) {
	args := mock.Called(input)
	return args.Get(0).([]*proto_iot.TaskDescription), args.Error(1)
}
func (mock *client) GetTasksByStatusWithContext(ctx context.Context, input proto_iot.GetTasksByStatusInput) ([]*proto_iot.TaskDescription, error) {
	args := mock.Called(ctx, input)
	return args.Get(0).([]*proto_iot.TaskDescription), args.Error(1)
}

func (mock *client) GetTaskByUUID(input string) (output *proto_iot.TaskDescription, err error) {
	args := mock.Called(input)
	return args.Get(0).(*proto_iot.TaskDescription), args.Error(1)
}
func (mock *client) GetTaskByUUIDWithContext(ctx context.Context, input string) (output *proto_iot.TaskDescription, err error) {
	args := mock.Called(ctx, input)
	return args.Get(0).(*proto_iot.TaskDescription), args.Error(1)
}
func (mock *client) GetTaskByLongId(input int64) (output *proto_iot.TaskDescription, err error) {
	args := mock.Called(input)
	return args.Get(0).(*proto_iot.TaskDescription), args.Error(1)
}
func (mock *client) GetTaskByLongIdWithContext(ctx context.Context, input int64) (output *proto_iot.TaskDescription, err error) {
	args := mock.Called(ctx, input)
	return args.Get(0).(*proto_iot.TaskDescription), args.Error(1)
}
