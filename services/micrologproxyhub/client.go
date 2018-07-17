package micrologproxyhub

import (
	"context"
	"time"

	api "github.com/SKF/go-enlight-sdk/services/micrologproxyhub/grpcapi"

	"github.com/SKF/go-utility/uuid"
	"google.golang.org/grpc"
)

type MicrologProxyHubClient interface {
	Dial(host, port string, opts ...grpc.DialOption) error
	Close()
	DeepPing() error

	SetTaskStatus(taskID, userID uuid.UUID, status api.TaskStatus) error
	GetTasksStream(dc chan<- api.GetTasksStreamOutput) error
}

func CreateClient() MicrologProxyHubClient {
	return &client{}
}

type client struct {
	api  api.MicrologProxyHubClient
	conn *grpc.ClientConn
}

func (c *client) Dial(host, port string, opts ...grpc.DialOption) (err error) {
	conn, err := grpc.Dial(host+":"+port, opts...)
	if err != nil {
		return
	}

	c.conn = conn
	c.api = api.NewMicrologProxyHubClient(conn)
	return
}

func (c *client) Close() {
	c.conn.Close()
}

func (c *client) DeepPing() (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	_, err = c.api.DeepPing(ctx, &api.Void{})
	return
}
