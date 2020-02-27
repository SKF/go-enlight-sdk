package iot

import (
	"context"
	"time"

	"github.com/SKF/proto/common"

	"google.golang.org/grpc"

	iot_grpcapi "github.com/SKF/proto/iot"
)

type IoTClient interface { // nolint: golint
	Dial(host, port string, opts ...grpc.DialOption) error
	DialWithContext(ctx context.Context, host, port string, opts ...grpc.DialOption) error
	Close()

	DeepPing() error
	DeepPingWithContext(ctx context.Context) error

	CreateTask(task iot_grpcapi.InitialTaskDescription) (string, error)
	CreateTaskWithContext(ctx context.Context, task iot_grpcapi.InitialTaskDescription) (string, error)

	DeleteTask(userID, taskID string) error
	DeleteTaskWithContext(ctx context.Context, userID, taskID string) error

	GetAllTasks(userID string) ([]iot_grpcapi.TaskDescription, error)
	GetAllTasksWithContext(ctx context.Context, userID string) ([]iot_grpcapi.TaskDescription, error)

	GetUncompletedTasks(userID string) ([]iot_grpcapi.TaskDescription, error)
	GetUncompletedTasksWithContext(ctx context.Context, userID string) ([]iot_grpcapi.TaskDescription, error)

	GetUncompletedTasksByHierarchy(nodeID string) (out []iot_grpcapi.TaskDescription, err error)
	GetUncompletedTasksByHierarchyWithContext(ctx context.Context, nodeID string) (out []iot_grpcapi.TaskDescription, err error)

	SetTaskStatus(input iot_grpcapi.SetTaskStatusInput) (err error)
	SetTaskStatusWithContext(ctx context.Context, input iot_grpcapi.SetTaskStatusInput) (err error)

	GetTasksByStatus(input iot_grpcapi.GetTasksByStatusInput) ([]*iot_grpcapi.TaskDescription, error)
	GetTasksByStatusWithContext(ctx context.Context, input iot_grpcapi.GetTasksByStatusInput) ([]*iot_grpcapi.TaskDescription, error)

	GetTaskByUUID(input string) (*iot_grpcapi.TaskDescription, error)
	GetTaskByUUIDWithContext(ctx context.Context, input string) (*iot_grpcapi.TaskDescription, error)

	GetTaskByLongId(input int64) (*iot_grpcapi.TaskDescription, error)
	GetTaskByLongIdWithContext(ctx context.Context, input int64) (*iot_grpcapi.TaskDescription, error)

	GetTasksModifiedSinceTimestamp(input iot_grpcapi.GetTasksModifiedSinceTimestampInput) (*iot_grpcapi.GetTasksModifiedSinceTimestampOutput, error)
	GetTasksModifiedSinceTimestampWithContext(ctx context.Context, input iot_grpcapi.GetTasksModifiedSinceTimestampInput) (*iot_grpcapi.GetTasksModifiedSinceTimestampOutput, error)

	IngestNodeData(input iot_grpcapi.IngestNodeDataInput) error
	IngestNodeDataWithContext(ctx context.Context, input iot_grpcapi.IngestNodeDataInput) error

	IngestNodesData(input iot_grpcapi.IngestNodesDataInput) error
	IngestNodesDataWithContext(ctx context.Context, input iot_grpcapi.IngestNodesDataInput) error

	GetLatestNodeData(input iot_grpcapi.GetLatestNodeDataInput) (*iot_grpcapi.NodeData, error)
	GetLatestNodeDataWithContext(ctx context.Context, input iot_grpcapi.GetLatestNodeDataInput) (*iot_grpcapi.NodeData, error)

	GetNodeData(input iot_grpcapi.GetNodeDataInput) ([]iot_grpcapi.NodeData, error)
	GetNodeDataWithContext(ctx context.Context, input iot_grpcapi.GetNodeDataInput) ([]iot_grpcapi.NodeData, error)

	GetMedia(input iot_grpcapi.GetMediaInput) (iot_grpcapi.Media, error)
	GetMediaWithContext(ctx context.Context, input iot_grpcapi.GetMediaInput) (iot_grpcapi.Media, error)

	RequestGetMediaSignedURL(in *iot_grpcapi.GetMediaSignedUrlInput) (*iot_grpcapi.GetMediaSignedUrlOutput, error)
	RequestGetMediaSignedURLWithContext(ctx context.Context, in *iot_grpcapi.GetMediaSignedUrlInput) (*iot_grpcapi.GetMediaSignedUrlOutput, error)

	RequestPutMediaSignedURL(in *iot_grpcapi.PutMediaSignedUrlInput) (*iot_grpcapi.PutMediaSignedUrlOutput, error)
	RequestPutMediaSignedURLWithContext(ctx context.Context, in *iot_grpcapi.PutMediaSignedUrlInput) (*iot_grpcapi.PutMediaSignedUrlOutput, error)

	DeleteNodeData(input iot_grpcapi.DeleteNodeDataInput) error
	DeleteNodeDataWithContext(ctx context.Context, input iot_grpcapi.DeleteNodeDataInput) error

	GetNodeEventLog(input iot_grpcapi.GetNodeEventLogInput) (*iot_grpcapi.GetNodeEventLogOutput, error)
	GetNodeEventLogWithContext(ctx context.Context, input iot_grpcapi.GetNodeEventLogInput) (*iot_grpcapi.GetNodeEventLogOutput, error)
}

type Client struct {
	conn *grpc.ClientConn
	api  iot_grpcapi.IoTClient
}

func CreateClient() IoTClient {
	return &Client{}
}

// Dial creates a client connection to the given host with background context and no timeout
func (c *Client) Dial(host, port string, opts ...grpc.DialOption) error {
	return c.DialWithContext(context.Background(), host, port, opts...)
}

// DialWithContext creates a client connection to the given host with context (for timeout and transaction id)
func (c *Client) DialWithContext(ctx context.Context, host, port string, opts ...grpc.DialOption) (err error) {
	conn, err := grpc.DialContext(ctx, host+":"+port, opts...)
	if err != nil {
		return
	}

	c.conn = conn
	c.api = iot_grpcapi.NewIoTClient(conn)
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
	_, err := c.api.DeepPing(ctx, &common.Void{})
	return err
}
