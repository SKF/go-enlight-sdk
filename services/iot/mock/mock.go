package mock

import (
	"context"

	iot_grpcapi "github.com/SKF/proto/iot"
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

func (mock *client) CreateTask(task iot_grpcapi.InitialTaskDescription) (string, error) {
	args := mock.Called(task)
	return args.String(0), args.Error(1)
}
func (mock *client) CreateTaskWithContext(ctx context.Context, task iot_grpcapi.InitialTaskDescription) (string, error) {
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

func (mock *client) GetAllTasks(userID string) ([]iot_grpcapi.TaskDescription, error) {
	args := mock.Called(userID)
	return args.Get(0).([]iot_grpcapi.TaskDescription), args.Error(1)
}
func (mock *client) GetAllTasksWithContext(ctx context.Context, userID string) ([]iot_grpcapi.TaskDescription, error) {
	args := mock.Called(ctx, userID)
	return args.Get(0).([]iot_grpcapi.TaskDescription), args.Error(1)
}

func (mock *client) GetUncompletedTasks(userID string) ([]iot_grpcapi.TaskDescription, error) {
	args := mock.Called(userID)
	return args.Get(0).([]iot_grpcapi.TaskDescription), args.Error(1)
}
func (mock *client) GetUncompletedTasksWithContext(ctx context.Context, userID string) ([]iot_grpcapi.TaskDescription, error) {
	args := mock.Called(ctx, userID)
	return args.Get(0).([]iot_grpcapi.TaskDescription), args.Error(1)
}

func (mock *client) GetUncompletedTasksByHierarchy(nodeID string) ([]iot_grpcapi.TaskDescription, error) {
	args := mock.Called(nodeID)
	return args.Get(0).([]iot_grpcapi.TaskDescription), args.Error(1)
}
func (mock *client) GetUncompletedTasksByHierarchyWithContext(ctx context.Context, nodeID string) ([]iot_grpcapi.TaskDescription, error) {
	args := mock.Called(ctx, nodeID)
	return args.Get(0).([]iot_grpcapi.TaskDescription), args.Error(1)
}

func (mock *client) SetTaskStatus(input iot_grpcapi.SetTaskStatusInput) error {
	args := mock.Called(input)
	return args.Error(0)
}
func (mock *client) SetTaskStatusWithContext(ctx context.Context, input iot_grpcapi.SetTaskStatusInput) error {
	args := mock.Called(ctx, input)
	return args.Error(0)
}

func (mock *client) GetTaskStream(input iot_grpcapi.GetTaskStreamInput, dc chan<- iot_grpcapi.GetTaskStreamOutput) error {
	args := mock.Called(input, dc)
	return args.Error(0)
}
func (mock *client) GetTaskStreamWithContext(ctx context.Context, input iot_grpcapi.GetTaskStreamInput, dc chan<- iot_grpcapi.GetTaskStreamOutput) error {
	args := mock.Called(ctx, input, dc)
	return args.Error(0)
}

func (mock *client) IngestNodeData(input iot_grpcapi.IngestNodeDataInput) error {
	args := mock.Called(input)
	return args.Error(0)
}
func (mock *client) IngestNodeDataWithContext(ctx context.Context, input iot_grpcapi.IngestNodeDataInput) error {
	args := mock.Called(ctx, input)
	return args.Error(0)
}

func (mock *client) IngestNodeDataStream(c <-chan iot_grpcapi.IngestNodeDataStreamInput) error {
	args := mock.Called(c)
	return args.Error(0)
}
func (mock *client) IngestNodeDataStreamWithContext(ctx context.Context, c <-chan iot_grpcapi.IngestNodeDataStreamInput) error {
	args := mock.Called(ctx, c)
	return args.Error(0)
}

func (mock *client) GetLatestNodeData(input iot_grpcapi.GetLatestNodeDataInput) (*iot_grpcapi.NodeData, error) {
	args := mock.Called(input)
	return args.Get(0).(*iot_grpcapi.NodeData), args.Error(1)
}

func (mock *client) GetLatestNodeDataWithContext(ctx context.Context, input iot_grpcapi.GetLatestNodeDataInput) (*iot_grpcapi.NodeData, error) {
	args := mock.Called(ctx, input)
	return args.Get(0).(*iot_grpcapi.NodeData), args.Error(1)
}

func (mock *client) GetNodeData(input iot_grpcapi.GetNodeDataInput) ([]iot_grpcapi.NodeData, error) {
	args := mock.Called(input)
	return args.Get(0).([]iot_grpcapi.NodeData), args.Error(1)
}
func (mock *client) GetNodeDataWithContext(ctx context.Context, input iot_grpcapi.GetNodeDataInput) ([]iot_grpcapi.NodeData, error) {
	args := mock.Called(ctx, input)
	return args.Get(0).([]iot_grpcapi.NodeData), args.Error(1)
}

func (mock *client) GetNodeDataStream(input iot_grpcapi.GetNodeDataStreamInput, c chan<- iot_grpcapi.GetNodeDataStreamOutput) error {
	args := mock.Called(input, c)
	return args.Error(0)
}
func (mock *client) GetNodeDataStreamWithContext(ctx context.Context, input iot_grpcapi.GetNodeDataStreamInput, c chan<- iot_grpcapi.GetNodeDataStreamOutput) error {
	args := mock.Called(ctx, input, c)
	return args.Error(0)
}

func (mock *client) GetMedia(input iot_grpcapi.GetMediaInput) (iot_grpcapi.Media, error) {
	args := mock.Called(input)
	return args.Get(0).(iot_grpcapi.Media), args.Error(1)
}
func (mock *client) GetMediaWithContext(ctx context.Context, input iot_grpcapi.GetMediaInput) (media iot_grpcapi.Media, err error) {
	args := mock.Called(ctx, input)
	return args.Get(0).(iot_grpcapi.Media), args.Error(1)
}

func (mock *client) GetTasksByStatus(input iot_grpcapi.GetTasksByStatusInput) ([]*iot_grpcapi.TaskDescription, error) {
	args := mock.Called(input)
	return args.Get(0).([]*iot_grpcapi.TaskDescription), args.Error(1)
}
func (mock *client) GetTasksByStatusWithContext(ctx context.Context, input iot_grpcapi.GetTasksByStatusInput) ([]*iot_grpcapi.TaskDescription, error) {
	args := mock.Called(ctx, input)
	return args.Get(0).([]*iot_grpcapi.TaskDescription), args.Error(1)
}

func (mock *client) GetTaskByUUID(input string) (output *iot_grpcapi.TaskDescription, err error) {
	args := mock.Called(input)
	return args.Get(0).(*iot_grpcapi.TaskDescription), args.Error(1)
}
func (mock *client) GetTaskByUUIDWithContext(ctx context.Context, input string) (output *iot_grpcapi.TaskDescription, err error) {
	args := mock.Called(ctx, input)
	return args.Get(0).(*iot_grpcapi.TaskDescription), args.Error(1)
}
func (mock *client) GetTaskByLongId(input int64) (output *iot_grpcapi.TaskDescription, err error) {
	args := mock.Called(input)
	return args.Get(0).(*iot_grpcapi.TaskDescription), args.Error(1)
}
func (mock *client) GetTaskByLongIdWithContext(ctx context.Context, input int64) (output *iot_grpcapi.TaskDescription, err error) {
	args := mock.Called(ctx, input)
	return args.Get(0).(*iot_grpcapi.TaskDescription), args.Error(1)
}
func (mock *client) DeleteNodeData(input iot_grpcapi.DeleteNodeDataInput) error {
	args := mock.Called(input)
	return args.Error(0)
}
func (mock *client) DeleteNodeDataWithContext(ctx context.Context, input iot_grpcapi.DeleteNodeDataInput) error {
	args := mock.Called(ctx, input)
	return args.Error(0)
}

func (mock *client) GetTasksModifiedSinceTimestamp(input iot_grpcapi.GetTasksModifiedSinceTimestampInput) (*iot_grpcapi.GetTasksModifiedSinceTimestampOutput, error) {
	args := mock.Called(input)
	return args.Get(0).(*iot_grpcapi.GetTasksModifiedSinceTimestampOutput), args.Error(1)
}
func (mock *client) GetTasksModifiedSinceTimestampWithContext(ctx context.Context, input iot_grpcapi.GetTasksModifiedSinceTimestampInput) (output *iot_grpcapi.GetTasksModifiedSinceTimestampOutput, err error) {
	args := mock.Called(ctx, input)
	return args.Get(0).(*iot_grpcapi.GetTasksModifiedSinceTimestampOutput), args.Error(1)
}

func (mock *client) GetNodeDataLog(input iot_grpcapi.GetNodeDataLogInput) (output *iot_grpcapi.GetNodeDataLogOutput, err error) {
	args := mock.Called(input)
	return args.Get(0).(*iot_grpcapi.GetNodeDataLogOutput), args.Error(1)
}
func (mock *client) GetNodeDataLogWithContext(ctx context.Context, input iot_grpcapi.GetNodeDataLogInput) (output *iot_grpcapi.GetNodeDataLogOutput, err error) {
	args := mock.Called(ctx, input)
	return args.Get(0).(*iot_grpcapi.GetNodeDataLogOutput), args.Error(1)
}
