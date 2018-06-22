package iam

import (
	"context"
	"time"

	"github.com/SKF/go-eventsource/eventsource"
	"google.golang.org/grpc"

	"github.com/SKF/go-enlight-sdk/services/iam/grpcapi"
)

type IAMClient interface {
	Dial(host, port string, opts ...grpc.DialOption) error
	Close()
	DeepPing() error

	CheckAuthentication(token, method string) (grpcapi.User, error)
	GetNodesByUser(userID string) (nodeIDs []string, err error)
	GetEventRecords(since int, limit *int32) ([]eventsource.Record, error)
}

type client struct {
	conn *grpc.ClientConn
	api  grpcapi.IAMClient
}

func CreateClient() IAMClient {
	return &client{}
}

func (c *client) Dial(host, port string, opts ...grpc.DialOption) (err error) {
	conn, err := grpc.Dial(host+":"+port, opts...)
	if err != nil {
		return
	}

	c.conn = conn
	c.api = grpcapi.NewIAMClient(conn)
	return
}

func (c *client) Close() {
	c.conn.Close()
}

func (c *client) DeepPing() error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	_, err := c.api.DeepPing(ctx, &grpcapi.PrimitiveVoid{})
	return err
}
