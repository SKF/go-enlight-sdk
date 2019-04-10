package mock

import (
	"context"

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

func (mock *client) GetAuthorizationResourcesByType(resourceType string) (resources []common.Origin, err error) {
	args := mock.Called(resourceType)
	return args.Get(0).([]common.Origin), args.Error(1)
}
func (mock *client) GetAuthorizationResourcesByTypeWithContext(ctx context.Context, resourceType string) (resources []common.Origin, err error) {
	args := mock.Called(ctx, resourceType)
	return args.Get(0).([]common.Origin), args.Error(1)
}

func (mock *client) GetAuthorizationResourceRelations(resource common.Origin) (resources []common.Origin, err error) {
	args := mock.Called(resource)
	return args.Get(0).([]common.Origin), args.Error(1)
}
func (mock *client) GetAuthorizationResourceRelationsWithContext(ctx context.Context, resource common.Origin) (resources []common.Origin, err error) {
	args := mock.Called(ctx, resource)
	return args.Get(0).([]common.Origin), args.Error(1)
}

func (mock *client) AddAuthorizationResource(resource common.Origin) error {
	args := mock.Called(resource)
	return args.Error(0)
}
func (mock *client) AddAuthorizationResourceWithContext(ctx context.Context, resource common.Origin) error {
	args := mock.Called(ctx, resource)
	return args.Error(0)
}

func (mock *client) AddAuthorizationResourceRelation(resource common.Origin, parent common.Origin) error {
	args := mock.Called(resource, parent)
	return args.Error(0)
}
func (mock *client) AddAuthorizationResourceRelationWithContext(ctx context.Context, resource common.Origin, parent common.Origin) error {
	args := mock.Called(ctx, resource, parent)
	return args.Error(0)
}

func (mock *client) RemoveAuthorizationResourceRelation(resource common.Origin, parent common.Origin) error {
	args := mock.Called(resource, parent)
	return args.Error(0)
}
func (mock *client) RemoveAuthorizationResourceRelationWithContext(ctx context.Context, resource common.Origin, parent common.Origin) error {
	args := mock.Called(ctx, resource, parent)
	return args.Error(0)
}

func (mock *client) RemoveAuthorizationResource(resource common.Origin) error {
	args := mock.Called(resource)
	return args.Error(0)
}
func (mock *client) RemoveAuthorizationResourceWithContext(ctx context.Context, resource common.Origin) error {
	args := mock.Called(ctx, resource)
	return args.Error(0)
}

func (mock *client) AddUserPermission(userID, action string, resource common.Origin) error {
	args := mock.Called(userID, action, resource)
	return args.Error(0)
}
func (mock *client) AddUserPermissionWithContext(ctx context.Context, userID, action string, resource common.Origin) error {
	args := mock.Called(ctx, userID, action, resource)
	return args.Error(0)
}

func (mock *client) RemoveUserPermission(userID, role string, resource common.Origin) error {
	args := mock.Called(userID, role, resource)
	return args.Error(0)
}
func (mock *client) RemoveUserPermissionWithContext(ctx context.Context, userID, role string, resource common.Origin) error {
	args := mock.Called(ctx, userID, role, resource)
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
