package reports

import (
	"context"
	"time"

	"github.com/SKF/go-enlight-sdk/services/reports/reportsgrpcapi"
)

func (c *client) GetAssetHealth(input reportsgrpcapi.GetAssetHealthInput) (output *reportsgrpcapi.GetAssetHealthOutput, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	return c.GetAssetHealthWithContext(ctx, input)
}
func (c *client) GetAssetHealthWithContext(ctx context.Context, input reportsgrpcapi.GetAssetHealthInput) (output *reportsgrpcapi.GetAssetHealthOutput, err error) {
	return c.api.GetAssetHealth(ctx, &input)
}

func (c *client) GetFunctionalLocationHealth(input reportsgrpcapi.GetFunctionalLocationHealthInput) (output *reportsgrpcapi.GetFunctionalLocationHealthOutput, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	return c.GetFunctionalLocationHealthWithContext(ctx, input)
}
func (c *client) GetFunctionalLocationHealthWithContext(ctx context.Context, input reportsgrpcapi.GetFunctionalLocationHealthInput) (output *reportsgrpcapi.GetFunctionalLocationHealthOutput, err error) {
	return c.api.GetFunctionalLocationHealth(ctx, &input)
}

func (c *client) GetComplianceLog(input reportsgrpcapi.GetComplianceLogInput) (output *reportsgrpcapi.GetComplianceLogOutput, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	return c.GetComplianceLogWithContext(ctx, input)
}
func (c *client) GetComplianceLogWithContext(ctx context.Context, input reportsgrpcapi.GetComplianceLogInput) (output *reportsgrpcapi.GetComplianceLogOutput, err error) {
	return c.api.GetComplianceLog(ctx, &input)
}
