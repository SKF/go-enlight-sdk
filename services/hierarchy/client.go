package hierarchy

import (
	"context"
	"time"

	"github.com/SKF/go-eventsource/eventsource"
	"google.golang.org/grpc"

	"github.com/SKF/go-enlight-sdk/services/hierarchy/grpcapi"
)

type HierarchyClient interface {
	Dial(host, port string, opts ...grpc.DialOption) error
	Close()
	DeepPing() error

	SaveNode(request grpcapi.SaveNodeInput) (string, error)
	GetNode(nodeID string) (grpcapi.Node, error)
	GetNodes(parentID string) ([]grpcapi.Node, error)
	DeleteNode(request grpcapi.DeleteNodeInput) error
	GetEvents(since int, limit *int32) ([]eventsource.Record, error)
	GetParentNode(nodeID string) (grpcapi.Node, error)
}

type client struct {
	conn *grpc.ClientConn
	api  grpcapi.HierarchyClient
}

func CreateClient() HierarchyClient {
	return &client{}
}

func (c *client) Dial(host, port string, opts ...grpc.DialOption) (err error) {
	conn, err := grpc.Dial(host+":"+port, opts...)
	if err != nil {
		return
	}

	c.conn = conn
	c.api = grpcapi.NewHierarchyClient(conn)
	return
}

func (c *client) Close() {
	c.conn.Close()
}

func (c *client) DeepPing() error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	_, err := c.api.DeepPing(ctx, &grpcapi.PrimitiveVoid{})
	return err
}
