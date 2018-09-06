package mock

import (
	"context"

	proto_reports "github.com/SKF/proto/reports"
	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc"

	"github.com/SKF/go-enlight-sdk/services/reports"
)

type client struct {
	mock.Mock
}

func Create() *client {
	return new(client)
}

var _ reports.ReportsClient = &client{}

func (mock *client) Dial(host, port string, opts ...grpc.DialOption) error {
	args := mock.Called(host, port, opts)
	return args.Error(0)
}
func (mock *client) Close() {
	mock.Called()
	return
}

func (mock *client) DeepPing() (output *proto_reports.DeepPingOutput, err error) {
	args := mock.Called()
	return args.Get(0).(*proto_reports.DeepPingOutput), args.Error(1)
}
func (mock *client) DeepPingWithContext(ctx context.Context) (output *proto_reports.DeepPingOutput, err error) {
	args := mock.Called(ctx)
	return args.Get(0).(*proto_reports.DeepPingOutput), args.Error(1)
}

func (mock *client) GetFunctionalLocationHealth(input proto_reports.GetFunctionalLocationHealthInput) (output *proto_reports.GetFunctionalLocationHealthOutput, err error) {
	args := mock.Called(input)
	return args.Get(0).(*proto_reports.GetFunctionalLocationHealthOutput), args.Error(1)
}
func (mock *client) GetFunctionalLocationHealthWithContext(ctx context.Context, input proto_reports.GetFunctionalLocationHealthInput) (output *proto_reports.GetFunctionalLocationHealthOutput, err error) {
	args := mock.Called(ctx, input)
	return args.Get(0).(*proto_reports.GetFunctionalLocationHealthOutput), args.Error(1)
}

func (mock *client) GetAssetHealth(input proto_reports.GetAssetHealthInput) (output *proto_reports.GetAssetHealthOutput, err error) {
	args := mock.Called(input)
	return args.Get(0).(*proto_reports.GetAssetHealthOutput), args.Error(1)
}
func (mock *client) GetAssetHealthWithContext(ctx context.Context, input proto_reports.GetAssetHealthInput) (output *proto_reports.GetAssetHealthOutput, err error) {
	args := mock.Called(ctx, input)
	return args.Get(0).(*proto_reports.GetAssetHealthOutput), args.Error(1)
}

func (mock *client) GetComplianceLog(input proto_reports.GetComplianceLogInput) (output *proto_reports.GetComplianceLogOutput, err error) {
	args := mock.Called(input)
	return args.Get(0).(*proto_reports.GetComplianceLogOutput), args.Error(1)
}
func (mock *client) GetComplianceLogWithContext(ctx context.Context, input proto_reports.GetComplianceLogInput) (output *proto_reports.GetComplianceLogOutput, err error) {
	args := mock.Called(ctx, input)
	return args.Get(0).(*proto_reports.GetComplianceLogOutput), args.Error(1)
}

func (mock *client) GetReports(input proto_reports.GetReportsInput) (output *proto_reports.GetReportsOutput, err error) {
	args := mock.Called(input)
	return args.Get(0).(*proto_reports.GetReportsOutput), args.Error(1)
}
func (mock *client) GetReportsWithContext(ctx context.Context, input proto_reports.GetReportsInput) (output *proto_reports.GetReportsOutput, err error) {
	args := mock.Called(ctx, input)
	return args.Get(0).(*proto_reports.GetReportsOutput), args.Error(1)
}

func (mock *client) GetComplianceSummary(input proto_reports.GetComplianceSummaryInput) (output *proto_reports.GetComplianceSummaryOutput, err error) {
	args := mock.Called(input)
	return args.Get(0).(*proto_reports.GetComplianceSummaryOutput), args.Error(1)
}
func (mock *client) GetComplianceSummaryWithContext(ctx context.Context, input proto_reports.GetComplianceSummaryInput) (output *proto_reports.GetComplianceSummaryOutput, err error) {
	args := mock.Called(ctx, input)
	return args.Get(0).(*proto_reports.GetComplianceSummaryOutput), args.Error(1)
}
