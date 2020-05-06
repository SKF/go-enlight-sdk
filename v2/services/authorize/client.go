package authorize

import (
	"context"
	"net"
	"os"

	"github.com/SKF/go-enlight-sdk/v2/interceptors/reconnect"
	"github.com/SKF/go-utility/v2/log"
	authorizeApi "github.com/SKF/proto/v2/authorize"
	"github.com/SKF/proto/v2/common"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

type client struct {
	conn *grpc.ClientConn
	api  authorizeApi.AuthorizeClient
}

type AuthorizeClient interface {
	Dial(ctx context.Context, host, port string, opts ...grpc.DialOption) error
	DialUsingCredentials(ctx context.Context, sess *session.Session, host, port, secretKey string, opts ...grpc.DialOption) error
	Close(ctx context.Context) error

	DeepPing(ctx context.Context) error

	IsAuthorized(ctx context.Context, userID, action string, resource *common.Origin) (bool, error)
	IsAuthorizedBulk(ctx context.Context, userID, action string, reqResources []common.Origin) ([]string, []bool, error)
	IsAuthorizedBulkWithResources(ctx context.Context, userID, action string, reqResources []common.Origin) ([]common.Origin, []bool, error)
	IsAuthorizedByEndpoint(ctx context.Context, api, method, endpoint, userID string) (bool, error)
	AddResource(ctx context.Context, resource common.Origin) error
	GetResource(ctx context.Context, id, originType string) (common.Origin, error)
	AddResources(ctx context.Context, resources []common.Origin) error
	RemoveResource(ctx context.Context, resource common.Origin) error
	RemoveResources(ctx context.Context, resources []common.Origin) error
	GetResourcesWithActionsAccess(ctx context.Context, actions []string, resourceType string, resource *common.Origin) ([]common.Origin, error)
	GetResourcesByUserAction(ctx context.Context, userID, actionName, resourceType string) ([]common.Origin, error)
	GetResourcesByType(ctx context.Context, resourceType string) (resources []common.Origin, err error)
	GetResourcesByOriginAndType(ctx context.Context, resource common.Origin, resourceType string, depth int32) (resources []common.Origin, err error)
	GetResourceParents(ctx context.Context, resource common.Origin, parentOriginType string) (resources []common.Origin, err error)
	GetResourceChildren(ctx context.Context, resource common.Origin, childOriginType string) (resources []common.Origin, err error)
	GetUserIDsWithAccessToResource(ctx context.Context, resource common.Origin) (resources []string, err error)
	AddResourceRelation(ctx context.Context, resource common.Origin, parent common.Origin) error
	AddResourceRelations(ctx context.Context, resources authorizeApi.AddResourceRelationsInput) error
	RemoveResourceRelation(ctx context.Context, resource common.Origin, parent common.Origin) error
	RemoveResourceRelations(ctx context.Context, resources authorizeApi.RemoveResourceRelationsInput) error
	ApplyUserAction(ctx context.Context, userID, action string, resource *common.Origin) error
	RemoveUserAction(ctx context.Context, userID, action string, resource *common.Origin) error
	GetActionsByUserRole(ctx context.Context, userRole string) ([]authorizeApi.Action, error)
	GetResourcesAndActionsByUser(ctx context.Context, userID string) ([]authorizeApi.ActionResource, error)
	GetResourcesAndActionsByUserAndResource(ctx context.Context, userID string, resource *common.Origin) ([]authorizeApi.ActionResource, error)
	AddAction(ctx context.Context, action authorizeApi.Action) error
	RemoveAction(ctx context.Context, name string) error
	GetAction(ctx context.Context, name string) (authorizeApi.Action, error)
	GetAllActions(ctx context.Context) ([]authorizeApi.Action, error)
	GetUserActions(ctx context.Context, userID string) ([]authorizeApi.Action, error)
	AddUserRole(ctx context.Context, role authorizeApi.UserRole) error
	GetUserRole(ctx context.Context, roleName string) (authorizeApi.UserRole, error)
	RemoveUserRole(ctx context.Context, roleName string) error
}

func CreateClient() AuthorizeClient {
	return &client{}
}

// DialWithContext creates a client connection to the given host with context (for timeout and transaction id)
func (c *client) Dial(ctx context.Context, host, port string, opts ...grpc.DialOption) (err error) {
	conn, err := grpc.DialContext(ctx, host+":"+port, opts...)
	if err != nil {
		return
	}

	c.conn = conn
	c.api = authorizeApi.NewAuthorizeClient(conn)
	err = c.logClientState(ctx, "opening connection")
	return
}

// DialUsingCredentialsWithContext creates a client connection to the given host with context (for timeout and transaction id)
func (c *client) DialUsingCredentials(ctx context.Context, sess *session.Session, host, port, secretKey string, opts ...grpc.DialOption) error {
	var newClientConn reconnect.NewConnectionFunc
	newClientConn = func(invokerCtx context.Context, invokerConn *grpc.ClientConn, invokerOptions ...grpc.CallOption) (context.Context, *grpc.ClientConn, []grpc.CallOption, error) {
		credOpt, err := getCredentialOption(invokerCtx, sess, host, secretKey)
		if err != nil {
			log.WithTracing(invokerCtx).WithError(err).Error("Failed to get credential options")
			return invokerCtx, invokerConn, invokerOptions, err
		}

		reconnectOpts := grpc.WithUnaryInterceptor(reconnect.UnaryInterceptor(
			reconnect.WithNewConnection(newClientConn),
		))

		dialOpts := append(opts, credOpt, reconnectOpts, grpc.WithBlock())
		newConn, err := grpc.DialContext(invokerCtx, net.JoinHostPort(host, port), dialOpts...)
		if err != nil {
			log.WithTracing(invokerCtx).WithError(err).Error("Failed to dial context")
			return invokerCtx, invokerConn, invokerOptions, err
		}
		_ = invokerConn.Close()

		c.conn = newConn
		c.api = authorizeApi.NewAuthorizeClient(c.conn)
		return invokerCtx, c.conn, invokerOptions, err
	}

	opt, err := getCredentialOption(ctx, sess, host, secretKey)
	if err != nil {
		return err
	}

	reconnectOpts := grpc.WithUnaryInterceptor(reconnect.UnaryInterceptor(
		reconnect.WithNewConnection(newClientConn),
	))
	newOpts := append(opts, opt, reconnectOpts)

	conn, err := grpc.DialContext(ctx, host+":"+port, newOpts...)
	if err != nil {
		return err
	}

	c.conn = conn
	c.api = authorizeApi.NewAuthorizeClient(c.conn)

	err = c.logClientState(ctx, "opening connection")
	return err
}

func (c *client) Close(ctx context.Context) (err error) {
	if c.conn != nil {
		defer func() {
			if errClose := c.conn.Close(); errClose != nil {
				if err != nil {
					err = errors.Wrapf(errClose, err.Error())
				} else {
					err = errClose
				}
			}
		}()

		err = c.logClientState(ctx, "closing connection")
		return err
	}
	return nil
}

func (c *client) DeepPing(ctx context.Context) error {
	_, err := c.api.DeepPing(ctx, &common.Void{})
	return err
}

func (c *client) logClientState(ctx context.Context, state string) error {
	hostname, err := os.Hostname()
	if err != nil {
		hostname = ""
	}
	_, err = c.api.LogClientState(ctx, &authorizeApi.LogClientStateInput{
		State:    state,
		Hostname: hostname,
	})
	return err
}
