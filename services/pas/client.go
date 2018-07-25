package pas

import (
	"context"
	"time"

	"github.com/SKF/go-enlight-sdk/services/pas/pasapi"

	"google.golang.org/grpc"
)

type PointAlarmStatusClient interface {
	Dial(host, port string, opts ...grpc.DialOption) error
	Close()

	DeepPing() error
	DeepPingWithContext(ctx context.Context) error

	SetPointThreshold(input pasapi.SetPointThresholdInput) error
	SetPointThresholdWithContext(ctx context.Context, input pasapi.SetPointThresholdInput) error

	GetPointThreshold(nodeID string) ([]pasapi.AlarmStatusInterval, error)
	GetPointThresholdWithContext(ctx context.Context, nodeID string) ([]pasapi.AlarmStatusInterval, error)

	SetPointStatus(input pasapi.SetPointStatusInput) error
	SetPointStatusWithContext(ctx context.Context, input pasapi.SetPointStatusInput) error

	GetPointStatus(input pasapi.GetPointStatusInput) (pasapi.AlarmStatus, error)
	GetPointStatusWithContext(ctx context.Context, input pasapi.GetPointStatusInput) (pasapi.AlarmStatus, error)

	GetPointStatusStream(dc chan<- pasapi.GetPointStatusStreamOutput) error
	GetPointStatusStreamWithContext(ctx context.Context, dc chan<- pasapi.GetPointStatusStreamOutput) error
}

func CreateClient() PointAlarmStatusClient {
	return &client{}
}

type client struct {
	api  pasapi.PointAlarmStatusClient
	conn *grpc.ClientConn
}

func (c *client) Dial(host, port string, opts ...grpc.DialOption) (err error) {
	conn, err := grpc.Dial(host+":"+port, opts...)
	if err != nil {
		return
	}

	c.conn = conn
	c.api = pasapi.NewPointAlarmStatusClient(conn)
	return
}

func (c *client) Close() {
	c.conn.Close()
}

func (c *client) DeepPing() (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	return c.DeepPingWithContext(ctx)
}

func (c *client) DeepPingWithContext(ctx context.Context) (err error) {
	_, err = c.api.DeepPing(ctx, &pasapi.Void{})
	return
}
