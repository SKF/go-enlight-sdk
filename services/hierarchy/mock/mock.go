package mock

import (
	"context"

	"github.com/SKF/go-eventsource/eventsource"
	proto_common "github.com/SKF/proto/common"
	proto_hierarchy "github.com/SKF/proto/hierarchy"
	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc"

	"github.com/SKF/go-enlight-sdk/services/hierarchy"
)

type client struct {
	mock.Mock
}

func Create() *client {
	return new(client)
}

var _ hierarchy.HierarchyClient = &client{}

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

func (mock *client) SaveNode(request proto_hierarchy.SaveNodeInput) (string, error) {
	args := mock.Called(request)
	return args.String(0), args.Error(1)
}
func (mock *client) SaveNodeWithContext(ctx context.Context, request proto_hierarchy.SaveNodeInput) (string, error) {
	args := mock.Called(ctx, request)
	return args.String(0), args.Error(1)
}

func (mock *client) GetNode(uuid string) (proto_hierarchy.Node, error) {
	args := mock.Called(uuid)
	return args.Get(0).(proto_hierarchy.Node), args.Error(1)
}
func (mock *client) GetNodeWithContext(ctx context.Context, uuid string) (proto_hierarchy.Node, error) {
	args := mock.Called(ctx, uuid)
	return args.Get(0).(proto_hierarchy.Node), args.Error(1)
}

func (mock *client) GetNodeIDByOrigin(origin proto_common.Origin) (string, error) {
	args := mock.Called(origin)
	return args.String(0), args.Error(1)
}
func (mock *client) GetNodeIDByOriginWithContext(ctx context.Context, origin proto_common.Origin) (string, error) {
	args := mock.Called(ctx, origin)
	return args.String(0), args.Error(1)
}

func (mock *client) GetNodes(parentID string) ([]proto_hierarchy.Node, error) {
	args := mock.Called(parentID)
	return args.Get(0).([]proto_hierarchy.Node), args.Error(1)
}
func (mock *client) GetNodesWithContext(ctx context.Context, parentID string) ([]proto_hierarchy.Node, error) {
	args := mock.Called(ctx, parentID)
	return args.Get(0).([]proto_hierarchy.Node), args.Error(1)
}

func (mock *client) DeleteNode(request proto_hierarchy.DeleteNodeInput) error {
	args := mock.Called(request)
	return args.Error(0)
}
func (mock *client) DeleteNodeWithContext(ctx context.Context, request proto_hierarchy.DeleteNodeInput) error {
	args := mock.Called(ctx, request)
	return args.Error(0)
}

func (mock *client) GetEvents(since int, limit *int32) ([]eventsource.Record, error) {
	args := mock.Called(since, limit)
	return args.Get(0).([]eventsource.Record), args.Error(1)
}
func (mock *client) GetEventsWithContext(ctx context.Context, since int, limit *int32) ([]eventsource.Record, error) {
	args := mock.Called(ctx, since, limit)
	return args.Get(0).([]eventsource.Record), args.Error(1)
}

func (mock *client) GetParentNode(nodeID string) (proto_hierarchy.Node, error) {
	args := mock.Called(nodeID)
	return args.Get(0).(proto_hierarchy.Node), args.Error(1)
}
func (mock *client) GetParentNodeWithContext(ctx context.Context, nodeID string) (proto_hierarchy.Node, error) {
	args := mock.Called(ctx, nodeID)
	return args.Get(0).(proto_hierarchy.Node), args.Error(1)
}

func (mock *client) GetAncestors(nodeID string) ([]proto_hierarchy.AncestorNode, error) {
	args := mock.Called(nodeID)
	return args.Get(0).([]proto_hierarchy.AncestorNode), args.Error(1)
}
func (mock *client) GetAncestorsWithContext(ctx context.Context, nodeID string) ([]proto_hierarchy.AncestorNode, error) {
	args := mock.Called(ctx, nodeID)
	return args.Get(0).([]proto_hierarchy.AncestorNode), args.Error(1)
}

func (mock *client) GetChildNodes(parentID string) ([]proto_hierarchy.Node, error) {
	args := mock.Called(parentID)
	return args.Get(0).([]proto_hierarchy.Node), args.Error(1)
}
func (mock *client) GetChildNodesWithContext(ctx context.Context, parentID string) ([]proto_hierarchy.Node, error) {
	args := mock.Called(ctx, parentID)
	return args.Get(0).([]proto_hierarchy.Node), args.Error(1)
}
