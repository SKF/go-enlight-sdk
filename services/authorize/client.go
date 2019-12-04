package authorize

import (
	"context"
	"github.com/SKF/go-enlight-sdk/interceptors/reconnect"
	"github.com/SKF/go-utility/log"
	"github.com/aws/aws-sdk-go/aws/session"
	"google.golang.org/grpc/codes"
	"os"
	"time"

	"github.com/SKF/proto/common"
	"google.golang.org/grpc"

	authorize_grpcapi "github.com/SKF/proto/authorize"
)

type client struct {
	conn           *grpc.ClientConn
	api            authorize_grpcapi.AuthorizeClient
	requestTimeout time.Duration
}

type AuthorizeClient interface { // nolint: golint
	Dial(sess *session.Session, host, port, secretKey string, opts ...grpc.DialOption) error
	DialWithContext(ctx context.Context, sess *session.Session, host, port, secretKey string, opts ...grpc.DialOption) error
	Close() error
	SetRequestTimeout(d time.Duration)

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

	GetResource(id, originType string) (common.Origin, error)
	GetResourceWithContext(ctx context.Context, id, originType string) (common.Origin, error)

	AddResources(resources []common.Origin) error
	AddResourcesWithContext(ctx context.Context, resources []common.Origin) error

	RemoveResource(resource common.Origin) error
	RemoveResourceWithContext(ctx context.Context, resource common.Origin) error

	RemoveResources(resources []common.Origin) error
	RemoveResourcesWithContext(ctx context.Context, resources []common.Origin) error

	GetResourcesWithActionsAccess(actions []string, resourceType string, resource *common.Origin) ([]common.Origin, error)
	GetResourcesWithActionsAccessWithContext(ctx context.Context, actions []string, resourceType string, resource *common.Origin) ([]common.Origin, error)

	GetResourcesByUserAction(userID, actionName, resourceType string) ([]common.Origin, error)
	GetResourcesByUserActionWithContext(ctx context.Context, userID, actionName, resourceType string) ([]common.Origin, error)

	GetResourcesByType(resourceType string) (resources []common.Origin, err error)
	GetResourcesByTypeWithContext(ctx context.Context, resourceType string) (resources []common.Origin, err error)

	GetResourcesByOriginAndType(resource common.Origin, resourceType string, depth int32) (resources []common.Origin, err error)
	GetResourcesByOriginAndTypeWithContext(ctx context.Context, resource common.Origin, resourceType string, depth int32) (resources []common.Origin, err error)

	GetResourceParents(resource common.Origin, parentOriginType string) (resources []common.Origin, err error)
	GetResourceParentsWithContext(ctx context.Context, resource common.Origin, parentOriginType string) (resources []common.Origin, err error)

	GetResourceChildren(resource common.Origin, childOriginType string) (resources []common.Origin, err error)
	GetResourceChildrenWithContext(ctx context.Context, resource common.Origin, childOriginType string) (resources []common.Origin, err error)

	GetUserIDsWithAccessToResource(resource common.Origin) (resources []string, err error)
	GetUserIDsWithAccessToResourceWithContext(ctx context.Context, resource common.Origin) (resources []string, err error)

	AddResourceRelation(resource common.Origin, parent common.Origin) error
	AddResourceRelationWithContext(ctx context.Context, resource common.Origin, parent common.Origin) error

	AddResourceRelations(resources authorize_grpcapi.AddResourceRelationsInput) error
	AddResourceRelationsWithContext(ctx context.Context, resources authorize_grpcapi.AddResourceRelationsInput) error

	RemoveResourceRelation(resource common.Origin, parent common.Origin) error
	RemoveResourceRelationWithContext(ctx context.Context, resource common.Origin, parent common.Origin) error

	RemoveResourceRelations(resources authorize_grpcapi.RemoveResourceRelationsInput) error
	RemoveResourceRelationsWithContext(ctx context.Context, resources authorize_grpcapi.RemoveResourceRelationsInput) error

	ApplyUserAction(userID, action string, resource *common.Origin) error
	ApplyUserActionWithContext(ctx context.Context, userID, action string, resource *common.Origin) error

	RemoveUserAction(userID, action string, resource *common.Origin) error
	RemoveUserActionWithContext(ctx context.Context, userID, action string, resource *common.Origin) error

	GetActionsByUserRole(userRole string) ([]authorize_grpcapi.Action, error)
	GetActionsByUserRoleWithContext(ctx context.Context, userRole string) ([]authorize_grpcapi.Action, error)

	GetResourcesAndActionsByUser(userID string) ([]authorize_grpcapi.ActionResource, error)
	GetResourcesAndActionsByUserWithContext(ctx context.Context, userID string) ([]authorize_grpcapi.ActionResource, error)

	AddAction(action authorize_grpcapi.Action) error
	AddActionWithContext(ctx context.Context, action authorize_grpcapi.Action) error

	RemoveAction(name string) error
	RemoveActionWithContext(ctx context.Context, name string) error

	GetAction(name string) (authorize_grpcapi.Action, error)
	GetActionWithContext(ctx context.Context, name string) (authorize_grpcapi.Action, error)

	GetAllActions() ([]authorize_grpcapi.Action, error)
	GetAllActionsWithContext(ctx context.Context) ([]authorize_grpcapi.Action, error)

	GetUserActions(userID string) ([]authorize_grpcapi.Action, error)
	GetUserActionsWithContext(ctx context.Context, userID string) ([]authorize_grpcapi.Action, error)

	AddUserRole(role authorize_grpcapi.UserRole) error
	AddUserRoleWithContext(ctx context.Context, role authorize_grpcapi.UserRole) error

	GetUserRole(roleName string) (authorize_grpcapi.UserRole, error)
	GetUserRoleWithContext(ctx context.Context, roleName string) (authorize_grpcapi.UserRole, error)

	RemoveUserRole(roleName string) error
	RemoveUserRoleWithContext(ctx context.Context, roleName string) error
}

func CreateClient() AuthorizeClient {
	return &client{
		requestTimeout: 60 * time.Second,
	}
}

// Dial creates a client connection to the given host with background context and no timeout
func (c *client) Dial(sess *session.Session, host, port, secretKey string, opts ...grpc.DialOption) error {
	ctx, cancel := context.WithTimeout(context.Background(), c.requestTimeout)
	defer cancel()
	return c.DialWithContext(ctx, sess, host, port, secretKey, opts...)
}

// DialWithContext creates a client connection to the given host with context (for timeout and transaction id)
func (c *client) DialWithContext(ctx context.Context, sess *session.Session, host, port, secretKey string, opts ...grpc.DialOption) error {
	reconnectOpts := grpc.WithUnaryInterceptor(reconnect.UnaryInterceptor(
		reconnect.WithCodes(codes.DeadlineExceeded),
		reconnect.WithNewConnection(
			func(invokerCtx context.Context, invokerConn *grpc.ClientConn, invokerOptions ...grpc.CallOption) (context.Context, *grpc.ClientConn, []grpc.CallOption, error) {
				log.WithTracing(invokerCtx).Debug("Retrying with new connection")
				opt, err := getCredentialOption(ctx, sess, host, secretKey)
				if err != nil {
					log.WithTracing(invokerCtx).WithError(err).Error("Failed to get credential options")
					return invokerCtx, invokerConn, invokerOptions, err
				}
				_ = c.conn.Close()
				c.conn, err = grpc.DialContext(invokerCtx, host+":"+port, append(opts, opt, grpc.WithBlock())...)
				if err != nil {
					log.WithTracing(invokerCtx).WithError(err).Error("Failed to dial context")
					return invokerCtx, invokerConn, invokerOptions, err
				}
				c.api = authorize_grpcapi.NewAuthorizeClient(c.conn)
				return invokerCtx, c.conn, invokerOptions, err
			}),
	),
	)

	opt, err := getCredentialOption(ctx, sess, host, secretKey)
	if err != nil {
		return err
	}
	newOpts := append(opts, opt, reconnectOpts)

	conn, err := grpc.DialContext(ctx, host+":"+port, newOpts...)
	if err != nil {
		return err
	}

	c.conn = conn
	c.api = authorize_grpcapi.NewAuthorizeClient(c.conn)

	err = c.logClientState(ctx, "opening connection")
	return err
}

func (c *client) Close() (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.requestTimeout)
	defer cancel()
	err = c.logClientState(ctx, "closing connection")
	if c.conn != nil {
		return c.conn.Close()
	}
	return nil
}

func (c *client) DeepPing() error {
	ctx, cancel := context.WithTimeout(context.Background(), c.requestTimeout)
	defer cancel()
	return c.DeepPingWithContext(ctx)
}

func (c *client) DeepPingWithContext(ctx context.Context) error {
	_, err := c.api.DeepPing(ctx, &common.Void{})
	return err
}

func (c *client) logClientState(ctx context.Context, state string) error {
	hostname, err := os.Hostname()
	if err != nil {
		hostname = ""
	}
	_, err = c.api.LogClientState(ctx, &authorize_grpcapi.LogClientStateInput{
		State:    state,
		Hostname: hostname,
	})
	return err
}
