package reports

import (
	"context"
	"time"

	proto_reports "github.com/SKF/proto/reports"
)

func (c *client) GetAssetHealth(input proto_reports.GetAssetHealthInput) (output *proto_reports.GetAssetHealthOutput, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	return c.GetAssetHealthWithContext(ctx, input)
}
func (c *client) GetAssetHealthWithContext(ctx context.Context, input proto_reports.GetAssetHealthInput) (output *proto_reports.GetAssetHealthOutput, err error) {
	return c.api.GetAssetHealth(ctx, &input)
}

func (c *client) GetFunctionalLocationHealth(input proto_reports.GetFunctionalLocationHealthInput) (output *proto_reports.GetFunctionalLocationHealthOutput, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	return c.GetFunctionalLocationHealthWithContext(ctx, input)
}
func (c *client) GetFunctionalLocationHealthWithContext(ctx context.Context, input proto_reports.GetFunctionalLocationHealthInput) (output *proto_reports.GetFunctionalLocationHealthOutput, err error) {
	return c.api.GetFunctionalLocationHealth(ctx, &input)
}

func (c *client) GetComplianceLog(input proto_reports.GetComplianceLogInput) (output *proto_reports.GetComplianceLogOutput, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	return c.GetComplianceLogWithContext(ctx, input)
}
func (c *client) GetComplianceLogWithContext(ctx context.Context, input proto_reports.GetComplianceLogInput) (output *proto_reports.GetComplianceLogOutput, err error) {
	return c.api.GetComplianceLog(ctx, &input)
}

func (c *client) GetReports(input proto_reports.GetReportsInput) (output *proto_reports.GetReportsOutput, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	return c.GetReportsWithContext(ctx, input)
}
func (c *client) GetReportsWithContext(ctx context.Context, input proto_reports.GetReportsInput) (output *proto_reports.GetReportsOutput, err error) {
	return c.api.GetReports(ctx, &input)
}

func (c *client) GetComplianceSummary(input proto_reports.GetComplianceSummaryInput) (output *proto_reports.GetComplianceSummaryOutput, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	return c.GetComplianceSummaryWithContext(ctx, input)
}
func (c *client) GetComplianceSummaryWithContext(ctx context.Context, input proto_reports.GetComplianceSummaryInput) (output *proto_reports.GetComplianceSummaryOutput, err error) {
	return c.api.GetComplianceSummary(ctx, &input)
}
