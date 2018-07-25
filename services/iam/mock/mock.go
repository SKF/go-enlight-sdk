package mock

import (
	"github.com/SKF/go-eventsource/eventsource"
	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc"

	"github.com/SKF/go-enlight-sdk/services/iam"
	"github.com/SKF/go-enlight-sdk/services/iam/grpcapi"
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

func (mock *client) CheckAuthentication(token, method string) (grpcapi.User, error) {
	args := mock.Called(token, method)
	return args.Get(0).(grpcapi.User), args.Error(1)
}
func (mock *client) GetNodesByUser(userID string) (nodeIDs []string, err error) {
	args := mock.Called(userID)
	return args.Get(0).([]string), args.Error(1)
}
func (mock *client) GetEventRecords(since int, limit *int32) ([]eventsource.Record, error) {
	args := mock.Called(since, limit)
	return args.Get(0).([]eventsource.Record), args.Error(1)
}
