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

	SetTaskStatus(input api.SetTaskStatusInput) (err error)
	SetTaskStatusWithContext(ctx context.Context, input api.SetTaskStatusInput) (err error)

	GetTaskStream(input api.GetTaskStreamInput, dc chan<- api.GetTaskStreamOutput) (err error)
	GetTaskStreamWithContext(ctx context.Context, input api.GetTaskStreamInput, dc chan<- api.GetTaskStreamOutput) (err error)

	GetTasksByStatus(input api.GetTasksByStatusInput) ([]*api.TaskDescription, error)
	GetTasksByStatusWithContext(ctx context.Context, input api.GetTasksByStatusInput) ([]*api.TaskDescription, error)

	IngestNodeData(input api.IngestNodeDataInput) error
	IngestNodeDataWithContext(ctx context.Context, input api.IngestNodeDataInput) error

	IngestNodeDataStream(c <-chan api.IngestNodeDataStreamInput) error
	IngestNodeDataStreamWithContext(ctx context.Context, c <-chan api.IngestNodeDataStreamInput) error

	GetLatestNodeData(input api.GetLatestNodeDataInput) (*api.NodeData, error)
	GetLatestNodeDataWithContext(ctx context.Context, input api.GetLatestNodeDataInput) (*api.NodeData, error)

	GetNodeData(input api.GetNodeDataInput) ([]api.NodeData, error)
	GetNodeDataWithContext(ctx context.Context, input api.GetNodeDataInput) ([]api.NodeData, error)

	GetNodeDataStream(input api.GetNodeDataStreamInput, c chan<- api.GetNodeDataStreamOutput) error
	GetNodeDataStreamWithContext(ctx context.Context, input api.GetNodeDataStreamInput, c chan<- api.GetNodeDataStreamOutput) error

	GetMedia(input api.GetMediaInput) (api.Media, error)
	GetMediaWithContext(ctx context.Context, input api.GetMediaInput) (api.Media, error)
}

type Client struct {
	conn *grpc.ClientConn
	api  api.IoTClient
}

func CreateClient() IoTClient {
	return &Client{}
}

func (c *Client) Dial(host, port string, opts ...grpc.DialOption) (err error) {
	conn, err := grpc.Dial(host+":"+port, opts...)
	if err != nil {
		return
	}

	c.conn = conn
	c.api = api.NewIoTClient(conn)
	return
}

func (c *Client) Close() {
	c.conn.Close()
}

func (c *Client) DeepPing() error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	return c.DeepPingWithContext(ctx)
}

func (c *Client) DeepPingWithContext(ctx context.Context) error {
	_, err := c.api.DeepPing(ctx, &api.PrimitiveVoid{})
	return err
}
