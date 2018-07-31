package mock

import (
	"github.com/SKF/go-enlight-sdk/services/reports"
	api "github.com/SKF/go-enlight-sdk/services/reports/reportsgrpcapi"
	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc"
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
func (mock *client) DeepPing() (output *api.DeepPingOutput, err error) {
	args := mock.Called()
	return args.Get(0).(&api.DeepPingOutput), args.Error(1)
}

func (mock *client) GetFunctionalLocationHealth(input api.GetFunctionalLocationHealthInput) (output *api.GetFunctionalLocationHealthOutput, err error) {
	args := mock.Called(input)
	return args.Get(0).(&api.GetFunctionalLocationHealthOutput), args.Error(1)
}

func (mock *client) GetAssetHealth(input api.GetAssetHealthInput) (output *api.GetAssetHealthOutput, err error) {
	args := mock.Called(input)
	return args.Get(0).(&api.GetAssetHealthOutput), args.Error(1)
}

func (mock *client) GetComplianceLog(input api.GetComplianceLogInput) (output *api.GetComplianceLogOutput, err error) {
	args := mock.Called(input)
	return args.Get(0).(&api.GetComplianceLogOutput), args.Error(1)
}
