package iot

import (
	"context"
	"time"

	"google.golang.org/grpc"

	api "github.com/SKF/go-enlight-sdk/services/iot/iotgrpcapi"
)

type IoTClient interface {
	Dial(host, port string, opts ...grpc.DialOption) error
	Close()

	DeepPing() error
	DeepPingWithContext(ctx context.Context) error

	CreateTask(task api.InitialTaskDescription) (string, error)
	CreateTaskWithContext(ctx context.Context, task api.InitialTaskDescription) (string, error)

	DeleteTask(userID, taskID string) error
	DeleteTaskWithContext(ctx context.Context, userID, taskID string) error

	SetTaskCompleted(userID, taskID string) error
	SetTaskCompletedWithContext(ctx context.Context, userID, taskID string) error

	GetAllTasks(userID string) ([]api.TaskDescription, error)
	GetAllTasksWithContext(ctx context.Context, userID string) ([]api.TaskDescription, error)

	GetUncompletedTasks(userID string) ([]api.TaskDescription, error)
	GetUncompletedTasksWithContext(ctx context.Context, userID string) ([]api.TaskDescription, error)

	GetUncompletedTasksByHierarchy(nodeID string) (out []api.TaskDescription, err error)
	GetUncompletedTasksByHierarchyWithContext(ctx context.Context, nodeID string) (out []api.TaskDescription, err error)

	SetTaskStatus(taskID, userID string, status api.TaskStatus) (err error)
	SetTaskStatusWithContext(ctx context.Context, taskID, userID string, status api.TaskStatus) (err error)

	GetTaskStream(input api.GetTaskStreamInput, dc chan<- api.GetTaskStreamOutput) (err error)
	GetTaskStreamWithContext(ctx context.Context, input api.GetTaskStreamInput, dc chan<- api.GetTaskStreamOutput) (err error)

	IngestNodeData(nodeID string, nodeData api.NodeData) error
	IngestNodeDataWithContext(ctx context.Context, nodeID string, nodeData api.NodeData) error

	IngestNodeDataStream(c <-chan api.IngestNodeDataStreamInput) error
	IngestNodeDataStreamWithContext(ctx context.Context, c <-chan api.IngestNodeDataStreamInput) error

	GetLatestNodeData(input api.GetLatestNodeDataInput) (api.NodeData, error)
	GetLatestNodeDataWithContext(ctx context.Context, input api.GetLatestNodeDataInput) (api.NodeData, error)

	GetNodeData(input api.GetNodeDataInput) ([]api.NodeData, error)
	GetNodeDataWithContext(ctx context.Context, input api.GetNodeDataInput) ([]api.NodeData, error)

	GetNodeDataStream(input api.GetNodeDataStreamInput, c chan<- api.GetNodeDataStreamOutput) error
	GetNodeDataStreamWithContext(ctx context.Context, input api.GetNodeDataStreamInput, c chan<- api.GetNodeDataStreamOutput) error
}

type client struct {
	conn *grpc.ClientConn
	api  api.IoTClient
}

func CreateClient() IoTClient {
	return &client{}
}

func (c *client) Dial(host, port string, opts ...grpc.DialOption) (err error) {
	conn, err := grpc.Dial(host+":"+port, opts...)
	if err != nil {
		return
	}

	c.conn = conn
	c.api = api.NewIoTClient(conn)
	return
}

func (c *client) Close() {
	c.conn.Close()
}

func (c *client) DeepPing() error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	return c.DeepPingWithContext(ctx)
}

func (c *client) DeepPingWithContext(ctx context.Context) error {
	_, err := c.api.DeepPing(ctx, &api.PrimitiveVoid{})
	return err
}
