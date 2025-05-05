package iam

import (
	"context"
	"time"

	"github.com/SKF/proto/v2/common"

	"google.golang.org/grpc"

	"github.com/SKF/go-eventsource/eventsource"
	iam_grpcapi "github.com/SKF/proto/v2/iam"
)

type IAMClient interface { // nolint: golint
	Dial(host, port string, opts ...grpc.DialOption) error
	DialWithContext(ctx context.Context, host, port string, opts ...grpc.DialOption) error
	Close()

	DeepPing() error
	DeepPingWithContext(ctx context.Context) error

	CheckAuthentication(token, arn string) (*iam_grpcapi.UserClaims, error)
	CheckAuthenticationWithContext(ctx context.Context, token, arn string) (*iam_grpcapi.UserClaims, error)

	CheckAuthenticationByEndpoint(token, api, method, endpoint string) (*iam_grpcapi.UserClaims, error)
	CheckAuthenticationByEndpointWithContext(ctx context.Context, token, api, method, endpoint string) (*iam_grpcapi.UserClaims, error)

	GetNodesByUser(userID string) (nodeIDs []string, err error)
	GetNodesByUserWithContext(ctx context.Context, userID string) (nodeIDs []string, err error)

	GetEventRecords(since int, limit *int32) ([]eventsource.Record, error)
	GetEventRecordsWithContext(ctx context.Context, since int, limit *int32) ([]eventsource.Record, error)
}

type client struct {
	conn *grpc.ClientConn
	api  iam_grpcapi.IAMClient
}

func CreateClient() IAMClient {
	return &client{}
}

// Dial creates a client connection to the given host with background context and no timeout
func (c *client) Dial(host, port string, opts ...grpc.DialOption) error {
	return c.DialWithContext(context.Background(), host, port, opts...)
}

// DialWithContext creates a client connection to the given host with context (for timeout and transaction id)
func (c *client) DialWithContext(ctx context.Context, host, port string, opts ...grpc.DialOption) (err error) {
	conn, err := grpc.DialContext(ctx, host+":"+port, opts...)
	if err != nil {
		return
	}

	c.conn = conn
	c.api = iam_grpcapi.NewIAMClient(conn)
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
	_, err := c.api.DeepPing(ctx, &common.Void{})
	return err
}
