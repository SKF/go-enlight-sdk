package iam

import (
	"context"
	"time"

	"github.com/SKF/proto/common"

	"github.com/SKF/go-eventsource/eventsource"
	iam_grpcapi "github.com/SKF/proto/iam"
	"google.golang.org/grpc"
)

type IAMClient interface {
	Dial(host, port string, opts ...grpc.DialOption) error
	Close()

	DeepPing() error
	DeepPingWithContext(ctx context.Context) error

	CheckAuthentication(token, method string) (iam_grpcapi.User, error)
	CheckAuthenticationWithContext(ctx context.Context, token, method string) (iam_grpcapi.User, error)

	GetNodesByUser(userID string) (nodeIDs []string, err error)
	GetNodesByUserWithContext(ctx context.Context, userID string) (nodeIDs []string, err error)

	GetEventRecords(since int, limit *int32) ([]eventsource.Record, error)
	GetEventRecordsWithContext(ctx context.Context, since int, limit *int32) ([]eventsource.Record, error)

	IsAuthorized(userID, action string, resource *common.Origin) error
	IsAuthorizedWithContext(ctx context.Context, userID, action string, resource *common.Origin) error

	AddResource(resource common.Origin) error
	AddResourceWithContext(ctx context.Context, resource common.Origin) error

	RemoveResource(resource common.Origin) error
	RemoveResourceWithContext(ctx context.Context, resource common.Origin) error

	AddResourceParent(resource common.Origin, parent common.Origin) error
	AddResourceParentWithContext(ctx context.Context, resource common.Origin, parent common.Origin) error

	RemoveResourceParent(resource common.Origin, parent common.Origin) error
	RemoveResourceParentWithContext(ctx context.Context, resource common.Origin, parent common.Origin) error

	AddUserRole(userID, role string, resource common.Origin) error
	AddUserRoleWithContext(ctx context.Context, userID, role string, resource common.Origin) error

	RemoveUserRole(userID, role string, resource common.Origin) error
	RemoveUserRoleWithContext(ctx context.Context, userID, role string, resource common.Origin) error
}

type client struct {
	conn *grpc.ClientConn
	api  iam_grpcapi.IAMClient
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
