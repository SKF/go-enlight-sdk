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

	SetPointAlarmThresholds(input pasapi.SetPointAlarmThresholdsInput) error
	SetPointAlarmThresholdsWithContext(ctx context.Context, input pasapi.SetPointAlarmThresholdsInput) error

	GetPointAlarmThresholds(nodeID string) ([]pasapi.AlarmStatusInterval, error)
	GetPointAlarmThresholdsWithContext(ctx context.Context, nodeID string) ([]pasapi.AlarmStatusInterval, error)

	SetPointAlarmStatus(input pasapi.SetPointAlarmStatusInput) error
	SetPointAlarmStatusWithContext(ctx context.Context, input pasapi.SetPointAlarmStatusInput) error

	GetPointAlarmStatus(input pasapi.GetPointAlarmStatusInput) (pasapi.AlarmStatus, error)
	GetPointAlarmStatusWithContext(ctx context.Context, input pasapi.GetPointAlarmStatusInput) (pasapi.AlarmStatus, error)

	GetPointAlarmStatusStream(dc chan<- pasapi.GetPointAlarmStatusStreamOutput) error
	GetPointAlarmStatusStreamWithContext(ctx context.Context, dc chan<- pasapi.GetPointAlarmStatusStreamOutput) error
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
