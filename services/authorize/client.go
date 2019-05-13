package authorize

import (
	"context"
	"time"

	"github.com/SKF/proto/common"

	"google.golang.org/grpc"

	authorize_grpcapi "github.com/SKF/proto/authorize"
)

type AuthorizeClient interface { // nolint: golint
	Dial(host, port string, opts ...grpc.DialOption) error
	DialWithContext(ctx context.Context, host, port string, opts ...grpc.DialOption) error
	Close()

	DeepPing() error
	DeepPingWithContext(ctx context.Context) error

	IsAuthorized(userID, action string, resource *common.Origin) (bool, error)
	IsAuthorizedWithContext(ctx context.Context, userID, action string, resource *common.Origin) (bool, error)
	IsAuthorizedBulk(userID, action string, resource []common.Origin) ([]string, []bool, error)
	IsAuthorizedBulkWithContext(ctx context.Context, userID, action string, resource []common.Origin) ([]string, []bool, error)
	IsAuthorizedByEndpoint(api, method, endpoint, userID string) (bool, error)
	IsAuthorizedByEndpointWithContext(ctx context.Context, api, method, endpoint, userID string) (bool, error)

	AddResource(resource common.Origin) error
	AddResourceWithContext(ctx context.Context, resource common.Origin) error

	AddResources(resources []common.Origin) error
	AddResourcesWithContext(ctx context.Context, resources []common.Origin) error

	RemoveResource(resource common.Origin) error
	RemoveResourceWithContext(ctx context.Context, resource common.Origin) error

	RemoveResources(resources []common.Origin) error
	RemoveResourcesWithContext(ctx context.Context, resources []common.Origin) error

	GetResourcesByType(resourceType string) (resources []common.Origin, err error)
	GetResourcesByTypeWithContext(ctx context.Context, resourceType string) (resources []common.Origin, err error)

	GetResourcesByOriginAndType(originID string, resourceType string) (resources []common.Origin, err error)
	GetResourcesByOriginAndTypeWithContext(ctx context.Context, originID string, resourceType string) (resources []common.Origin, err error)

	GetUserIDsWithAccessToResource(originID string) (resources []string, err error)
	GetUserIDsWithAccessToResourceWithContext(ctx context.Context, originID string) (resources []string, err error)

	AddResourceRelation(resource common.Origin, parent common.Origin) error
	AddResourceRelationWithContext(ctx context.Context, resource common.Origin, parent common.Origin) error

	AddResourceRelations(resources authorize_grpcapi.AddResourceRelationsInput) error
	AddResourceRelationsWithContext(ctx context.Context, resources authorize_grpcapi.AddResourceRelationsInput) error

	RemoveResourceRelation(resource common.Origin, parent common.Origin) error
	RemoveResourceRelationWithContext(ctx context.Context, resource common.Origin, parent common.Origin) error

	RemoveResourceRelations(resources authorize_grpcapi.RemoveResourceRelationsInput) error
	RemoveResourceRelationsWithContext(ctx context.Context, resources authorize_grpcapi.RemoveResourceRelationsInput) error

	AddUserPermission(userID, action string, resource *common.Origin) error
	AddUserPermissionWithContext(ctx context.Context, userID, action string, resource *common.Origin) error

	RemoveUserPermission(userID, action string, resource *common.Origin) error
	RemoveUserPermissionWithContext(ctx context.Context, userID, action string, resource *common.Origin) error

	GetActionsByUserRole(userRole string) ([]authorize_grpcapi.Action, error)
	GetActionsByUserRoleWithContext(ctx context.Context, userRole string) ([]authorize_grpcapi.Action, error)

	GetResourcesAndActionsByUser(userID string) ([]authorize_grpcapi.ActionResource, error)
	GetResourcesAndActionsByUserWithContext(ctx context.Context, userID string) ([]authorize_grpcapi.ActionResource, error)
}

type client struct {
	conn *grpc.ClientConn
	api  authorize_grpcapi.AuthorizeClient
}

func CreateClient() AuthorizeClient {
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
	c.api = authorize_grpcapi.NewAuthorizeClient(conn)
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
