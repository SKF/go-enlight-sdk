package mock

import (
	"context"

	"github.com/SKF/proto/common"

	"github.com/SKF/go-eventsource/eventsource"
	iam_grpcapi "github.com/SKF/proto/iam"
	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc"

	"github.com/SKF/go-enlight-sdk/services/iam"
)

type client struct {
	mock.Mock
}

func Create() *client {
	return new(client)
}

var _ iam.IAMClient = &client{}

func (mock *client) Dial(host, port string, opts ...grpc.DialOption) error {
	args := mock.Called(host, port, opts)
	return args.Error(0)
}
func (mock *client) Close() {
	mock.Called()
	return
}

func (mock *client) DeepPing() error {
	args := mock.Called()
	return args.Error(0)
}
func (mock *client) DeepPingWithContext(ctx context.Context) error {
	args := mock.Called(ctx)
	return args.Error(0)
}

func (mock *client) CheckAuthentication(token, method string) (iam_grpcapi.User, error) {
	args := mock.Called(token, method)
	return args.Get(0).(iam_grpcapi.User), args.Error(1)
}
func (mock *client) CheckAuthenticationWithContext(ctx context.Context, token, method string) (iam_grpcapi.User, error) {
	args := mock.Called(ctx, token, method)
	return args.Get(0).(iam_grpcapi.User), args.Error(1)
}

func (mock *client) GetNodesByUser(userID string) (nodeIDs []string, err error) {
	args := mock.Called(userID)
	return args.Get(0).([]string), args.Error(1)
}
func (mock *client) GetNodesByUserWithContext(ctx context.Context, userID string) (nodeIDs []string, err error) {
	args := mock.Called(ctx, userID)
	return args.Get(0).([]string), args.Error(1)
}

func (mock *client) GetEventRecords(since int, limit *int32) ([]eventsource.Record, error) {
	args := mock.Called(since, limit)
	return args.Get(0).([]eventsource.Record), args.Error(1)
}
func (mock *client) GetEventRecordsWithContext(ctx context.Context, since int, limit *int32) ([]eventsource.Record, error) {
	args := mock.Called(ctx, since, limit)
	return args.Get(0).([]eventsource.Record), args.Error(1)
}

func (mock *client) IsAuthorized(userID, action string, resource *common.Origin) error {
	args := mock.Called(userID, action, resource)
	return args.Error(0)
}
func (mock *client) IsAuthorizedWithContext(ctx context.Context, userID, action string, resource *common.Origin) error {
	args := mock.Called(ctx, userID, action, resource)
	return args.Error(0)
}

func (mock *client) AddResource(resource common.Origin) error {
	args := mock.Called(resource)
	return args.Error(0)
}
func (mock *client) AddResourceWithContext(ctx context.Context, resource common.Origin) error {
	args := mock.Called(ctx, resource)
	return args.Error(0)
}

func (mock *client) AddResourceParent(resource common.Origin, parent common.Origin) error {
	args := mock.Called(resource, parent)
	return args.Error(0)
}
func (mock *client) AddResourceParentWithContext(ctx context.Context, resource common.Origin, parent common.Origin) error {
	args := mock.Called(ctx, resource, parent)
	return args.Error(0)
}

func (mock *client) RemoveResourceParent(resource common.Origin, parent common.Origin) error {
	args := mock.Called(resource, parent)
	return args.Error(0)
}
func (mock *client) RemoveResourceParentWithContext(ctx context.Context, resource common.Origin, parent common.Origin) error {
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

func (mock *client) AddUserPermission(userID, action string, resource common.Origin) error {
	args := mock.Called(userID, action, resource)
	return args.Error(0)
}
func (mock *client) AddUserPermissionWithContext(ctx context.Context, userID, action string, resource common.Origin) error {
	args := mock.Called(ctx, userID, action, resource)
	return args.Error(0)
}

func (mock *client) RemoveUserRole(userID, role string, resource common.Origin) error {
	args := mock.Called(userID, role, resource)
	return args.Error(0)
}
func (mock *client) RemoveUserRoleWithContext(ctx context.Context, userID, role string, resource common.Origin) error {
	args := mock.Called(ctx, userID, role, resource)
	return args.Error(0)
}
