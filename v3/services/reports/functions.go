package reports

import (
	"context"
	"time"

	reports_grpcapi "github.com/SKF/proto/v2/reports"
)

func (c *client) GetAssetHealth(input *reports_grpcapi.GetAssetHealthInput) (output *reports_grpcapi.GetAssetHealthOutput, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	return c.GetAssetHealthWithContext(ctx, input)
}
func (c *client) GetAssetHealthWithContext(ctx context.Context, input *reports_grpcapi.GetAssetHealthInput) (output *reports_grpcapi.GetAssetHealthOutput, err error) {
	return c.api.GetAssetHealth(ctx, input)
}

func (c *client) GetFunctionalLocationHealth(input *reports_grpcapi.GetFunctionalLocationHealthInput) (output *reports_grpcapi.GetFunctionalLocationHealthOutput, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	return c.GetFunctionalLocationHealthWithContext(ctx, input)
}
func (c *client) GetFunctionalLocationHealthWithContext(ctx context.Context, input *reports_grpcapi.GetFunctionalLocationHealthInput) (output *reports_grpcapi.GetFunctionalLocationHealthOutput, err error) {
	return c.api.GetFunctionalLocationHealth(ctx, input)
}

func (c *client) GetComplianceLog(input *reports_grpcapi.GetComplianceLogInput) (output *reports_grpcapi.GetComplianceLogOutput, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	return c.GetComplianceLogWithContext(ctx, input)
}
func (c *client) GetComplianceLogWithContext(ctx context.Context, input *reports_grpcapi.GetComplianceLogInput) (output *reports_grpcapi.GetComplianceLogOutput, err error) {
	return c.api.GetComplianceLog(ctx, input)
}

func (c *client) GetReports(input *reports_grpcapi.GetReportsInput) (output *reports_grpcapi.GetReportsOutput, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	return c.GetReportsWithContext(ctx, input)
}
func (c *client) GetReportsWithContext(ctx context.Context, input *reports_grpcapi.GetReportsInput) (output *reports_grpcapi.GetReportsOutput, err error) {
	return c.api.GetReports(ctx, input)
}

func (c *client) GetComplianceSummary(input *reports_grpcapi.GetComplianceSummaryInput) (output *reports_grpcapi.GetComplianceSummaryOutput, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	return c.GetComplianceSummaryWithContext(ctx, input)
}
func (c *client) GetComplianceSummaryWithContext(ctx context.Context, input *reports_grpcapi.GetComplianceSummaryInput) (output *reports_grpcapi.GetComplianceSummaryOutput, err error) {
	return c.api.GetComplianceSummary(ctx, input)
}
