package hierarchy

import (
	"context"
	"time"

	"github.com/SKF/go-eventsource/eventsource"
	"google.golang.org/grpc"

	"github.com/SKF/go-enlight-sdk/services/hierarchy/grpcapi"
)

// HierarchyClient provides the API operation methods for making
// requests to Enlight Hierarchy Management Service. See this
// package's package overview docs for details on the service.
type HierarchyClient interface {
	Dial(host, port string, opts ...grpc.DialOption) error
	Close()

	DeepPing() error
	DeepPingWithContext(ctx context.Context) error

	SaveNode(request grpcapi.SaveNodeInput) (string, error)
	SaveNodeWithContext(ctx context.Context, request grpcapi.SaveNodeInput) (string, error)

	GetNode(nodeID string) (grpcapi.Node, error)
	GetNodeWithContext(ctx context.Context, nodeID string) (grpcapi.Node, error)

	GetNodes(parentID string) ([]grpcapi.Node, error)
	GetNodesWithContext(ctx context.Context, parentID string) ([]grpcapi.Node, error)

	GetChildNodes(parentID string) ([]grpcapi.Node, error)
	GetChildNodesWithContext(ctx context.Context, parentID string) ([]grpcapi.Node, error)

	DeleteNode(request grpcapi.DeleteNodeInput) error
	DeleteNodeWithContext(ctx context.Context, request grpcapi.DeleteNodeInput) error

	GetAncestors(nodeID string) ([]grpcapi.AncestorNode, error)
	GetAncestorsWithContext(ctx context.Context, nodeID string) ([]grpcapi.AncestorNode, error)

	GetEvents(since int, limit *int32) ([]eventsource.Record, error)
	GetEventsWithContext(ctx context.Context, since int, limit *int32) ([]eventsource.Record, error)

	GetParentNode(nodeID string) (grpcapi.Node, error)
	GetParentNodeWithContext(ctx context.Context, nodeID string) (grpcapi.Node, error)

	GetNodeIDByOrigin(origin grpcapi.Origin) (string, error)
	GetNodeIDByOriginWithContext(ctx context.Context, origin grpcapi.Origin) (string, error)
}

// Client implements the HierarchyClient and holds the connection.
type Client struct {
	conn *grpc.ClientConn
	api  grpcapi.HierarchyClient
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
	c.api = grpcapi.NewHierarchyClient(conn)
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
	_, err := c.api.DeepPing(ctx, &grpcapi.PrimitiveVoid{})
	return err
}
