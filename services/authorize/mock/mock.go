package mock

import (
	"context"

	grpcapi "github.com/SKF/proto/authorize"

	"github.com/SKF/proto/common"

	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc"

	"github.com/SKF/go-enlight-sdk/services/authorize"
)

type client struct {
	mock.Mock
}

func Create() *client { // nolint: golint
	return new(client)
}

var _ authorize.AuthorizeClient = &client{}

func (mock *client) Dial(host, port string, opts ...grpc.DialOption) error {
	args := mock.Called(host, port, opts)
	return args.Error(0)
}
func (mock *client) DialWithContext(ctx context.Context, host, port string, opts ...grpc.DialOption) error {
	args := mock.Called(ctx, host, port, opts)
	return args.Error(0)
}

func (mock *client) Close() {
	mock.Called()
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

func (mock *client) IsAuthorizedByEndpoint(api, method, endpoint, userID string) (bool, error) {
	args := mock.Called(api, method, endpoint, userID)
	return args.Bool(0), args.Error(1)
}
func (mock *client) IsAuthorizedByEndpointWithContext(ctx context.Context, api, method, endpoint, userID string) (bool, error) {
	args := mock.Called(ctx, api, method, endpoint, userID)
	return args.Bool(0), args.Error(1)
}

func (mock *client) GetResourcesByType(resourceType string) (resources []common.Origin, err error) {
	args := mock.Called(resourceType)
	return args.Get(0).([]common.Origin), args.Error(1)
}
func (mock *client) GetResourcesByTypeWithContext(ctx context.Context, resourceType string) (resources []common.Origin, err error) {
	args := mock.Called(ctx, resourceType)
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

func (mock *client) AddUserPermission(userID, action string, resource *common.Origin) error {
	args := mock.Called(userID, action, resource)
	return args.Error(0)
}
func (mock *client) AddUserPermissionWithContext(ctx context.Context, userID, action string, resource *common.Origin) error {
	args := mock.Called(ctx, userID, action, resource)
	return args.Error(0)
}

func (mock *client) RemoveUserPermission(userID, action string, resource *common.Origin) error {
	args := mock.Called(userID, action, resource)
	return args.Error(0)
}
func (mock *client) RemoveUserPermissionWithContext(ctx context.Context, userID, action string, resource *common.Origin) error {
	args := mock.Called(ctx, userID, action, resource)
	return args.Error(0)
}

func (mock *client) GetResourcesByOriginAndType(originID string, resourceType string) (resources []common.Origin, err error) {
	args := mock.Called(originID, resourceType)
	return args.Get(0).([]common.Origin), args.Error(1)
}
func (mock *client) GetResourcesByOriginAndTypeWithContext(ctx context.Context, originID string, resourceType string) (resources []common.Origin, err error) {
	args := mock.Called(ctx, originID, resourceType)
	return args.Get(0).([]common.Origin), args.Error(1)
}

func (mock *client) GetUserIDsWithAccessToResource(originID string) (resources []string, err error) {
	args := mock.Called(originID)
	return args.Get(0).([]string), args.Error(1)
}

func (mock *client) GetUserIDsWithAccessToResourceWithContext(ctx context.Context, originID string) (resources []string, err error) {
	args := mock.Called(ctx, originID)
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
