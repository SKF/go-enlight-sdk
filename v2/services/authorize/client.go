package authorize

import (
	"context"
	_ "embed"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws/session"
	"google.golang.org/grpc"
	"google.golang.org/grpc/resolver"

	"github.com/SKF/go-enlight-sdk/v2/services/authorize/credentialsmanager"
	authorizeApi "github.com/SKF/proto/v2/authorize"
	"github.com/SKF/proto/v2/common"
)

//go:embed service_config.json
var defaultServiceConfig string

type client struct {
	conn               *grpc.ClientConn
	api                authorizeApi.AuthorizeClient
	requestTimeout     time.Duration
	credentialsFetcher credentialsmanager.CredentialsFetcher
}

type AuthorizeClient interface {
	Dial(host, port string, opts ...grpc.DialOption) error
	DialWithContext(ctx context.Context, host, port string, opts ...grpc.DialOption) error
	DialUsingCredentials(sess *session.Session, host, port, secretKey string, opts ...grpc.DialOption) error
	DialUsingCredentialsWithContext(ctx context.Context, sess *session.Session, host, port, secretKey string, opts ...grpc.DialOption) error
	DialUsingCredentialsManager(ctx context.Context, cf credentialsmanager.CredentialsFetcher, host, port, secretKey string, opts ...grpc.DialOption) error

	Close() error
	SetRequestTimeout(d time.Duration)

	DeepPing() error
	DeepPingWithContext(ctx context.Context) error

	IsAuthorized(userID, action string, resource *common.Origin) (bool, error)
	IsAuthorizedWithContext(ctx context.Context, userID, action string, resource *common.Origin) (bool, error)
	IsAuthorizedBulk(userID, action string, resResources []common.Origin) ([]string, []bool, error)
	IsAuthorizedBulkWithContext(ctx context.Context, userID, action string, reqResources []common.Origin) ([]string, []bool, error)
	IsAuthorizedBulkWithResources(ctx context.Context, userID, action string, reqResources []common.Origin) ([]common.Origin, []bool, error)
	IsAuthorizedByEndpoint(api, method, endpoint, userID string) (bool, error)
	IsAuthorizedByEndpointWithContext(ctx context.Context, api, method, endpoint, userID string) (bool, error)
	IsAuthorizedWithReason(userID, action string, resource *common.Origin) (bool, string, error)
	IsAuthorizedWithReasonWithContext(ctx context.Context, userID, action string, resource *common.Origin) (bool, string, error)

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

	AddResourceRelations(resources authorizeApi.AddResourceRelationsInput) error
	AddResourceRelationsWithContext(ctx context.Context, resources authorizeApi.AddResourceRelationsInput) error

	RemoveResourceRelation(resource common.Origin, parent common.Origin) error
	RemoveResourceRelationWithContext(ctx context.Context, resource common.Origin, parent common.Origin) error

	RemoveResourceRelations(resources authorizeApi.RemoveResourceRelationsInput) error
	RemoveResourceRelationsWithContext(ctx context.Context, resources authorizeApi.RemoveResourceRelationsInput) error

	ApplyUserAction(userID, action string, resource *common.Origin) error
	ApplyUserActionWithContext(ctx context.Context, userID, action string, resource *common.Origin) error

	ApplyRolesForUserOnResources(userID string, roles []string, resources []common.Origin) error
	ApplyRolesForUserOnResourcesWithContext(ctx context.Context, userID string, roles []string, resources []common.Origin) error

	RemoveUserAction(userID, action string, resource *common.Origin) error
	RemoveUserActionWithContext(ctx context.Context, userID, action string, resource *common.Origin) error

	GetActionsByUserRole(userRole string) ([]authorizeApi.Action, error)
	GetActionsByUserRoleWithContext(ctx context.Context, userRole string) ([]authorizeApi.Action, error)

	GetResourcesAndActionsByUser(userID string) ([]authorizeApi.ActionResource, error)
	GetResourcesAndActionsByUserWithContext(ctx context.Context, userID string) ([]authorizeApi.ActionResource, error)

	GetResourcesAndActionsByUserAndResource(userID string, resource *common.Origin) ([]authorizeApi.ActionResource, error)
	GetResourcesAndActionsByUserAndResourceWithContext(ctx context.Context, userID string, resource *common.Origin) ([]authorizeApi.ActionResource, error)

	AddAction(action authorizeApi.Action) error
	AddActionWithContext(ctx context.Context, action authorizeApi.Action) error

	RemoveAction(name string) error
	RemoveActionWithContext(ctx context.Context, name string) error

	GetAction(name string) (authorizeApi.Action, error)
	GetActionWithContext(ctx context.Context, name string) (authorizeApi.Action, error)

	GetAllActions() ([]authorizeApi.Action, error)
	GetAllActionsWithContext(ctx context.Context) ([]authorizeApi.Action, error)

	GetUserActions(userID string) ([]authorizeApi.Action, error)
	GetUserActionsWithContext(ctx context.Context, userID string) ([]authorizeApi.Action, error)

	AddUserRole(role authorizeApi.UserRole) error
	AddUserRoleWithContext(ctx context.Context, role authorizeApi.UserRole) error

	GetUserRole(roleName string) (authorizeApi.UserRole, error)
	GetUserRoleWithContext(ctx context.Context, roleName string) (authorizeApi.UserRole, error)

	RemoveUserRole(roleName string) error
	RemoveUserRoleWithContext(ctx context.Context, roleName string) error
}

func CreateClient() AuthorizeClient {
	return &client{
		requestTimeout: 60 * time.Second,
	}
}

func (c *client) withCredentialsFetcher(credentialsFetcher credentialsmanager.CredentialsFetcher) *client {
	c.credentialsFetcher = credentialsFetcher

	return c
}

// Dial creates a client connection to the given host with background context and no timeout
func (c *client) Dial(host, port string, opts ...grpc.DialOption) error {
	ctx, cancel := context.WithTimeout(context.Background(), c.requestTimeout)
	defer cancel()
	return c.DialWithContext(ctx, host, port, opts...)
}

// DialWithContext creates a client connection to the given host with context (for timeout and transaction id)
func (c *client) DialWithContext(ctx context.Context, host, port string, opts ...grpc.DialOption) (err error) {
	resolver.SetDefaultScheme("dns")
	opts = append([]grpc.DialOption{grpc.WithDefaultServiceConfig(defaultServiceConfig)}, opts...)

	conn, err := grpc.DialContext(ctx, host+":"+port, opts...)
	if err != nil {
		return
	}

	c.conn = conn
	c.api = authorizeApi.NewAuthorizeClient(conn)
	err = c.logClientState(ctx, "opening connection")
	return
}

// DialUsingCredentials creates a client connection to the given host with background context and no timeout
func (c *client) DialUsingCredentials(sess *session.Session, host, port, secretKey string, opts ...grpc.DialOption) error {
	ctx, cancel := context.WithTimeout(context.Background(), c.requestTimeout)
	defer cancel()
	return c.DialUsingCredentialsWithContext(ctx, sess, host, port, secretKey, opts...)
}

// DialUsingCredentialsWithContext creates a client connection to the given host with context (for timeout and transaction id)
func (c *client) DialUsingCredentialsWithContext(ctx context.Context, sess *session.Session, host, port, secretKey string, opts ...grpc.DialOption) error {
	cm := credentialsmanager.New().UsingSDKV1Session(sess)
	return c.DialUsingCredentialsManager(ctx, cm, host, port, secretKey, opts...)
}

func (c *client) DialUsingCredentialsManager(ctx context.Context, cf credentialsmanager.CredentialsFetcher, host, port, secretKey string, opts ...grpc.DialOption) error {
	return c.withCredentialsFetcher(cf).
		dialUsingCredentials(ctx, host, port, secretKey, opts...)
}

func (c *client) dialUsingCredentials(ctx context.Context, host, port, secretKey string, opts ...grpc.DialOption) error {
	resolver.SetDefaultScheme("dns")
	opts = append([]grpc.DialOption{grpc.WithDefaultServiceConfig(defaultServiceConfig)}, opts...)

	opt, err := getCredentialOption(ctx, host, secretKey, c.credentialsFetcher)
	if err != nil {
		return err
	}

	newOpts := append(opts, opt)

	conn, err := grpc.DialContext(ctx, host+":"+port, newOpts...)
	if err != nil {
		return err
	}

	c.conn = conn
	c.api = authorizeApi.NewAuthorizeClient(c.conn)

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
	_, err = c.api.LogClientState(ctx, &authorizeApi.LogClientStateInput{
		State:    state,
		Hostname: hostname,
	})
	return err
}
