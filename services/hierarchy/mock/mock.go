package mock

import (
	"github.com/SKF/go-eventsource/eventsource"
	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc"

	"github.com/SKF/go-enlight-sdk/services/hierarchy"
	"github.com/SKF/go-enlight-sdk/services/hierarchy/grpcapi"
)

type client struct {
	mock.Mock
	hierarchy.HierarchyClient
}

// Create returns an empty mock
func Create() *client {
	return new(client)
}

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

func (mock *client) SaveNode(request grpcapi.SaveNodeInput) (string, error) {
	args := mock.Called(request)
	return args.String(0), args.Error(1)
}
func (mock *client) GetNode(uuid string) (grpcapi.Node, error) {
	args := mock.Called(uuid)
	return args.Get(0).(grpcapi.Node), args.Error(1)
}
func (mock *client) GetNodes(parentID string) ([]grpcapi.Node, error) {
	args := mock.Called(parentID)
	return args.Get(0).([]grpcapi.Node), args.Error(1)
}
func (mock *client) DeleteNode(request grpcapi.DeleteNodeInput) error {
	args := mock.Called(request)
	return args.Error(0)
}
func (mock *client) GetEvents(since int, limit *int32) ([]eventsource.Record, error) {
	args := mock.Called(since, limit)
	return args.Get(0).([]eventsource.Record), args.Error(1)
}

func (mock *client) GetParentNode(nodeID string) (grpcapi.Node, error) {
	args := mock.Called(nodeID)
	return args.String(0), args.Error(1)
}
