package mock

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go/aws/session"

	grpcapi "github.com/SKF/proto/v2/authorize"

	"github.com/SKF/proto/v2/common"

	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc"

	"github.com/SKF/go-enlight-sdk/v2/services/authorize"
)

type client struct {
	mock.Mock
}

func Create() *client { // nolint: golint
	return new(client)
}

var _ authorize.AuthorizeClient = &client{}

func (mock *client) SetRequestTimeout(d time.Duration) {
	mock.Called(d)
}

func (mock *client) Dial(host, port string, opts ...grpc.DialOption) error {
	args := mock.Called(host, port, opts)
	return args.Error(0)
}

func (mock *client) DialWithContext(ctx context.Context, host, port string, opts ...grpc.DialOption) error {
	args := mock.Called(ctx, host, port, opts)
	return args.Error(0)
}

func (mock *client) DialUsingCredentials(sess *session.Session, host, port, secretKey string, opts ...grpc.DialOption) error {
	args := mock.Called(sess, host, port, secretKey, opts)
	return args.Error(0)
}

func (mock *client) DialUsingCredentialsWithContext(ctx context.Context, sess *session.Session, host, port, secretKey string, opts ...grpc.DialOption) error {
	args := mock.Called(ctx, sess, host, port, secretKey, opts)
	return args.Error(0)
}

func (mock *client) Close() error {
	args := mock.Called()
	return args.Error(0)
}

func (mock *client) DeepPing() error {
	args := mock.Called()
	return args.Error(0)
}
func (mock *client) DeepPingWithContext(ctx context.Context) error {
	args := mock.Called(ctx)
	return args.Error(0)
}

func (mock *client) IsAuthorized(userID, action string, resource *common.Origin) (bool, error) {
	args := mock.Called(userID, action, resource)
	return args.Bool(0), args.Error(1)
}
func (mock *client) IsAuthorizedWithContext(ctx context.Context, userID, action string, resource *common.Origin) (bool, error) {
	args := mock.Called(ctx, userID, action, resource)
	return args.Bool(0), args.Error(1)
}

func (mock *client) IsAuthorizedBulk(userID, action string, resources []common.Origin) ([]string, []bool, error) {
	args := mock.Called(userID, action, resources)
	return args.Get(0).([]string), args.Get(1).([]bool), args.Error(2)
}
func (mock *client) IsAuthorizedBulkWithContext(ctx context.Context, userID, action string, resources []common.Origin) ([]string, []bool, error) {
	args := mock.Called(userID, action, resources)
	return args.Get(0).([]string), args.Get(1).([]bool), args.Error(2)
}
func (mock *client) IsAuthorizedBulkWithResources(ctx context.Context, userID, action string, resources []common.Origin) ([]common.Origin, []bool, error) {
	args := mock.Called(userID, action, resources)
	return args.Get(0).([]common.Origin), args.Get(1).([]bool), args.Error(2)
}

func (mock *client) IsAuthorizedByEndpoint(api, method, endpoint, userID string) (bool, error) {
	args := mock.Called(api, method, endpoint, userID)
	return args.Bool(0), args.Error(1)
}
func (mock *client) IsAuthorizedByEndpointWithContext(ctx context.Context, api, method, endpoint, userID string) (bool, error) {
	args := mock.Called(ctx, api, method, endpoint, userID)
	return args.Bool(0), args.Error(1)
}

func (mock *client) GetResourcesWithActionsAccess(actions []string, resourceType string, resource *common.Origin) ([]common.Origin, error) {
	args := mock.Called(actions, resourceType, resource)
	return args.Get(0).([]common.Origin), args.Error(1)
}
func (mock *client) GetResourcesWithActionsAccessWithContext(ctx context.Context, actions []string, resourceType string, resource *common.Origin) ([]common.Origin, error) {
	args := mock.Called(ctx, actions, resourceType, resource)
	return args.Get(0).([]common.Origin), args.Error(1)
}

func (mock *client) GetResourcesByUserAction(userID, actionName, resourceType string) ([]common.Origin, error) {
	args := mock.Called(userID, actionName, resourceType)
	return args.Get(0).([]common.Origin), args.Error(1)
}
func (mock *client) GetResourcesByUserActionWithContext(ctx context.Context, userID, actionName, resourceType string) ([]common.Origin, error) {
	args := mock.Called(ctx, userID, actionName, resourceType)
	return args.Get(0).([]common.Origin), args.Error(1)
}

func (mock *client) GetResourcesByType(resourceType string) (resources []common.Origin, err error) {
	args := mock.Called(resourceType)
	return args.Get(0).([]common.Origin), args.Error(1)
}
func (mock *client) GetResourcesByTypeWithContext(ctx context.Context, resourceType string) (resources []common.Origin, err error) {
	args := mock.Called(ctx, resourceType)
	return args.Get(0).([]common.Origin), args.Error(1)
}

func (mock *client) GetResourceParents(resource common.Origin, parentOriginType string) (resources []common.Origin, err error) {
	args := mock.Called(resource, parentOriginType)
	return args.Get(0).([]common.Origin), args.Error(1)
}
func (mock *client) GetResourceParentsWithContext(ctx context.Context, resource common.Origin, parentOriginType string) (resources []common.Origin, err error) {
	args := mock.Called(ctx, resource, parentOriginType)
	return args.Get(0).([]common.Origin), args.Error(1)
}

func (mock *client) GetResourceChildren(resource common.Origin, childOriginType string) (resources []common.Origin, err error) {
	args := mock.Called(resource, childOriginType)
	return args.Get(0).([]common.Origin), args.Error(1)
}
func (mock *client) GetResourceChildrenWithContext(ctx context.Context, resource common.Origin, childOriginType string) (resources []common.Origin, err error) {
	args := mock.Called(ctx, resource, childOriginType)
	return args.Get(0).([]common.Origin), args.Error(1)
}

func (mock *client) AddResource(resource common.Origin) error {
	args := mock.Called(resource)
	return args.Error(0)
}
func (mock *client) AddResourceWithContext(ctx context.Context, resource common.Origin) error {
	args := mock.Called(ctx, resource)
	return args.Error(0)
}

func (mock *client) GetResource(id string, originType string) (common.Origin, error) {
	args := mock.Called(id)
	return args.Get(0).(common.Origin), args.Error(1)
}
func (mock *client) GetResourceWithContext(ctx context.Context, id string, originType string) (common.Origin, error) {
	args := mock.Called(id, originType)
	return args.Get(0).(common.Origin), args.Error(1)
}

func (mock *client) AddResourceRelation(resource common.Origin, parent common.Origin) error {
	args := mock.Called(resource, parent)
	return args.Error(0)
}
func (mock *client) AddResourceRelationWithContext(ctx context.Context, resource common.Origin, parent common.Origin) error {
	args := mock.Called(ctx, resource, parent)
	return args.Error(0)
}

func (mock *client) RemoveResourceRelation(resource common.Origin, parent common.Origin) error {
	args := mock.Called(resource, parent)
	return args.Error(0)
}
func (mock *client) RemoveResourceRelationWithContext(ctx context.Context, resource common.Origin, parent common.Origin) error {
	args := mock.Called(ctx, resource, parent)
	return args.Error(0)
}

func (mock *client) RemoveResource(resource common.Origin) error {
	args := mock.Called(resource)
	return args.Error(0)
}

func (mock *client) RemoveResourceWithContext(ctx context.Context, resource common.Origin) error {
	args := mock.Called(ctx, resource)
	return args.Error(0)
}

func (mock *client) ApplyUserAction(userID, action string, resource *common.Origin) error {
	args := mock.Called(userID, action, resource)
	return args.Error(0)
}
func (mock *client) ApplyUserActionWithContext(ctx context.Context, userID, action string, resource *common.Origin) error {
	args := mock.Called(ctx, userID, action, resource)
	return args.Error(0)
}

func (mock *client) RemoveUserAction(userID, action string, resource *common.Origin) error {
	args := mock.Called(userID, action, resource)
	return args.Error(0)
}
func (mock *client) RemoveUserActionWithContext(ctx context.Context, userID, action string, resource *common.Origin) error {
	args := mock.Called(ctx, userID, action, resource)
	return args.Error(0)
}

func (mock *client) GetResourcesByOriginAndType(resource common.Origin, resourceType string, depth int32) (resources []common.Origin, err error) {
	args := mock.Called(resource, resourceType, depth)
	return args.Get(0).([]common.Origin), args.Error(1)
}
func (mock *client) GetResourcesByOriginAndTypeWithContext(ctx context.Context, resource common.Origin, resourceType string, depth int32) (resources []common.Origin, err error) {
	args := mock.Called(ctx, resource, resourceType, depth)
	return args.Get(0).([]common.Origin), args.Error(1)
}

func (mock *client) GetUserIDsWithAccessToResource(resource common.Origin) (resources []string, err error) {
	args := mock.Called(resource)
	return args.Get(0).([]string), args.Error(1)
}

func (mock *client) GetUserIDsWithAccessToResourceWithContext(ctx context.Context, resource common.Origin) (resources []string, err error) {
	args := mock.Called(ctx, resource)
	return args.Get(0).([]string), args.Error(1)
}

func (mock *client) AddResources(resources []common.Origin) error {
	args := mock.Called(resources)
	return args.Error(0)
}

func (mock *client) AddResourcesWithContext(ctx context.Context, resources []common.Origin) error {
	args := mock.Called(ctx, resources)
	return args.Error(0)
}

func (mock *client) RemoveResources(resources []common.Origin) error {
	args := mock.Called(resources)
	return args.Error(0)
}

func (mock *client) RemoveResourcesWithContext(ctx context.Context, resources []common.Origin) error {
	args := mock.Called(ctx, resources)
	return args.Error(0)
}

func (mock *client) AddResourceRelations(resources grpcapi.AddResourceRelationsInput) error {
	args := mock.Called(resources)
	return args.Error(0)
}

func (mock *client) AddResourceRelationsWithContext(ctx context.Context, resources grpcapi.AddResourceRelationsInput) error {
	args := mock.Called(ctx, resources)
	return args.Error(0)
}

func (mock *client) RemoveResourceRelations(resources grpcapi.RemoveResourceRelationsInput) error {
	args := mock.Called(resources)
	return args.Error(0)
}

func (mock *client) RemoveResourceRelationsWithContext(ctx context.Context, resources grpcapi.RemoveResourceRelationsInput) error {
	args := mock.Called(ctx, resources)
	return args.Error(0)
}

func (mock *client) GetActionsByUserRole(userRole string) ([]grpcapi.Action, error) {
	args := mock.Called(userRole)
	return args.Get(0).([]grpcapi.Action), args.Error(1)
}
func (mock *client) GetActionsByUserRoleWithContext(ctx context.Context, userRole string) ([]grpcapi.Action, error) {
	args := mock.Called(ctx, userRole)
	return args.Get(0).([]grpcapi.Action), args.Error(1)
}

func (mock *client) GetResourcesAndActionsByUser(userID string) ([]grpcapi.ActionResource, error) {
	args := mock.Called(userID)
	return args.Get(0).([]grpcapi.ActionResource), args.Error(1)
}
func (mock *client) GetResourcesAndActionsByUserWithContext(ctx context.Context, userID string) ([]grpcapi.ActionResource, error) {
	args := mock.Called(ctx, userID)
	return args.Get(0).([]grpcapi.ActionResource), args.Error(1)
}

func (mock *client) GetResourcesAndActionsByUserAndResource(userID string, resource *common.Origin) ([]grpcapi.ActionResource, error) {
	args := mock.Called(userID, resource)
	return args.Get(0).([]grpcapi.ActionResource), args.Error(1)
}

func (mock *client) GetResourcesAndActionsByUserAndResourceWithContext(ctx context.Context, userID string, resource *common.Origin) ([]grpcapi.ActionResource, error) {
	args := mock.Called(ctx, userID, resource)
	return args.Get(0).([]grpcapi.ActionResource), args.Error(1)
}

func (mock *client) AddAction(action grpcapi.Action) error {
	args := mock.Called(action)
	return args.Error(0)
}
func (mock *client) AddActionWithContext(ctx context.Context, action grpcapi.Action) error {
	args := mock.Called(ctx, action)
	return args.Error(0)
}
func (mock *client) RemoveAction(name string) error {
	args := mock.Called(name)
	return args.Error(0)
}
func (mock *client) RemoveActionWithContext(ctx context.Context, name string) error {
	args := mock.Called(ctx, name)
	return args.Error(0)
}
func (mock *client) GetAction(name string) (grpcapi.Action, error) {
	args := mock.Called(name)
	return args.Get(0).(grpcapi.Action), args.Error(1)
}
func (mock *client) GetActionWithContext(ctx context.Context, name string) (grpcapi.Action, error) {
	args := mock.Called(ctx, name)
	return args.Get(0).(grpcapi.Action), args.Error(1)
}
func (mock *client) GetAllActions() ([]grpcapi.Action, error) {
	args := mock.Called()
	return args.Get(0).([]grpcapi.Action), args.Error(1)
}
func (mock *client) GetAllActionsWithContext(ctx context.Context) ([]grpcapi.Action, error) {
	args := mock.Called(ctx)
	return args.Get(0).([]grpcapi.Action), args.Error(1)
}

func (mock *client) GetUserActions(userID string) ([]grpcapi.Action, error) {
	args := mock.Called(userID)
	return args.Get(0).([]grpcapi.Action), args.Error(1)
}
func (mock *client) GetUserActionsWithContext(ctx context.Context, userID string) ([]grpcapi.Action, error) {
	args := mock.Called(ctx, userID)
	return args.Get(0).([]grpcapi.Action), args.Error(1)
}

func (mock *client) AddUserRole(role grpcapi.UserRole) error {
	args := mock.Called(role)
	return args.Error(0)
}
func (mock *client) AddUserRoleWithContext(ctx context.Context, role grpcapi.UserRole) error {
	args := mock.Called(ctx, role)
	return args.Error(0)
}

func (mock *client) GetUserRole(roleName string) (grpcapi.UserRole, error) {
	args := mock.Called(roleName)
	return args.Get(0).(grpcapi.UserRole), args.Error(1)
}
func (mock *client) GetUserRoleWithContext(ctx context.Context, roleName string) (grpcapi.UserRole, error) {
	args := mock.Called(ctx, roleName)
	return args.Get(0).(grpcapi.UserRole), args.Error(1)
}

func (mock *client) RemoveUserRole(roleName string) error {
	args := mock.Called(roleName)
	return args.Error(0)
}

func (mock *client) RemoveUserRoleWithContext(ctx context.Context, roleName string) error {
	args := mock.Called(ctx, roleName)
	return args.Error(0)
}
