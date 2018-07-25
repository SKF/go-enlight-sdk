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

	SetTaskStatus(taskID, userID uuid.UUID, status mhubapi.TaskStatus) error
	AvailableDSKFStream(dc chan<- mhubapi.AvailableDSKFStreamOutput) error
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

func (c *client) DeepPing() (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	_, err = c.api.DeepPing(ctx, &mhubapi.Void{})
	return
}
