package mock

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go/aws/session"

	grpcapi "github.com/SKF/proto/authorize"

	"github.com/SKF/proto/common"

	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc"

	"github.com/SKF/go-enlight-sdk/services/authorize"
	"github.com/SKF/go-enlight-sdk/services/authorize/credentialsmanager"
)

type Client struct {
	mock.Mock
}

func Create() *Client { // nolint: golint
	return new(Client)
}

var _ authorize.AuthorizeClient = &Client{}

func (mock *Client) SetRequestTimeout(d time.Duration) {
	mock.Called(d)
}

func (mock *Client) Dial(host, port string, opts ...grpc.DialOption) error {
	args := mock.Called(host, port, opts)
	return args.Error(0)
}

func (mock *Client) DialWithContext(ctx context.Context, host, port string, opts ...grpc.DialOption) error {
	args := mock.Called(ctx, host, port, opts)
	return args.Error(0)
}

func (mock *Client) DialUsingCredentials(sess *session.Session, host, port, secretKey string, opts ...grpc.DialOption) error {
	args := mock.Called(sess, host, port, secretKey, opts)
	return args.Error(0)
}

func (mock *Client) DialUsingCredentialsWithContext(ctx context.Context, sess *session.Session, host, port, secretKey string, opts ...grpc.DialOption) error {
	args := mock.Called(ctx, sess, host, port, secretKey, opts)
	return args.Error(0)
}

func (mock *Client) DialUsingCredentialsManager(ctx context.Context, cm *credentialsmanager.CredentialsManager, host, port, secretKey string, opts ...grpc.DialOption) error {
	args := mock.Called(ctx, cm, host, port, secretKey, opts)
	return args.Error(0)
}

func (mock *Client) Close() error {
	args := mock.Called()
	return args.Error(0)
}

func (mock *Client) DeepPing() error {
	args := mock.Called()
	return args.Error(0)
}
func (mock *Client) DeepPingWithContext(ctx context.Context) error {
	args := mock.Called(ctx)
	return args.Error(0)
}

func (mock *Client) IsAuthorized(userID, action string, resource *common.Origin) (bool, error) {
	args := mock.Called(userID, action, resource)
	return args.Bool(0), args.Error(1)
}
func (mock *Client) IsAuthorizedWithContext(ctx context.Context, userID, action string, resource *common.Origin) (bool, error) {
	args := mock.Called(ctx, userID, action, resource)
	return args.Bool(0), args.Error(1)
}

func (mock *Client) IsAuthorizedBulk(userID, action string, resources []common.Origin) ([]string, []bool, error) {
	args := mock.Called(userID, action, resources)
	return args.Get(0).([]string), args.Get(1).([]bool), args.Error(2)
}
func (mock *Client) IsAuthorizedBulkWithContext(ctx context.Context, userID, action string, resources []common.Origin) ([]string, []bool, error) {
	args := mock.Called(userID, action, resources)
	return args.Get(0).([]string), args.Get(1).([]bool), args.Error(2)
}

func (mock *Client) IsAuthorizedByEndpoint(api, method, endpoint, userID string) (bool, error) {
	args := mock.Called(api, method, endpoint, userID)
	return args.Bool(0), args.Error(1)
}
func (mock *Client) IsAuthorizedByEndpointWithContext(ctx context.Context, api, method, endpoint, userID string) (bool, error) {
	args := mock.Called(ctx, api, method, endpoint, userID)
	return args.Bool(0), args.Error(1)
}

func (mock *Client) GetResourcesWithActionsAccess(actions []string, resourceType string, resource *common.Origin) ([]common.Origin, error) {
	args := mock.Called(actions, resourceType, resource)
	return args.Get(0).([]common.Origin), args.Error(1)
}
func (mock *Client) GetResourcesWithActionsAccessWithContext(ctx context.Context, actions []string, resourceType string, resource *common.Origin) ([]common.Origin, error) {
	args := mock.Called(ctx, actions, resourceType, resource)
	return args.Get(0).([]common.Origin), args.Error(1)
}

func (mock *Client) GetResourcesByUserAction(userID, actionName, resourceType string) ([]common.Origin, error) {
	args := mock.Called(userID, actionName, resourceType)
	return args.Get(0).([]common.Origin), args.Error(1)
}
func (mock *Client) GetResourcesByUserActionWithContext(ctx context.Context, userID, actionName, resourceType string) ([]common.Origin, error) {
	args := mock.Called(ctx, userID, actionName, resourceType)
	return args.Get(0).([]common.Origin), args.Error(1)
}

func (mock *Client) GetResourcesByType(resourceType string) (resources []common.Origin, err error) {
	args := mock.Called(resourceType)
	return args.Get(0).([]common.Origin), args.Error(1)
}
func (mock *Client) GetResourcesByTypeWithContext(ctx context.Context, resourceType string) (resources []common.Origin, err error) {
	args := mock.Called(ctx, resourceType)
	return args.Get(0).([]common.Origin), args.Error(1)
}

func (mock *Client) GetResourceParents(resource common.Origin, parentOriginType string) (resources []common.Origin, err error) {
	args := mock.Called(resource, parentOriginType)
	return args.Get(0).([]common.Origin), args.Error(1)
}
func (mock *Client) GetResourceParentsWithContext(ctx context.Context, resource common.Origin, parentOriginType string) (resources []common.Origin, err error) {
	args := mock.Called(ctx, resource, parentOriginType)
	return args.Get(0).([]common.Origin), args.Error(1)
}

func (mock *Client) GetResourceChildren(resource common.Origin, childOriginType string) (resources []common.Origin, err error) {
	args := mock.Called(resource, childOriginType)
	return args.Get(0).([]common.Origin), args.Error(1)
}
func (mock *Client) GetResourceChildrenWithContext(ctx context.Context, resource common.Origin, childOriginType string) (resources []common.Origin, err error) {
	args := mock.Called(ctx, resource, childOriginType)
	return args.Get(0).([]common.Origin), args.Error(1)
}

func (mock *Client) AddResource(resource common.Origin) error {
	args := mock.Called(resource)
	return args.Error(0)
}
func (mock *Client) AddResourceWithContext(ctx context.Context, resource common.Origin) error {
	args := mock.Called(ctx, resource)
	return args.Error(0)
}

func (mock *Client) GetResource(id string, originType string) (common.Origin, error) {
	args := mock.Called(id)
	return args.Get(0).(common.Origin), args.Error(1)
}
func (mock *Client) GetResourceWithContext(ctx context.Context, id string, originType string) (common.Origin, error) {
	args := mock.Called(id, originType)
	return args.Get(0).(common.Origin), args.Error(1)
}

func (mock *Client) AddResourceRelation(resource common.Origin, parent common.Origin) error {
	args := mock.Called(resource, parent)
	return args.Error(0)
}
func (mock *Client) AddResourceRelationWithContext(ctx context.Context, resource common.Origin, parent common.Origin) error {
	args := mock.Called(ctx, resource, parent)
	return args.Error(0)
}

func (mock *Client) RemoveResourceRelation(resource common.Origin, parent common.Origin) error {
	args := mock.Called(resource, parent)
	return args.Error(0)
}
func (mock *Client) RemoveResourceRelationWithContext(ctx context.Context, resource common.Origin, parent common.Origin) error {
	args := mock.Called(ctx, resource, parent)
	return args.Error(0)
}

func (mock *Client) RemoveResource(resource common.Origin) error {
	args := mock.Called(resource)
	return args.Error(0)
}

func (mock *Client) RemoveResourceWithContext(ctx context.Context, resource common.Origin) error {
	args := mock.Called(ctx, resource)
	return args.Error(0)
}

func (mock *Client) ApplyUserAction(userID, action string, resource *common.Origin) error {
	args := mock.Called(userID, action, resource)
	return args.Error(0)
}
func (mock *Client) ApplyUserActionWithContext(ctx context.Context, userID, action string, resource *common.Origin) error {
	args := mock.Called(ctx, userID, action, resource)
	return args.Error(0)
}

func (mock *Client) RemoveUserAction(userID, action string, resource *common.Origin) error {
	args := mock.Called(userID, action, resource)
	return args.Error(0)
}
func (mock *Client) RemoveUserActionWithContext(ctx context.Context, userID, action string, resource *common.Origin) error {
	args := mock.Called(ctx, userID, action, resource)
	return args.Error(0)
}

func (mock *Client) GetResourcesByOriginAndType(resource common.Origin, resourceType string, depth int32) (resources []common.Origin, err error) {
	args := mock.Called(resource, resourceType, depth)
	return args.Get(0).([]common.Origin), args.Error(1)
}
func (mock *Client) GetResourcesByOriginAndTypeWithContext(ctx context.Context, resource common.Origin, resourceType string, depth int32) (resources []common.Origin, err error) {
	args := mock.Called(ctx, resource, resourceType, depth)
	return args.Get(0).([]common.Origin), args.Error(1)
}

func (mock *Client) GetUserIDsWithAccessToResource(resource common.Origin) (resources []string, err error) {
	args := mock.Called(resource)
	return args.Get(0).([]string), args.Error(1)
}

func (mock *Client) GetUserIDsWithAccessToResourceWithContext(ctx context.Context, resource common.Origin) (resources []string, err error) {
	args := mock.Called(ctx, resource)
	return args.Get(0).([]string), args.Error(1)
}

func (mock *Client) AddResources(resources []common.Origin) error {
	args := mock.Called(resources)
	return args.Error(0)
}

func (mock *Client) AddResourcesWithContext(ctx context.Context, resources []common.Origin) error {
	args := mock.Called(ctx, resources)
	return args.Error(0)
}

func (mock *Client) RemoveResources(resources []common.Origin) error {
	args := mock.Called(resources)
	return args.Error(0)
}

func (mock *Client) RemoveResourcesWithContext(ctx context.Context, resources []common.Origin) error {
	args := mock.Called(ctx, resources)
	return args.Error(0)
}

func (mock *Client) AddResourceRelations(resources grpcapi.AddResourceRelationsInput) error {
	args := mock.Called(resources)
	return args.Error(0)
}

func (mock *Client) AddResourceRelationsWithContext(ctx context.Context, resources grpcapi.AddResourceRelationsInput) error {
	args := mock.Called(ctx, resources)
	return args.Error(0)
}

func (mock *Client) RemoveResourceRelations(resources grpcapi.RemoveResourceRelationsInput) error {
	args := mock.Called(resources)
	return args.Error(0)
}

func (mock *Client) RemoveResourceRelationsWithContext(ctx context.Context, resources grpcapi.RemoveResourceRelationsInput) error {
	args := mock.Called(ctx, resources)
	return args.Error(0)
}

func (mock *Client) GetActionsByUserRole(userRole string) ([]grpcapi.Action, error) {
	args := mock.Called(userRole)
	return args.Get(0).([]grpcapi.Action), args.Error(1)
}
func (mock *Client) GetActionsByUserRoleWithContext(ctx context.Context, userRole string) ([]grpcapi.Action, error) {
	args := mock.Called(ctx, userRole)
	return args.Get(0).([]grpcapi.Action), args.Error(1)
}

func (mock *Client) GetResourcesAndActionsByUser(userID string) ([]grpcapi.ActionResource, error) {
	args := mock.Called(userID)
	return args.Get(0).([]grpcapi.ActionResource), args.Error(1)
}
func (mock *Client) GetResourcesAndActionsByUserWithContext(ctx context.Context, userID string) ([]grpcapi.ActionResource, error) {
	args := mock.Called(ctx, userID)
	return args.Get(0).([]grpcapi.ActionResource), args.Error(1)
}
func (mock *Client) AddAction(action grpcapi.Action) error {
	args := mock.Called(action)
	return args.Error(0)
}
func (mock *Client) AddActionWithContext(ctx context.Context, action grpcapi.Action) error {
	args := mock.Called(ctx, action)
	return args.Error(0)
}
func (mock *Client) RemoveAction(name string) error {
	args := mock.Called(name)
	return args.Error(0)
}
func (mock *Client) RemoveActionWithContext(ctx context.Context, name string) error {
	args := mock.Called(ctx, name)
	return args.Error(0)
}
func (mock *Client) GetAction(name string) (grpcapi.Action, error) {
	args := mock.Called(name)
	return args.Get(0).(grpcapi.Action), args.Error(1)
}
func (mock *Client) GetActionWithContext(ctx context.Context, name string) (grpcapi.Action, error) {
	args := mock.Called(ctx, name)
	return args.Get(0).(grpcapi.Action), args.Error(1)
}
func (mock *Client) GetAllActions() ([]grpcapi.Action, error) {
	args := mock.Called()
	return args.Get(0).([]grpcapi.Action), args.Error(1)
}
func (mock *Client) GetAllActionsWithContext(ctx context.Context) ([]grpcapi.Action, error) {
	args := mock.Called(ctx)
	return args.Get(0).([]grpcapi.Action), args.Error(1)
}

func (mock *Client) GetUserActions(userID string) ([]grpcapi.Action, error) {
	args := mock.Called(userID)
	return args.Get(0).([]grpcapi.Action), args.Error(1)
}
func (mock *Client) GetUserActionsWithContext(ctx context.Context, userID string) ([]grpcapi.Action, error) {
	args := mock.Called(ctx, userID)
	return args.Get(0).([]grpcapi.Action), args.Error(1)
}

func (mock *Client) AddUserRole(role grpcapi.UserRole) error {
	args := mock.Called(role)
	return args.Error(0)
}
func (mock *Client) AddUserRoleWithContext(ctx context.Context, role grpcapi.UserRole) error {
	args := mock.Called(ctx, role)
	return args.Error(0)
}

func (mock *Client) GetUserRole(roleName string) (grpcapi.UserRole, error) {
	args := mock.Called(roleName)
	return args.Get(0).(grpcapi.UserRole), args.Error(1)
}
func (mock *Client) GetUserRoleWithContext(ctx context.Context, roleName string) (grpcapi.UserRole, error) {
	args := mock.Called(ctx, roleName)
	return args.Get(0).(grpcapi.UserRole), args.Error(1)
}

func (mock *Client) RemoveUserRole(roleName string) error {
	args := mock.Called(roleName)
	return args.Error(0)
}

func (mock *Client) RemoveUserRoleWithContext(ctx context.Context, roleName string) error {
	args := mock.Called(ctx, roleName)
	return args.Error(0)
}
