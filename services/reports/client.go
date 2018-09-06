package reports

import (
	"context"
	"time"

	proto_common "github.com/SKF/proto/common"
	proto_reports "github.com/SKF/proto/reports"
	"google.golang.org/grpc"
)

type ReportsClient interface {
	Dial(host, port string, opts ...grpc.DialOption) error
	Close()

	DeepPing() (output *proto_reports.DeepPingOutput, err error)
	DeepPingWithContext(ctx context.Context) (output *proto_reports.DeepPingOutput, err error)

	GetFunctionalLocationHealth(input proto_reports.GetFunctionalLocationHealthInput) (output *proto_reports.GetFunctionalLocationHealthOutput, err error)
	GetFunctionalLocationHealthWithContext(ctx context.Context, input proto_reports.GetFunctionalLocationHealthInput) (output *proto_reports.GetFunctionalLocationHealthOutput, err error)

	GetAssetHealth(input proto_reports.GetAssetHealthInput) (output *proto_reports.GetAssetHealthOutput, err error)
	GetAssetHealthWithContext(ctx context.Context, input proto_reports.GetAssetHealthInput) (output *proto_reports.GetAssetHealthOutput, err error)

	GetComplianceLog(input proto_reports.GetComplianceLogInput) (output *proto_reports.GetComplianceLogOutput, err error)
	GetComplianceLogWithContext(ctx context.Context, input proto_reports.GetComplianceLogInput) (output *proto_reports.GetComplianceLogOutput, err error)

	GetReports(input proto_reports.GetReportsInput) (output *proto_reports.GetReportsOutput, err error)
	GetReportsWithContext(ctx context.Context, input proto_reports.GetReportsInput) (output *proto_reports.GetReportsOutput, err error)

	GetComplianceSummary(input proto_reports.GetComplianceSummaryInput) (output *proto_reports.GetComplianceSummaryOutput, err error)
	GetComplianceSummaryWithContext(ctx context.Context, input proto_reports.GetComplianceSummaryInput) (output *proto_reports.GetComplianceSummaryOutput, err error)
}

type client struct {
	api  proto_reports.ReportsClient
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
	c.api = proto_reports.NewReportsClient(conn)
	return
}

func (c *client) DeepPing() (output *proto_reports.DeepPingOutput, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	return c.DeepPingWithContext(ctx)
}
func (c *client) DeepPingWithContext(ctx context.Context) (output *proto_reports.DeepPingOutput, err error) {
	return c.api.DeepPing(ctx, &proto_common.Void{})
}
