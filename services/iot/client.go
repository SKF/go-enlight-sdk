package iot

import (
	"context"
	"time"

	"google.golang.org/grpc"

	"github.com/SKF/go-enlight-sdk/services/iot/iotgrpcapi"
)

type IoTClient interface {
	Dial(host, port string, opts ...grpc.DialOption) error
	Close()
	DeepPing() error

	CreateTask(task iotgrpcapi.InitialTaskDescription) (string, error)
	DeleteTask(userID, taskID string) error
	SetTaskCompleted(userID, taskID string) error
	GetAllTasks(userID string) ([]iotgrpcapi.TaskDescription, error)
	GetUncompletedTasks(userID string) ([]iotgrpcapi.TaskDescription, error)
}

type client struct {
	conn *grpc.ClientConn
	api  iotgrpcapi.IoTClient
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
	c.api = iotgrpcapi.NewIoTClient(conn)
	return
}

func (c *client) Close() {
	c.conn.Close()
}

func (c *client) DeepPing() error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	_, err := c.api.DeepPing(ctx, &iotgrpcapi.PrimitiveVoid{})
	return err
}
