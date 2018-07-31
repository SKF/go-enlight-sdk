package mock

import (
	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc"
	"skfdc.visualstudio.com/enlightcentre/go-enlight-sdk/services/reports"
	"skfdc.visualstudio.com/enlightcentre/go-enlight-sdk/services/reports/reportsgrpcapi"
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
func (mock *client) DeepPing() error {
	args := mock.Called()
	return args.Error(0)
}

func (mock *client) GetFunctionalLocationHealth(input reportsgrpcapi.GetFunctionalLocationHealthInput) (output reportsgrpcapi.GetFunctionalLocationHealthOutput, err error) {
	args := mock.Called(input)
	return args.Error(0)
}

func (mock *client) GetAssetHealth(input reportsgrpcapi.GetAssetHealthInput) (output reportsgrpcapi.GetAssetHealthOutput, err error) {
	args := mock.Called(input)
	return args.Error(0)
}

func (mock *client) GetComplianceLog(input reportsgrpcapi.GetComplianceLogInput) (output reportsgrpcapi.GetComplianceLogOutput, err error) {
	args := mock.Called(input)
	return args.Error(0)
}
