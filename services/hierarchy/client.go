package hierarchy

import (
	"context"
	"time"

	"github.com/SKF/go-eventsource/eventsource"
	proto_common "github.com/SKF/proto/common"
	proto_hierarchy "github.com/SKF/proto/hierarchy"
	"google.golang.org/grpc"
)

// HierarchyClient provides the API operation methods for making
// requests to Enlight Hierarchy Management Service. See this
// package's package overview docs for details on the service.
type HierarchyClient interface {
	Dial(host, port string, opts ...grpc.DialOption) error
	Close()

	DeepPing() error
	DeepPingWithContext(ctx context.Context) error

	SaveNode(request proto_hierarchy.SaveNodeInput) (string, error)
	SaveNodeWithContext(ctx context.Context, request proto_hierarchy.SaveNodeInput) (string, error)

	GetNode(nodeID string) (proto_hierarchy.Node, error)
	GetNodeWithContext(ctx context.Context, nodeID string) (proto_hierarchy.Node, error)

	GetNodes(parentID string) ([]proto_hierarchy.Node, error)
	GetNodesWithContext(ctx context.Context, parentID string) ([]proto_hierarchy.Node, error)

	GetChildNodes(parentID string) ([]proto_hierarchy.Node, error)
	GetChildNodesWithContext(ctx context.Context, parentID string) ([]proto_hierarchy.Node, error)

	DeleteNode(request proto_hierarchy.DeleteNodeInput) error
	DeleteNodeWithContext(ctx context.Context, request proto_hierarchy.DeleteNodeInput) error

	GetAncestors(nodeID string) ([]proto_hierarchy.AncestorNode, error)
	GetAncestorsWithContext(ctx context.Context, nodeID string) ([]proto_hierarchy.AncestorNode, error)

	GetEvents(since int, limit *int32) ([]eventsource.Record, error)
	GetEventsWithContext(ctx context.Context, since int, limit *int32) ([]eventsource.Record, error)

	GetParentNode(nodeID string) (proto_hierarchy.Node, error)
	GetParentNodeWithContext(ctx context.Context, nodeID string) (proto_hierarchy.Node, error)

	GetNodeIDByOrigin(origin proto_common.Origin) (string, error)
	GetNodeIDByOriginWithContext(ctx context.Context, origin proto_common.Origin) (string, error)
}

// Client implements the HierarchyClient and holds the connection.
type Client struct {
	conn *grpc.ClientConn
	api  proto_hierarchy.HierarchyClient
}

// CreateClient creates an instance of the HierarchyClient.
func CreateClient() HierarchyClient {
	return &Client{}
}

// Dial creates a client connection to the given host.
func (c *Client) Dial(host, port string, opts ...grpc.DialOption) (err error) {
	conn, err := grpc.Dial(host+":"+port, opts...)
	if err != nil {
		return
	}

	c.conn = conn
	c.api = proto_hierarchy.NewHierarchyClient(conn)
	return
}

// Close tears down the ClientConn and all underlying connections.
func (c *Client) Close() {
	c.conn.Close()
}

// DeepPing pings the service to see if it is alive.
func (c *Client) DeepPing() error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	return c.DeepPingWithContext(ctx)
}

// DeepPingWithContext pings the service to see if it is alive.
func (c *Client) DeepPingWithContext(ctx context.Context) error {
	_, err := c.api.DeepPing(ctx, &proto_common.Void{})
	return err
}
