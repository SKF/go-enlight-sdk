package pas

import (
	"context"
	"time"

	proto_common "github.com/SKF/proto/common"
	proto_pas "github.com/SKF/proto/pas"
	"google.golang.org/grpc"
)

// PointAlarmStatusClient provides the API operation methods for making
// requests to the Enlight Point Alarm Status Service. See this
// package's package overview docs for details on the service.
type PointAlarmStatusClient interface {
	Dial(host, port string, opts ...grpc.DialOption) error
	Close()

	DeepPing() error
	DeepPingWithContext(ctx context.Context) error

	SetPointAlarmThreshold(input proto_pas.SetPointAlarmThresholdInput) error
	SetPointAlarmThresholdWithContext(ctx context.Context, input proto_pas.SetPointAlarmThresholdInput) error

	GetPointAlarmThreshold(nodeID string) (proto_pas.GetPointAlarmThresholdOutput, error)
	GetPointAlarmThresholdWithContext(ctx context.Context, nodeID string) (proto_pas.GetPointAlarmThresholdOutput, error)

	SetPointAlarmStatus(input proto_pas.SetPointAlarmStatusInput) error
	SetPointAlarmStatusWithContext(ctx context.Context, input proto_pas.SetPointAlarmStatusInput) error

	GetPointAlarmStatus(input proto_pas.GetPointAlarmStatusInput) (proto_pas.AlarmStatus, error)
	GetPointAlarmStatusWithContext(ctx context.Context, input proto_pas.GetPointAlarmStatusInput) (proto_pas.AlarmStatus, error)

	GetPointAlarmStatusStream(dc chan<- proto_pas.GetPointAlarmStatusStreamOutput) error
	GetPointAlarmStatusStreamWithContext(ctx context.Context, dc chan<- proto_pas.GetPointAlarmStatusStreamOutput) error
}

// Client implements the PointAlarmStatusClient and holds the connection.
type Client struct {
	api  proto_pas.PointAlarmStatusClient
	conn *grpc.ClientConn
}

// CreateClient creates an instance of the PointAlarmStatusClient
func CreateClient() PointAlarmStatusClient {
	return &Client{}
}

// Dial creates a client connection to the given host.
func (c *Client) Dial(host, port string, opts ...grpc.DialOption) (err error) {
	conn, err := grpc.Dial(host+":"+port, opts...)
	if err != nil {
		return
	}

	c.conn = conn
	c.api = proto_pas.NewPointAlarmStatusClient(conn)
	return
}

// Close tears down the ClientConn and all underlying connections.
func (c *Client) Close() {
	c.conn.Close()
}

// DeepPing pings the service to see if it is alive.
func (c *Client) DeepPing() (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	return c.DeepPingWithContext(ctx)
}

// DeepPingWithContext pings the service to see if it is alive.
func (c *Client) DeepPingWithContext(ctx context.Context) (err error) {
	_, err = c.api.DeepPing(ctx, &proto_common.Void{})
	return
}
