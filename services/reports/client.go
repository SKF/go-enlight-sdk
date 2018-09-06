package reports

import (
	"context"
	"time"

	reports_grpcapi "github.com/SKF/proto/reports"
	"google.golang.org/grpc"
)

type ReportsClient interface {
	Dial(host, port string, opts ...grpc.DialOption) error
	Close()

	DeepPing() (output *reports_grpcapi.DeepPingOutput, err error)
	DeepPingWithContext(ctx context.Context) (output *reports_grpcapi.DeepPingOutput, err error)

	GetFunctionalLocationHealth(input reports_grpcapi.GetFunctionalLocationHealthInput) (output *reports_grpcapi.GetFunctionalLocationHealthOutput, err error)
	GetFunctionalLocationHealthWithContext(ctx context.Context, input reports_grpcapi.GetFunctionalLocationHealthInput) (output *reports_grpcapi.GetFunctionalLocationHealthOutput, err error)

	GetAssetHealth(input reports_grpcapi.GetAssetHealthInput) (output *reports_grpcapi.GetAssetHealthOutput, err error)
	GetAssetHealthWithContext(ctx context.Context, input reports_grpcapi.GetAssetHealthInput) (output *reports_grpcapi.GetAssetHealthOutput, err error)

	GetComplianceLog(input reports_grpcapi.GetComplianceLogInput) (output *reports_grpcapi.GetComplianceLogOutput, err error)
	GetComplianceLogWithContext(ctx context.Context, input reports_grpcapi.GetComplianceLogInput) (output *reports_grpcapi.GetComplianceLogOutput, err error)

	GetReports(input reports_grpcapi.GetReportsInput) (output *reports_grpcapi.GetReportsOutput, err error)
	GetReportsWithContext(ctx context.Context, input reports_grpcapi.GetReportsInput) (output *reports_grpcapi.GetReportsOutput, err error)

	GetComplianceSummary(input reports_grpcapi.GetComplianceSummaryInput) (output *reports_grpcapi.GetComplianceSummaryOutput, err error)
	GetComplianceSummaryWithContext(ctx context.Context, input reports_grpcapi.GetComplianceSummaryInput) (output *reports_grpcapi.GetComplianceSummaryOutput, err error)
}

type client struct {
	api  reports_grpcapi.ReportsClient
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
	c.api = reports_grpcapi.NewReportsClient(conn)
	return
}

func (c *client) DeepPing() (output *reports_grpcapi.DeepPingOutput, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	return c.DeepPingWithContext(ctx)
}
func (c *client) DeepPingWithContext(ctx context.Context) (output *reports_grpcapi.DeepPingOutput, err error) {
	return c.api.DeepPing(ctx, &reports_grpcapi.PrimitiveVoid{})
}
