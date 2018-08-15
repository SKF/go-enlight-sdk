package mhub

import (
	"context"
	"time"

	"github.com/SKF/go-enlight-sdk/services/mhub/mhubapi"

	"github.com/SKF/go-utility/uuid"
	"google.golang.org/grpc"
)

type MicrologProxyHubClient interface {
	Dial(host, port string, opts ...grpc.DialOption) error
	Close()

	DeepPing() error
	DeepPingWithContext(ctx context.Context) error

	SetTaskStatus(id int64, userID uuid.UUID, status mhubapi.TaskStatus) error
	SetTaskStatusWithContext(ctx context.Context, id int64, userID uuid.UUID, status mhubapi.TaskStatus) error

	AvailableDSKFStream(dc chan<- mhubapi.AvailableDSKFStreamOutput) error
	AvailableDSKFStreamWithContext(ctx context.Context, dc chan<- mhubapi.AvailableDSKFStreamOutput) error
}

func CreateClient() MicrologProxyHubClient {
	return &client{}
}

type client struct {
	api  mhubapi.MicrologProxyHubClient
	conn *grpc.ClientConn
}

func (c *client) Dial(host, port string, opts ...grpc.DialOption) (err error) {
	conn, err := grpc.Dial(host+":"+port, opts...)
	if err != nil {
		return
	}

	c.conn = conn
	c.api = mhubapi.NewMicrologProxyHubClient(conn)
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
	_, err := c.api.DeepPing(ctx, &mhubapi.Void{})
	return err
}
