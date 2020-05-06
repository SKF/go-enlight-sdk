package mock

import (
	"context"

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

func (mock *client) Dial(ctx context.Context, host, port string, opts ...grpc.DialOption) error {
	args := mock.Called(ctx, host, port, opts)
	return args.Error(0)
}

func (mock *client) DialUsingCredentials(ctx context.Context, sess *session.Session, host, port, secretKey string, opts ...grpc.DialOption) error {
	args := mock.Called(ctx, sess, host, port, secretKey, opts)
	return args.Error(0)
}

func (mock *client) Close(ctx context.Context) error {
	args := mock.Called(ctx)
	return args.Error(0)
}

func (mock *client) DeepPing(ctx context.Context) error {
	args := mock.Called(ctx)
	return args.Error(0)
}

func (mock *client) IsAuthorized(ctx context.Context, userID, action string, resource *common.Origin) (bool, error) {
	args := mock.Called(ctx, userID, action, resource)
	return args.Bool(0), args.Error(1)
}

func (mock *client) IsAuthorizedBulk(ctx context.Context, userID, action string, resources []common.Origin) ([]string, []bool, error) {
	args := mock.Called(userID, action, resources)
	return args.Get(0).([]string), args.Get(1).([]bool), args.Error(2)
}
func (mock *client) IsAuthorizedBulkWithResources(ctx context.Context, userID, action string, resources []common.Origin) ([]common.Origin, []bool, error) {
	args := mock.Called(userID, action, resources)
	return args.Get(0).([]common.Origin), args.Get(1).([]bool), args.Error(2)
}

func (mock *client) IsAuthorizedByEndpoint(ctx context.Context, api, method, endpoint, userID string) (bool, error) {
	args := mock.Called(ctx, api, method, endpoint, userID)
	return args.Bool(0), args.Error(1)
}

func (mock *client) GetResourcesWithActionsAccess(ctx context.Context, actions []string, resourceType string, resource *common.Origin) ([]common.Origin, error) {
	args := mock.Called(ctx, actions, resourceType, resource)
	return args.Get(0).([]common.Origin), args.Error(1)
}

func (mock *client) GetResourcesByUserAction(ctx context.Context, userID, actionName, resourceType string) ([]common.Origin, error) {
	args := mock.Called(ctx, userID, actionName, resourceType)
	return args.Get(0).([]common.Origin), args.Error(1)
}

func (mock *client) GetResourcesByType(ctx context.Context, resourceType string) (resources []common.Origin, err error) {
	args := mock.Called(ctx, resourceType)
	return args.Get(0).([]common.Origin), args.Error(1)
}

func (mock *client) GetResourceParents(ctx context.Context, resource common.Origin, parentOriginType string) (resources []common.Origin, err error) {
	args := mock.Called(ctx, resource, parentOriginType)
	return args.Get(0).([]common.Origin), args.Error(1)
}

func (mock *client) GetResourceChildren(ctx context.Context, resource common.Origin, childOriginType string) (resources []common.Origin, err error) {
	args := mock.Called(ctx, resource, childOriginType)
	return args.Get(0).([]common.Origin), args.Error(1)
}

func (mock *client) AddResource(ctx context.Context, resource common.Origin) error {
	args := mock.Called(ctx, resource)
	return args.Error(0)
}

func (mock *client) GetResource(ctx context.Context, id string, originType string) (common.Origin, error) {
	args := mock.Called(id, originType)
	return args.Get(0).(common.Origin), args.Error(1)
}

func (mock *client) AddResourceRelation(ctx context.Context, resource common.Origin, parent common.Origin) error {
	args := mock.Called(ctx, resource, parent)
	return args.Error(0)
}

func (mock *client) RemoveResourceRelation(ctx context.Context, resource common.Origin, parent common.Origin) error {
	args := mock.Called(ctx, resource, parent)
	return args.Error(0)
}

func (mock *client) RemoveResource(ctx context.Context, resource common.Origin) error {
	args := mock.Called(ctx, resource)
	return args.Error(0)
}

func (mock *client) ApplyUserAction(ctx context.Context, userID, action string, resource *common.Origin) error {
	args := mock.Called(ctx, userID, action, resource)
	return args.Error(0)
}

func (mock *client) RemoveUserAction(ctx context.Context, userID, action string, resource *common.Origin) error {
	args := mock.Called(ctx, userID, action, resource)
	return args.Error(0)
}

func (mock *client) GetResourcesByOriginAndType(ctx context.Context, resource common.Origin, resourceType string, depth int32) (resources []common.Origin, err error) {
	args := mock.Called(ctx, resource, resourceType, depth)
	return args.Get(0).([]common.Origin), args.Error(1)
}

func (mock *client) GetUserIDsWithAccessToResource(ctx context.Context, resource common.Origin) (resources []string, err error) {
	args := mock.Called(ctx, resource)
	return args.Get(0).([]string), args.Error(1)
}

func (mock *client) AddResources(ctx context.Context, resources []common.Origin) error {
	args := mock.Called(ctx, resources)
	return args.Error(0)
}

func (mock *client) RemoveResources(ctx context.Context, resources []common.Origin) error {
	args := mock.Called(ctx, resources)
	return args.Error(0)
}

func (mock *client) AddResourceRelations(ctx context.Context, resources grpcapi.AddResourceRelationsInput) error {
	args := mock.Called(ctx, resources)
	return args.Error(0)
}

func (mock *client) RemoveResourceRelations(ctx context.Context, resources grpcapi.RemoveResourceRelationsInput) error {
	args := mock.Called(ctx, resources)
	return args.Error(0)
}

func (mock *client) GetActionsByUserRole(ctx context.Context, userRole string) ([]grpcapi.Action, error) {
	args := mock.Called(ctx, userRole)
	return args.Get(0).([]grpcapi.Action), args.Error(1)
}

func (mock *client) GetResourcesAndActionsByUser(ctx context.Context, userID string) ([]grpcapi.ActionResource, error) {
	args := mock.Called(ctx, userID)
	return args.Get(0).([]grpcapi.ActionResource), args.Error(1)
}

func (mock *client) GetResourcesAndActionsByUserAndResource(ctx context.Context, userID string, resource *common.Origin) ([]grpcapi.ActionResource, error) {
	args := mock.Called(ctx, userID, resource)
	return args.Get(0).([]grpcapi.ActionResource), args.Error(1)
}

func (mock *client) AddAction(ctx context.Context, action grpcapi.Action) error {
	args := mock.Called(ctx, action)
	return args.Error(0)
}

func (mock *client) RemoveAction(ctx context.Context, name string) error {
	args := mock.Called(ctx, name)
	return args.Error(0)
}

func (mock *client) GetAction(ctx context.Context, name string) (grpcapi.Action, error) {
	args := mock.Called(ctx, name)
	return args.Get(0).(grpcapi.Action), args.Error(1)
}

func (mock *client) GetAllActions(ctx context.Context) ([]grpcapi.Action, error) {
	args := mock.Called(ctx)
	return args.Get(0).([]grpcapi.Action), args.Error(1)
}

func (mock *client) GetUserActions(ctx context.Context, userID string) ([]grpcapi.Action, error) {
	args := mock.Called(ctx, userID)
	return args.Get(0).([]grpcapi.Action), args.Error(1)
}

func (mock *client) AddUserRole(ctx context.Context, role grpcapi.UserRole) error {
	args := mock.Called(ctx, role)
	return args.Error(0)
}

func (mock *client) GetUserRole(ctx context.Context, roleName string) (grpcapi.UserRole, error) {
	args := mock.Called(ctx, roleName)
	return args.Get(0).(grpcapi.UserRole), args.Error(1)
}

func (mock *client) RemoveUserRole(ctx context.Context, roleName string) error {
	args := mock.Called(ctx, roleName)
	return args.Error(0)
}
