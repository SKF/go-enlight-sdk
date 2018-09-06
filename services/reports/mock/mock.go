package mock

import (
	"context"

	reports_grpcapi "github.com/SKF/proto/reports"
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

func (mock *client) DeepPing() (output *reports_grpcapi.DeepPingOutput, err error) {
	args := mock.Called()
	return args.Get(0).(*reports_grpcapi.DeepPingOutput), args.Error(1)
}
func (mock *client) DeepPingWithContext(ctx context.Context) (output *reports_grpcapi.DeepPingOutput, err error) {
	args := mock.Called(ctx)
	return args.Get(0).(*reports_grpcapi.DeepPingOutput), args.Error(1)
}

func (mock *client) GetFunctionalLocationHealth(input reports_grpcapi.GetFunctionalLocationHealthInput) (output *reports_grpcapi.GetFunctionalLocationHealthOutput, err error) {
	args := mock.Called(input)
	return args.Get(0).(*reports_grpcapi.GetFunctionalLocationHealthOutput), args.Error(1)
}
func (mock *client) GetFunctionalLocationHealthWithContext(ctx context.Context, input reports_grpcapi.GetFunctionalLocationHealthInput) (output *reports_grpcapi.GetFunctionalLocationHealthOutput, err error) {
	args := mock.Called(ctx, input)
	return args.Get(0).(*reports_grpcapi.GetFunctionalLocationHealthOutput), args.Error(1)
}

func (mock *client) GetAssetHealth(input reports_grpcapi.GetAssetHealthInput) (output *reports_grpcapi.GetAssetHealthOutput, err error) {
	args := mock.Called(input)
	return args.Get(0).(*reports_grpcapi.GetAssetHealthOutput), args.Error(1)
}
func (mock *client) GetAssetHealthWithContext(ctx context.Context, input reports_grpcapi.GetAssetHealthInput) (output *reports_grpcapi.GetAssetHealthOutput, err error) {
	args := mock.Called(ctx, input)
	return args.Get(0).(*reports_grpcapi.GetAssetHealthOutput), args.Error(1)
}

func (mock *client) GetComplianceLog(input reports_grpcapi.GetComplianceLogInput) (output *reports_grpcapi.GetComplianceLogOutput, err error) {
	args := mock.Called(input)
	return args.Get(0).(*reports_grpcapi.GetComplianceLogOutput), args.Error(1)
}
func (mock *client) GetComplianceLogWithContext(ctx context.Context, input reports_grpcapi.GetComplianceLogInput) (output *reports_grpcapi.GetComplianceLogOutput, err error) {
	args := mock.Called(ctx, input)
	return args.Get(0).(*reports_grpcapi.GetComplianceLogOutput), args.Error(1)
}

func (mock *client) GetReports(input reports_grpcapi.GetReportsInput) (output *reports_grpcapi.GetReportsOutput, err error) {
	args := mock.Called(input)
	return args.Get(0).(*reports_grpcapi.GetReportsOutput), args.Error(1)
}
func (mock *client) GetReportsWithContext(ctx context.Context, input reports_grpcapi.GetReportsInput) (output *reports_grpcapi.GetReportsOutput, err error) {
	args := mock.Called(ctx, input)
	return args.Get(0).(*reports_grpcapi.GetReportsOutput), args.Error(1)
}

func (mock *client) GetComplianceSummary(input reports_grpcapi.GetComplianceSummaryInput) (output *reports_grpcapi.GetComplianceSummaryOutput, err error) {
	args := mock.Called(input)
	return args.Get(0).(*reports_grpcapi.GetComplianceSummaryOutput), args.Error(1)
}
func (mock *client) GetComplianceSummaryWithContext(ctx context.Context, input reports_grpcapi.GetComplianceSummaryInput) (output *reports_grpcapi.GetComplianceSummaryOutput, err error) {
	args := mock.Called(ctx, input)
	return args.Get(0).(*reports_grpcapi.GetComplianceSummaryOutput), args.Error(1)
}
