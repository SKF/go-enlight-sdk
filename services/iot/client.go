package iot

import (
	"context"
	"time"

	proto_common "github.com/SKF/proto/common"
	proto_iot "github.com/SKF/proto/iot"
	"google.golang.org/grpc"
)

type IoTClient interface {
	Dial(host, port string, opts ...grpc.DialOption) error
	Close()

	DeepPing() error
	DeepPingWithContext(ctx context.Context) error

	CreateTask(task proto_iot.InitialTaskDescription) (string, error)
	CreateTaskWithContext(ctx context.Context, task proto_iot.InitialTaskDescription) (string, error)

	DeleteTask(userID, taskID string) error
	DeleteTaskWithContext(ctx context.Context, userID, taskID string) error

	SetTaskCompleted(userID, taskID string) error
	SetTaskCompletedWithContext(ctx context.Context, userID, taskID string) error

	GetAllTasks(userID string) ([]proto_iot.TaskDescription, error)
	GetAllTasksWithContext(ctx context.Context, userID string) ([]proto_iot.TaskDescription, error)

	GetUncompletedTasks(userID string) ([]proto_iot.TaskDescription, error)
	GetUncompletedTasksWithContext(ctx context.Context, userID string) ([]proto_iot.TaskDescription, error)

	GetUncompletedTasksByHierarchy(nodeID string) (out []proto_iot.TaskDescription, err error)
	GetUncompletedTasksByHierarchyWithContext(ctx context.Context, nodeID string) (out []proto_iot.TaskDescription, err error)

	SetTaskStatus(input proto_iot.SetTaskStatusInput) (err error)
	SetTaskStatusWithContext(ctx context.Context, input proto_iot.SetTaskStatusInput) (err error)

	GetTaskStream(input proto_iot.GetTaskStreamInput, dc chan<- proto_iot.GetTaskStreamOutput) (err error)
	GetTaskStreamWithContext(ctx context.Context, input proto_iot.GetTaskStreamInput, dc chan<- proto_iot.GetTaskStreamOutput) (err error)

	GetTasksByStatus(input proto_iot.GetTasksByStatusInput) ([]*proto_iot.TaskDescription, error)
	GetTasksByStatusWithContext(ctx context.Context, input proto_iot.GetTasksByStatusInput) ([]*proto_iot.TaskDescription, error)

	GetTaskByUUID(input string) (*proto_iot.TaskDescription, error)
	GetTaskByUUIDWithContext(ctx context.Context, input string) (*proto_iot.TaskDescription, error)

	GetTaskByLongId(input int64) (*proto_iot.TaskDescription, error)
	GetTaskByLongIdWithContext(ctx context.Context, input int64) (*proto_iot.TaskDescription, error)

	IngestNodeData(input proto_iot.IngestNodeDataInput) error
	IngestNodeDataWithContext(ctx context.Context, input proto_iot.IngestNodeDataInput) error

	IngestNodeDataStream(c <-chan proto_iot.IngestNodeDataStreamInput) error
	IngestNodeDataStreamWithContext(ctx context.Context, c <-chan proto_iot.IngestNodeDataStreamInput) error

	GetLatestNodeData(input proto_iot.GetLatestNodeDataInput) (*proto_iot.NodeData, error)
	GetLatestNodeDataWithContext(ctx context.Context, input proto_iot.GetLatestNodeDataInput) (*proto_iot.NodeData, error)

	GetNodeData(input proto_iot.GetNodeDataInput) ([]proto_iot.NodeData, error)
	GetNodeDataWithContext(ctx context.Context, input proto_iot.GetNodeDataInput) ([]proto_iot.NodeData, error)

	GetNodeDataStream(input proto_iot.GetNodeDataStreamInput, c chan<- proto_iot.GetNodeDataStreamOutput) error
	GetNodeDataStreamWithContext(ctx context.Context, input proto_iot.GetNodeDataStreamInput, c chan<- proto_iot.GetNodeDataStreamOutput) error

	GetMedia(input proto_iot.GetMediaInput) (proto_iot.Media, error)
	GetMediaWithContext(ctx context.Context, input proto_iot.GetMediaInput) (proto_iot.Media, error)
}

type Client struct {
	conn *grpc.ClientConn
	api  proto_iot.IoTClient
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
	c.api = proto_iot.NewIoTClient(conn)
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
	_, err := c.api.DeepPing(ctx, &proto_common.Void{})
	return err
}
