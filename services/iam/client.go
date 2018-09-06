package iam

import (
	"context"
	"time"

	"github.com/SKF/go-eventsource/eventsource"
	proto_common "github.com/SKF/proto/common"
	proto_iam "github.com/SKF/proto/iam"
	"google.golang.org/grpc"
)

type IAMClient interface {
	Dial(host, port string, opts ...grpc.DialOption) error
	Close()

	DeepPing() error
	DeepPingWithContext(ctx context.Context) error

	CheckAuthentication(token, method string) (proto_iam.User, error)
	CheckAuthenticationWithContext(ctx context.Context, token, method string) (proto_iam.User, error)

	GetNodesByUser(userID string) (nodeIDs []string, err error)
	GetNodesByUserWithContext(ctx context.Context, userID string) (nodeIDs []string, err error)

	GetEventRecords(since int, limit *int32) ([]eventsource.Record, error)
	GetEventRecordsWithContext(ctx context.Context, since int, limit *int32) ([]eventsource.Record, error)
}

type client struct {
	conn *grpc.ClientConn
	api  proto_iam.IAMClient
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
	c.api = proto_iam.NewIAMClient(conn)
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
	_, err := c.api.DeepPing(ctx, &proto_common.Void{})
	return err
}
