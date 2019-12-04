package authorize

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/SKF/go-enlight-sdk/grpc"
	"github.com/SKF/go-enlight-sdk/interceptors/reconnect"
	"github.com/SKF/go-utility/log"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"os"
	"time"

	"github.com/SKF/proto/common"

	googleGrpc "google.golang.org/grpc"

	authorize_grpcapi "github.com/SKF/proto/authorize"
)

type AuthorizeClient interface { // nolint: golint
	Dial(sess *session.Session, host, port, secretKey string, opts ...googleGrpc.DialOption) error
	DialWithContext(ctx context.Context, sess *session.Session, host, port, secretKey string, opts ...googleGrpc.DialOption) error
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

func getSecret(sess *session.Session, secretsName string, out interface{}) (err error) {
	// credentials - default
	svc := secretsmanager.New(sess)
	input := &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(secretsName),
	}

	result, err := svc.GetSecretValue(input)
	if err != nil {
		log.WithError(err).
			WithField("secretsName", secretsName).
			Error("failed to get secrets")
		err = errors.Wrapf(err, "failed to get secret value from '%s'", secretsName)
		return
	}

	if err = json.Unmarshal([]byte(*result.SecretString), out); err != nil {
		log.WithError(err).
			WithField("secretsName", secretsName).
			Error("failed to unmarshal secret")
		err = errors.Wrapf(err, "failed to unmarshal secret from '%s'", secretsName)
	}

	return err
}

type dataStore struct {
	CA  []byte `json:"ca"`
	Key []byte `json:"key"`
	Crt []byte `json:"crt"`
}

func getCredentialOption(sess *session.Session, host, secretKeyName string) (googleGrpc.DialOption, error) {
	var clientCert dataStore
	if err := getSecret(sess, secretKeyName, &clientCert); err != nil {
		panic(err)
	}

	return grpc.WithTransportCredentialsPEM(
		host,
		clientCert.Crt, clientCert.Key, clientCert.CA,
	)
}

type client struct {
	conn           *googleGrpc.ClientConn
	api            authorize_grpcapi.AuthorizeClient
	requestTimeout time.Duration
}

func CreateClient() AuthorizeClient {
	return &client{
		requestTimeout: 60 * time.Second,
	}
}

func GetSecretKeyName(service, stage string) string {
	return fmt.Sprintf("authorize/%s/grpc/client/%s", stage, service)
}

func GetSecretKeyArn(accountId, region, service, stage string) string {
	return fmt.Sprintf("arn:aws:secretsmanager:%s:%s:secret:%s", region, accountId, GetSecretKeyName(service, stage))
}

// Dial creates a client connection to the given host with background context and no timeout
func (c *client) Dial(sess *session.Session, host, port, secretKey string, opts ...googleGrpc.DialOption) error {
	ctx, cancel := context.WithTimeout(context.Background(), c.requestTimeout)
	defer cancel()
	return c.DialWithContext(ctx, sess, host, port, secretKey, opts...)
}

// DialWithContext creates a client connection to the given host with context (for timeout and transaction id)
func (c *client) DialWithContext(ctx context.Context, sess *session.Session, host, port, secretKey string, opts ...googleGrpc.DialOption) error {
	reconnectOpts := googleGrpc.WithUnaryInterceptor(reconnect.UnaryInterceptor(
		reconnect.WithCodes(codes.DeadlineExceeded),
		reconnect.WithNewConnection(
			func(_ context.Context, invokerConn *googleGrpc.ClientConn, invokerOptions ...googleGrpc.CallOption) (context.Context, *googleGrpc.ClientConn, []googleGrpc.CallOption, error) {
				opt, err := getCredentialOption(sess, host, secretKey)
				if err != nil {
					return context.Background(), invokerConn, invokerOptions, err
				}
				c.conn, err = googleGrpc.DialContext(context.Background(), host+":"+port, append(opts, opt)...)
				if err != nil {
					return context.Background(), invokerConn, invokerOptions, err
				}
				c.api = authorize_grpcapi.NewAuthorizeClient(c.conn)
				return context.Background(), c.conn, invokerOptions, err
			}),
	),
	)

	opt, err := getCredentialOption(sess, host, secretKey)
	if err != nil {
		return err
	}
	newOpts := append(opts, opt, reconnectOpts)

	conn, err := googleGrpc.DialContext(ctx, host+":"+port, newOpts...)
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
	return c.conn.Close()
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
