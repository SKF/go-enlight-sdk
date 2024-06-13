package pas

import (
	"context"
	"time"

	"github.com/SKF/go-utility/v2/log"
	"github.com/SKF/proto/v2/common"
	pas_api "github.com/SKF/proto/v2/pas"
	grpc_retry "github.com/grpc-ecosystem/go-grpc-middleware/retry"

	"github.com/SKF/go-enlight-sdk/v2/interceptors/reconnect"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
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

	SetPointAlarmThreshold(input *pas_api.SetPointAlarmThresholdInput) error
	SetPointAlarmThresholdWithContext(ctx context.Context, input *pas_api.SetPointAlarmThresholdInput) error

	GetPointAlarmThreshold(nodeID string) (*pas_api.GetPointAlarmThresholdOutput, error)
	GetPointAlarmThresholdWithContext(ctx context.Context, nodeID string) (*pas_api.GetPointAlarmThresholdOutput, error)

	SetPointAlarmStatus(input *pas_api.SetPointAlarmStatusInput) error
	SetPointAlarmStatusWithContext(ctx context.Context, input *pas_api.SetPointAlarmStatusInput) error

	// Deprecated: GetPointAlarmStatus only returns worst status
	// use GetPointAlarmStatusV2 instead
	GetPointAlarmStatus(input *pas_api.GetPointAlarmStatusInput) (pas_api.AlarmStatus, error)
	// Deprecated: GetPointAlarmStatusWithContext only returns worst status
	// use GetPointAlarmStatusV2WithContext instead
	GetPointAlarmStatusWithContext(ctx context.Context, input *pas_api.GetPointAlarmStatusInput) (pas_api.AlarmStatus, error)

	GetPointAlarmStatusV2(input *pas_api.GetPointAlarmStatusInput) (*pas_api.GetPointAlarmStatusOutput, error)
	GetPointAlarmStatusV2WithContext(ctx context.Context, input *pas_api.GetPointAlarmStatusInput) (*pas_api.GetPointAlarmStatusOutput, error)

	CalculateAndSetPointAlarmStatus(input *pas_api.CalculateAndSetPointAlarmStatusInput) error
	CalculateAndSetPointAlarmStatusWithContext(ctx context.Context, input *pas_api.CalculateAndSetPointAlarmStatusInput) error
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
	var newClientConn reconnect.NewConnectionFunc
	newClientConn = func(invokerCtx context.Context, invokerConn *grpc.ClientConn, invokerOptions ...grpc.CallOption) (context.Context, *grpc.ClientConn, []grpc.CallOption, error) {
		dialOptsReconnectRetry := reconnectRetryInterceptor(newClientConn)

		dialOpts := append(opts, dialOptsReconnectRetry, grpc.WithBlock())
		newConn, dialErr := grpc.DialContext(invokerCtx, host+":"+port, dialOpts...)
		if dialErr != nil {
			log.
				WithTracing(invokerCtx).
				WithError(dialErr).
				Error("Failed to connect to PAS gRPC server")
			return invokerCtx, invokerConn, invokerOptions, dialErr
		}
		_ = invokerConn.Close()

		c.conn = newConn
		c.api = pas_api.NewPointAlarmStatusClient(c.conn)
		return invokerCtx, c.conn, invokerOptions, err
	}

	dialOptsReconnectRetry := reconnectRetryInterceptor(newClientConn)
	dialOpts := append(opts, dialOptsReconnectRetry)

	conn, err := grpc.DialContext(ctx, host+":"+port, dialOpts...)
	if err != nil {
		log.
			WithTracing(ctx).
			WithError(err).
			Error("Failed to connect to PAS gRPC server")
		return err
	}

	c.conn = conn
	c.api = pas_api.NewPointAlarmStatusClient(c.conn)
	return
}

func reconnectRetryInterceptor(newClientConn reconnect.NewConnectionFunc) grpc.DialOption {
	retryIC := grpc_retry.UnaryClientInterceptor(
		grpc_retry.WithBackoff(grpc_retry.BackoffLinear(100*time.Millisecond)),
		grpc_retry.WithCodes(codes.Unavailable, codes.ResourceExhausted, codes.Aborted),
	)

	reconnectIC := reconnect.UnaryInterceptor(
		reconnect.WithNewConnection(newClientConn),
	)

	dialOptsReconnectRetry := grpc.WithChainUnaryInterceptor(reconnectIC, retryIC) // first one is outer, being called last
	return dialOptsReconnectRetry
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
