package reports

import (
	"context"
	"time"

	"github.com/SKF/go-enlight-sdk/services/reports/reportsgrpcapi"
	"google.golang.org/grpc"
)

type ReportsClient interface {
	Dial(host, port string, opts ...grpc.DialOption) error
	Close()

	DeepPing() (output *reportsgrpcapi.DeepPingOutput, err error)
	GetFunctionalLocationHealth(input reportsgrpcapi.GetFunctionalLocationHealthInput) (output *reportsgrpcapi.GetFunctionalLocationHealthOutput, err error)
	GetAssetHealth(input reportsgrpcapi.GetAssetHealthInput) (output *reportsgrpcapi.GetAssetHealthOutput, err error)
	GetComplianceLog(input reportsgrpcapi.GetComplianceLogInput) (output *reportsgrpcapi.GetComplianceLogOutput, err error)
}

type client struct {
	api  reportsgrpcapi.ReportsClient
	conn *grpc.ClientConn
}

func (c *client) Close() {
	c.conn.Close()
}

func CreateClient() ReportsClient {
	return &client{}
}

func (c *client) Dial(host, port string, opts ...grpc.DialOption) (err error) {
	conn, err := grpc.Dial(host+":"+port, opts...)
	if err != nil {
		return
	}

	c.conn = conn
	c.api = reportsgrpcapi.NewReportsClient(conn)
	return
}

func (c *client) DeepPing() (output *reportsgrpcapi.DeepPingOutput, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	return c.api.DeepPing(ctx, &reportsgrpcapi.PrimitiveVoid{})
}
