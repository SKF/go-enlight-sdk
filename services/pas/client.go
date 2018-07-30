package pas

import (
	"context"
	"time"

	"github.com/SKF/go-enlight-sdk/services/pas/pasapi"

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

	SetPointAlarmThreshold(input pasapi.SetPointAlarmThresholdInput) error
	SetPointAlarmThresholdWithContext(ctx context.Context, input pasapi.SetPointAlarmThresholdInput) error

	GetPointAlarmThreshold(nodeID string) ([]pasapi.AlarmStatusInterval, error)
	GetPointAlarmThresholdWithContext(ctx context.Context, nodeID string) ([]pasapi.AlarmStatusInterval, error)

	SetPointAlarmStatus(input pasapi.SetPointAlarmStatusInput) error
	SetPointAlarmStatusWithContext(ctx context.Context, input pasapi.SetPointAlarmStatusInput) error

	GetPointAlarmStatus(input pasapi.GetPointAlarmStatusInput) (pasapi.AlarmStatus, error)
	GetPointAlarmStatusWithContext(ctx context.Context, input pasapi.GetPointAlarmStatusInput) (pasapi.AlarmStatus, error)

	GetPointAlarmStatusStream(dc chan<- pasapi.GetPointAlarmStatusStreamOutput) error
	GetPointAlarmStatusStreamWithContext(ctx context.Context, dc chan<- pasapi.GetPointAlarmStatusStreamOutput) error
}

// Client implements the PointAlarmStatusClient and holds the connection.
type Client struct {
	api  pasapi.PointAlarmStatusClient
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
	c.api = pasapi.NewPointAlarmStatusClient(conn)
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
	_, err = c.api.DeepPing(ctx, &pasapi.Void{})
	return
}
