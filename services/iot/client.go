package iot

import (
	"context"
	"time"

	"google.golang.org/grpc"

	api "github.com/SKF/go-enlight-sdk/services/iot/iot_grpc_api"
)

type IoTClient interface {
	Dial(host, port string, opts ...grpc.DialOption) error
	Close()
	DeepPing() error

	CreateTask(task api.InitialTaskDescription) (string, error)
	DeleteTask(userID, taskID string) error
	SetTaskCompleted(userID, taskID string) error
	GetAllTasks(userID string) ([]api.TaskDescription, error)
	GetUncompletedTasks(userID string) ([]api.TaskDescription, error)
	GetUncompletedTasksByHierarchy(nodeID string) (out []api.TaskDescription, err error)
	SetTaskStatus(taskID, userID string, status api.TaskStatus) (err error)

	IngestNodeData(nodeID string, nodeData api.NodeData) error
	GetNodeData(input api.GetNodeDataInput) ([]api.NodeData, error)
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

	_, err := c.api.DeepPing(ctx, &api.PrimitiveVoid{})
	return err
}
