package reports

import (
	"context"
	"time"

	"github.com/SKF/go-enlight-sdk/services/reports/reportsgrpcapi"
)

func (c *client) DeepPing() (_ reportsgrpcapi.PrimitiveString, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	_, err = c.api.DeepPing(ctx, &reportsgrpcapi.PrimitiveVoid{})
	return
}

func (c *client) GetAssetHealth(input reportsgrpcapi.GetAssetHealthInput) (output reportsgrpcapi.GetAssetHealthOutput, err error) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	output, err = c.api.GetAssetHealth(ctx, &input)
	return
}

func (c *client) GetFunctionalLocationHealth(input reportsgrpcapi.GetFunctionalLocationHealthInput) (output reportsgrpcapi.GetFunctionalLocationHealthOutput, err error) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	output, err = c.api.GetFunctionalLocationHealth(ctx, &input)
	return
}

func (c *client) GetComplianceLog(input reportsgrpcapi.GetComplianceLogInput) (output reportsgrpcapi.GetComplianceLogOutput, err error) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	output, err = c.api.GetComplianceLog(ctx, &input)
	return
}
