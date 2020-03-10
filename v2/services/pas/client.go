package pas

import (
	"context"
	"time"

	"github.com/SKF/proto/common"
	pas_api "github.com/SKF/proto/pas"

	"google.golang.org/grpc"
)

// PointAlarmStatusClient provides the API operation methods for making
// requests to the Enlight Point Alarm Status Service. See this
// package's package overview docs for details on the service.
type PointAlarmStatusClient interface {
	Dial(host, port string, opts ...grpc.DialOption) error
	DialWithContext(ctx context.Context, host, port string, opts ...grpc.DialOption) error
	Close()

	DeepPing() error
	DeepPingWithContext(ctx context.Context) error

	SetPointAlarmThreshold(input pas_api.SetPointAlarmThresholdInput) error
	SetPointAlarmThresholdWithContext(ctx context.Context, input pas_api.SetPointAlarmThresholdInput) error

	GetPointAlarmThreshold(nodeID string) (pas_api.GetPointAlarmThresholdOutput, error)
	GetPointAlarmThresholdWithContext(ctx context.Context, nodeID string) (pas_api.GetPointAlarmThresholdOutput, error)

	SetPointAlarmStatus(input pas_api.SetPointAlarmStatusInput) error
	SetPointAlarmStatusWithContext(ctx context.Context, input pas_api.SetPointAlarmStatusInput) error

	GetPointAlarmStatus(input pas_api.GetPointAlarmStatusInput) (pas_api.AlarmStatus, error)
	GetPointAlarmStatusWithContext(ctx context.Context, input pas_api.GetPointAlarmStatusInput) (pas_api.AlarmStatus, error)

	GetPointAlarmStatusEventLog(seqID string) (events pas_api.GetPointAlarmStatusEventLogOutput, err error)
	GetPointAlarmStatusEventLogWithContext(ctx context.Context, seqID string) (events pas_api.GetPointAlarmStatusEventLogOutput, err error)

	CalculateAndSetPointAlarmStatus(input pas_api.CalculateAndSetPointAlarmStatusInput) error
	CalculateAndSetPointAlarmStatusWithContext(ctx context.Context, input pas_api.CalculateAndSetPointAlarmStatusInput) error
}

// Client implements the PointAlarmStatusClient and holds the connection.
type Client struct {
	api  pas_api.PointAlarmStatusClient
	conn *grpc.ClientConn
}

// CreateClient creates an instance of the PointAlarmStatusClient
func CreateClient() PointAlarmStatusClient {
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
	c.api = pas_api.NewPointAlarmStatusClient(conn)
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
	_, err = c.api.DeepPing(ctx, &common.Void{})
	return
}
