package mock

import (
	"context"

	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc"

	"github.com/SKF/go-eventsource/eventsource"
	iam_grpcapi "github.com/SKF/proto/v2/iam"

	"github.com/SKF/go-enlight-sdk/v2/services/iam"
)

type client struct {
	mock.Mock
}

func Create() *client { // nolint: golint
	return new(client)
}

var _ iam.IAMClient = &client{}

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

func (mock *client) CheckAuthentication(token, arn string) (*iam_grpcapi.UserClaims, error) {
	args := mock.Called(token, arn)
	return args.Get(0).(*iam_grpcapi.UserClaims), args.Error(1)
}
func (mock *client) CheckAuthenticationWithContext(ctx context.Context, token, arn string) (*iam_grpcapi.UserClaims, error) {
	args := mock.Called(ctx, token, arn)
	return args.Get(0).(*iam_grpcapi.UserClaims), args.Error(1)
}

func (mock *client) CheckAuthenticationByEndpoint(token, api, method, endpoint string) (*iam_grpcapi.UserClaims, error) {
	args := mock.Called(token, api, method, endpoint)
	return args.Get(0).(*iam_grpcapi.UserClaims), args.Error(1)
}
func (mock *client) CheckAuthenticationByEndpointWithContext(ctx context.Context, token, api, method, endpoint string) (*iam_grpcapi.UserClaims, error) {
	args := mock.Called(ctx, token, api, method, endpoint)
	return args.Get(0).(*iam_grpcapi.UserClaims), args.Error(1)
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
