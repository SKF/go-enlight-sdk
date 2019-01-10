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

	IsAuthorized(userID, action string, resource *common.Origin) (bool, error)
	IsAuthorizedWithContext(ctx context.Context, userID, action string, resource *common.Origin) (bool, error)

	AddAuthorizationResource(resource common.Origin) error
	AddAuthorizationResourceWithContext(ctx context.Context, resource common.Origin) error

	RemoveAuthorizationResource(resource common.Origin) error
	RemoveAuthorizationResourceWithContext(ctx context.Context, resource common.Origin) error

	AddAuthorizationResourceRelation(resource common.Origin, parent common.Origin) error
	AddAuthorizationResourceRelationWithContext(ctx context.Context, resource common.Origin, parent common.Origin) error

	RemoveAuthorizationResourceRelation(resource common.Origin, parent common.Origin) error
	RemoveAuthorizationResourceRelationWithContext(ctx context.Context, resource common.Origin, parent common.Origin) error

	AddUserPermission(userID, role string, resource common.Origin) error
	AddUserPermissionWithContext(ctx context.Context, userID, role string, resource common.Origin) error

	RemoveUserPermission(userID, role string, resource common.Origin) error
	RemoveUserPermissionWithContext(ctx context.Context, userID, role string, resource common.Origin) error
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
