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

type ClientV2 struct {
	mock.Mock
}

func Create() *ClientV2 { // nolint: golint
	return new(ClientV2)
}

var _ authorize.AuthorizeClient = &ClientV2{}

func (mock *ClientV2) SetRequestTimeout(d time.Duration) {
	mock.Called(d)
}

func (mock *ClientV2) Dial(host, port string, opts ...grpc.DialOption) error {
	args := mock.Called(host, port, opts)
	return args.Error(0)
}

func (mock *ClientV2) DialWithContext(ctx context.Context, host, port string, opts ...grpc.DialOption) error {
	args := mock.Called(ctx, host, port, opts)
	return args.Error(0)
}

func (mock *ClientV2) DialUsingCredentials(sess *session.Session, host, port, secretKey string, opts ...grpc.DialOption) error {
	args := mock.Called(sess, host, port, secretKey, opts)
	return args.Error(0)
}

func (mock *ClientV2) DialUsingCredentialsWithContext(ctx context.Context, sess *session.Session, host, port, secretKey string, opts ...grpc.DialOption) error {
	args := mock.Called(ctx, sess, host, port, secretKey, opts)
	return args.Error(0)
}

func (mock *ClientV2) Close() error {
	args := mock.Called()
	return args.Error(0)
}

func (mock *ClientV2) DeepPing() error {
	args := mock.Called()
	return args.Error(0)
}
func (mock *ClientV2) DeepPingWithContext(ctx context.Context) error {
	args := mock.Called(ctx)
	return args.Error(0)
}

func (mock *ClientV2) IsAuthorized(userID, action string, resource *common.Origin) (bool, error) {
	args := mock.Called(userID, action, resource)
	return args.Bool(0), args.Error(1)
}
func (mock *ClientV2) IsAuthorizedWithContext(ctx context.Context, userID, action string, resource *common.Origin) (bool, error) {
	args := mock.Called(ctx, userID, action, resource)
	return args.Bool(0), args.Error(1)
}

func (mock *ClientV2) IsAuthorizedBulk(userID, action string, resources []common.Origin) ([]string, []bool, error) {
	args := mock.Called(userID, action, resources)
	return args.Get(0).([]string), args.Get(1).([]bool), args.Error(2)
}
func (mock *ClientV2) IsAuthorizedBulkWithContext(ctx context.Context, userID, action string, resources []common.Origin) ([]string, []bool, error) {
	args := mock.Called(userID, action, resources)
	return args.Get(0).([]string), args.Get(1).([]bool), args.Error(2)
}
func (mock *ClientV2) IsAuthorizedBulkWithResources(ctx context.Context, userID, action string, resources []common.Origin) ([]common.Origin, []bool, error) {
	args := mock.Called(userID, action, resources)
	return args.Get(0).([]common.Origin), args.Get(1).([]bool), args.Error(2)
}

func (mock *ClientV2) IsAuthorizedByEndpoint(api, method, endpoint, userID string) (bool, error) {
	args := mock.Called(api, method, endpoint, userID)
	return args.Bool(0), args.Error(1)
}
func (mock *ClientV2) IsAuthorizedByEndpointWithContext(ctx context.Context, api, method, endpoint, userID string) (bool, error) {
	args := mock.Called(ctx, api, method, endpoint, userID)
	return args.Bool(0), args.Error(1)
}

func (mock *ClientV2) GetResourcesWithActionsAccess(actions []string, resourceType string, resource *common.Origin) ([]common.Origin, error) {
	args := mock.Called(actions, resourceType, resource)
	return args.Get(0).([]common.Origin), args.Error(1)
}
func (mock *ClientV2) GetResourcesWithActionsAccessWithContext(ctx context.Context, actions []string, resourceType string, resource *common.Origin) ([]common.Origin, error) {
	args := mock.Called(ctx, actions, resourceType, resource)
	return args.Get(0).([]common.Origin), args.Error(1)
}

func (mock *ClientV2) GetResourcesByUserAction(userID, actionName, resourceType string) ([]common.Origin, error) {
	args := mock.Called(userID, actionName, resourceType)
	return args.Get(0).([]common.Origin), args.Error(1)
}
func (mock *ClientV2) GetResourcesByUserActionWithContext(ctx context.Context, userID, actionName, resourceType string) ([]common.Origin, error) {
	args := mock.Called(ctx, userID, actionName, resourceType)
	return args.Get(0).([]common.Origin), args.Error(1)
}

func (mock *ClientV2) GetResourcesByType(resourceType string) (resources []common.Origin, err error) {
	args := mock.Called(resourceType)
	return args.Get(0).([]common.Origin), args.Error(1)
}
func (mock *ClientV2) GetResourcesByTypeWithContext(ctx context.Context, resourceType string) (resources []common.Origin, err error) {
	args := mock.Called(ctx, resourceType)
	return args.Get(0).([]common.Origin), args.Error(1)
}

func (mock *ClientV2) GetResourceParents(resource common.Origin, parentOriginType string) (resources []common.Origin, err error) {
	args := mock.Called(resource, parentOriginType)
	return args.Get(0).([]common.Origin), args.Error(1)
}
func (mock *ClientV2) GetResourceParentsWithContext(ctx context.Context, resource common.Origin, parentOriginType string) (resources []common.Origin, err error) {
	args := mock.Called(ctx, resource, parentOriginType)
	return args.Get(0).([]common.Origin), args.Error(1)
}

func (mock *ClientV2) GetResourceChildren(resource common.Origin, childOriginType string) (resources []common.Origin, err error) {
	args := mock.Called(resource, childOriginType)
	return args.Get(0).([]common.Origin), args.Error(1)
}
func (mock *ClientV2) GetResourceChildrenWithContext(ctx context.Context, resource common.Origin, childOriginType string) (resources []common.Origin, err error) {
	args := mock.Called(ctx, resource, childOriginType)
	return args.Get(0).([]common.Origin), args.Error(1)
}

func (mock *ClientV2) AddResource(resource common.Origin) error {
	args := mock.Called(resource)
	return args.Error(0)
}
func (mock *ClientV2) AddResourceWithContext(ctx context.Context, resource common.Origin) error {
	args := mock.Called(ctx, resource)
	return args.Error(0)
}

func (mock *ClientV2) GetResource(id string, originType string) (common.Origin, error) {
	args := mock.Called(id)
	return args.Get(0).(common.Origin), args.Error(1)
}
func (mock *ClientV2) GetResourceWithContext(ctx context.Context, id string, originType string) (common.Origin, error) {
	args := mock.Called(id, originType)
	return args.Get(0).(common.Origin), args.Error(1)
}

func (mock *ClientV2) AddResourceRelation(resource common.Origin, parent common.Origin) error {
	args := mock.Called(resource, parent)
	return args.Error(0)
}
func (mock *ClientV2) AddResourceRelationWithContext(ctx context.Context, resource common.Origin, parent common.Origin) error {
	args := mock.Called(ctx, resource, parent)
	return args.Error(0)
}

func (mock *ClientV2) RemoveResourceRelation(resource common.Origin, parent common.Origin) error {
	args := mock.Called(resource, parent)
	return args.Error(0)
}
func (mock *ClientV2) RemoveResourceRelationWithContext(ctx context.Context, resource common.Origin, parent common.Origin) error {
	args := mock.Called(ctx, resource, parent)
	return args.Error(0)
}

func (mock *ClientV2) RemoveResource(resource common.Origin) error {
	args := mock.Called(resource)
	return args.Error(0)
}

func (mock *ClientV2) RemoveResourceWithContext(ctx context.Context, resource common.Origin) error {
	args := mock.Called(ctx, resource)
	return args.Error(0)
}

func (mock *ClientV2) ApplyUserAction(userID, action string, resource *common.Origin) error {
	args := mock.Called(userID, action, resource)
	return args.Error(0)
}
func (mock *ClientV2) ApplyUserActionWithContext(ctx context.Context, userID, action string, resource *common.Origin) error {
	args := mock.Called(ctx, userID, action, resource)
	return args.Error(0)
}

func (mock *ClientV2) ApplyRolesForUserOnResources(userID string, roles []string, resources []common.Origin) error {
	args := mock.Called(userID, roles, resources)
	return args.Error(0)
}

func (mock *ClientV2) ApplyRolesForUserOnResourcesWithContext(ctx context.Context, userID string, roles []string, resources []common.Origin) error {
	args := mock.Called(ctx, userID, roles, resources)
	return args.Error(0)
}

func (mock *ClientV2) RemoveUserAction(userID, action string, resource *common.Origin) error {
	args := mock.Called(userID, action, resource)
	return args.Error(0)
}
func (mock *ClientV2) RemoveUserActionWithContext(ctx context.Context, userID, action string, resource *common.Origin) error {
	args := mock.Called(ctx, userID, action, resource)
	return args.Error(0)
}

func (mock *ClientV2) GetResourcesByOriginAndType(resource common.Origin, resourceType string, depth int32) (resources []common.Origin, err error) {
	args := mock.Called(resource, resourceType, depth)
	return args.Get(0).([]common.Origin), args.Error(1)
}
func (mock *ClientV2) GetResourcesByOriginAndTypeWithContext(ctx context.Context, resource common.Origin, resourceType string, depth int32) (resources []common.Origin, err error) {
	args := mock.Called(ctx, resource, resourceType, depth)
	return args.Get(0).([]common.Origin), args.Error(1)
}

func (mock *ClientV2) GetUserIDsWithAccessToResource(resource common.Origin) (resources []string, err error) {
	args := mock.Called(resource)
	return args.Get(0).([]string), args.Error(1)
}

func (mock *ClientV2) GetUserIDsWithAccessToResourceWithContext(ctx context.Context, resource common.Origin) (resources []string, err error) {
	args := mock.Called(ctx, resource)
	return args.Get(0).([]string), args.Error(1)
}

func (mock *ClientV2) AddResources(resources []common.Origin) error {
	args := mock.Called(resources)
	return args.Error(0)
}

func (mock *ClientV2) AddResourcesWithContext(ctx context.Context, resources []common.Origin) error {
	args := mock.Called(ctx, resources)
	return args.Error(0)
}

func (mock *ClientV2) RemoveResources(resources []common.Origin) error {
	args := mock.Called(resources)
	return args.Error(0)
}

func (mock *ClientV2) RemoveResourcesWithContext(ctx context.Context, resources []common.Origin) error {
	args := mock.Called(ctx, resources)
	return args.Error(0)
}

func (mock *ClientV2) AddResourceRelations(resources grpcapi.AddResourceRelationsInput) error {
	args := mock.Called(resources)
	return args.Error(0)
}

func (mock *ClientV2) AddResourceRelationsWithContext(ctx context.Context, resources grpcapi.AddResourceRelationsInput) error {
	args := mock.Called(ctx, resources)
	return args.Error(0)
}

func (mock *ClientV2) RemoveResourceRelations(resources grpcapi.RemoveResourceRelationsInput) error {
	args := mock.Called(resources)
	return args.Error(0)
}

func (mock *ClientV2) RemoveResourceRelationsWithContext(ctx context.Context, resources grpcapi.RemoveResourceRelationsInput) error {
	args := mock.Called(ctx, resources)
	return args.Error(0)
}

func (mock *ClientV2) GetActionsByUserRole(userRole string) ([]grpcapi.Action, error) {
	args := mock.Called(userRole)
	return args.Get(0).([]grpcapi.Action), args.Error(1)
}
func (mock *ClientV2) GetActionsByUserRoleWithContext(ctx context.Context, userRole string) ([]grpcapi.Action, error) {
	args := mock.Called(ctx, userRole)
	return args.Get(0).([]grpcapi.Action), args.Error(1)
}

func (mock *ClientV2) GetResourcesAndActionsByUser(userID string) ([]grpcapi.ActionResource, error) {
	args := mock.Called(userID)
	return args.Get(0).([]grpcapi.ActionResource), args.Error(1)
}
func (mock *ClientV2) GetResourcesAndActionsByUserWithContext(ctx context.Context, userID string) ([]grpcapi.ActionResource, error) {
	args := mock.Called(ctx, userID)
	return args.Get(0).([]grpcapi.ActionResource), args.Error(1)
}

func (mock *ClientV2) GetResourcesAndActionsByUserAndResource(userID string, resource *common.Origin) ([]grpcapi.ActionResource, error) {
	args := mock.Called(userID, resource)
	return args.Get(0).([]grpcapi.ActionResource), args.Error(1)
}

func (mock *ClientV2) GetResourcesAndActionsByUserAndResourceWithContext(ctx context.Context, userID string, resource *common.Origin) ([]grpcapi.ActionResource, error) {
	args := mock.Called(ctx, userID, resource)
	return args.Get(0).([]grpcapi.ActionResource), args.Error(1)
}

func (mock *ClientV2) AddAction(action grpcapi.Action) error {
	args := mock.Called(action)
	return args.Error(0)
}
func (mock *ClientV2) AddActionWithContext(ctx context.Context, action grpcapi.Action) error {
	args := mock.Called(ctx, action)
	return args.Error(0)
}
func (mock *ClientV2) RemoveAction(name string) error {
	args := mock.Called(name)
	return args.Error(0)
}
func (mock *ClientV2) RemoveActionWithContext(ctx context.Context, name string) error {
	args := mock.Called(ctx, name)
	return args.Error(0)
}
func (mock *ClientV2) GetAction(name string) (grpcapi.Action, error) {
	args := mock.Called(name)
	return args.Get(0).(grpcapi.Action), args.Error(1)
}
func (mock *ClientV2) GetActionWithContext(ctx context.Context, name string) (grpcapi.Action, error) {
	args := mock.Called(ctx, name)
	return args.Get(0).(grpcapi.Action), args.Error(1)
}
func (mock *ClientV2) GetAllActions() ([]grpcapi.Action, error) {
	args := mock.Called()
	return args.Get(0).([]grpcapi.Action), args.Error(1)
}
func (mock *ClientV2) GetAllActionsWithContext(ctx context.Context) ([]grpcapi.Action, error) {
	args := mock.Called(ctx)
	return args.Get(0).([]grpcapi.Action), args.Error(1)
}

func (mock *ClientV2) GetUserActions(userID string) ([]grpcapi.Action, error) {
	args := mock.Called(userID)
	return args.Get(0).([]grpcapi.Action), args.Error(1)
}
func (mock *ClientV2) GetUserActionsWithContext(ctx context.Context, userID string) ([]grpcapi.Action, error) {
	args := mock.Called(ctx, userID)
	return args.Get(0).([]grpcapi.Action), args.Error(1)
}

func (mock *ClientV2) AddUserRole(role grpcapi.UserRole) error {
	args := mock.Called(role)
	return args.Error(0)
}
func (mock *ClientV2) AddUserRoleWithContext(ctx context.Context, role grpcapi.UserRole) error {
	args := mock.Called(ctx, role)
	return args.Error(0)
}

func (mock *ClientV2) GetUserRole(roleName string) (grpcapi.UserRole, error) {
	args := mock.Called(roleName)
	return args.Get(0).(grpcapi.UserRole), args.Error(1)
}
func (mock *ClientV2) GetUserRoleWithContext(ctx context.Context, roleName string) (grpcapi.UserRole, error) {
	args := mock.Called(ctx, roleName)
	return args.Get(0).(grpcapi.UserRole), args.Error(1)
}

func (mock *ClientV2) RemoveUserRole(roleName string) error {
	args := mock.Called(roleName)
	return args.Error(0)
}

func (mock *ClientV2) RemoveUserRoleWithContext(ctx context.Context, roleName string) error {
	args := mock.Called(ctx, roleName)
	return args.Error(0)
}

func (mock *ClientV2) IsAuthorizedWithReason(userID, action string, resource *common.Origin) (bool, string, error) {
	args := mock.Called(userID, action, resource)
	return args.Bool(0), args.String(1), args.Error(2)
}

func (mock *ClientV2) IsAuthorizedWithReasonWithContext(ctx context.Context, userID, action string, resource *common.Origin) (bool, string, error) {
	args := mock.Called(ctx, userID, action, resource)
	return args.Bool(0), args.String(1), args.Error(2)
}
