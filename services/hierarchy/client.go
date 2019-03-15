package hierarchy

import (
	"context"
	"time"

	"github.com/SKF/proto/common"

	"google.golang.org/grpc"

	"github.com/SKF/go-eventsource/eventsource"
	hierarchy_grpcapi "github.com/SKF/proto/hierarchy"
)

// HierarchyClient provides the API operation methods for making
// requests to Enlight Hierarchy Management Service. See this
// package's package overview docs for details on the service.
type HierarchyClient interface { // nolint: golint
	Dial(host, port string, opts ...grpc.DialOption) error
	DialWithContext(ctx context.Context, host, port string, opts ...grpc.DialOption) error
	Close()

	DeepPing() error
	DeepPingWithContext(ctx context.Context) error

	SaveNode(request hierarchy_grpcapi.SaveNodeInput) (string, error)
	SaveNodeWithContext(ctx context.Context, request hierarchy_grpcapi.SaveNodeInput) (string, error)

	GetNode(nodeID string) (hierarchy_grpcapi.Node, error)
	GetNodeWithContext(ctx context.Context, nodeID string) (hierarchy_grpcapi.Node, error)

	GetNodes(parentID string) ([]hierarchy_grpcapi.Node, error)
	GetNodesWithContext(ctx context.Context, parentID string) ([]hierarchy_grpcapi.Node, error)

	GetChildNodes(parentID string) ([]hierarchy_grpcapi.Node, error)
	GetChildNodesWithContext(ctx context.Context, parentID string) ([]hierarchy_grpcapi.Node, error)

	GetSubTree(parentID string, depth int, nodeTypes ...string) ([]hierarchy_grpcapi.Node, error)
	GetSubTreeWithContext(ctx context.Context, parentID string, depth int, nodeTypes ...string) ([]hierarchy_grpcapi.Node, error)

	DeleteNode(request hierarchy_grpcapi.DeleteNodeInput) error
	DeleteNodeWithContext(ctx context.Context, request hierarchy_grpcapi.DeleteNodeInput) error

	GetAncestors(nodeID string) ([]hierarchy_grpcapi.AncestorNode, error)
	GetAncestorsWithContext(ctx context.Context, nodeID string) ([]hierarchy_grpcapi.AncestorNode, error)

	GetEventsByCustomer(since int, limit *int32, customerID *string) ([]eventsource.Record, error)
	GetEventsByCustomerWithContext(ctx context.Context, since int, limit *int32, customerID *string) ([]eventsource.Record, error)

	GetEvents(since int, limit *int32) ([]eventsource.Record, error)
	GetEventsWithContext(ctx context.Context, since int, limit *int32) ([]eventsource.Record, error)

	GetParentNode(nodeID string) (hierarchy_grpcapi.Node, error)
	GetParentNodeWithContext(ctx context.Context, nodeID string) (hierarchy_grpcapi.Node, error)

	GetNodeIDByOrigin(origin common.Origin) (string, error)
	GetNodeIDByOriginWithContext(ctx context.Context, origin common.Origin) (string, error)
}

// Client implements the HierarchyClient and holds the connection.
type Client struct {
	conn *grpc.ClientConn
	api  hierarchy_grpcapi.HierarchyClient
}

// CreateClient creates an instance of the HierarchyClient.
func CreateClient() HierarchyClient {
	return &Client{}
}

// Dial creates a client connection to the given host with background context and no timeout
func (c *Client) Dial(host, port string, opts ...grpc.DialOption) error {
	return c.DialWithContext(context.Background(), host, port, opts...)
}

// DialWithContext creates a client connection to the given host with context (for timeout and transaction id)
func (c *Client) DialWithContext(ctx context.Context, host, port string, opts ...grpc.DialOption) (err error) {
	conn, err := grpc.DialContext(ctx, host+":"+port, opts...)
	if err != nil {
		return
	}

	c.conn = conn
	c.api = hierarchy_grpcapi.NewHierarchyClient(conn)
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
	_, err := c.api.DeepPing(ctx, &common.Void{})
	return err
}
