package mock

import (
	"github.com/aws/aws-sdk-go/service/secretsmanager"
	"github.com/stretchr/testify/mock"
)

type APIV1Mock struct {
	mock.Mock
}

func CreateAPIV1() *APIV1Mock {
	return &APIV1Mock{
		mock.Mock{},
	}
}

func (a *APIV1Mock) GetSecretValue(input *secretsmanager.GetSecretValueInput) (*secretsmanager.GetSecretValueOutput, error) {
	args := a.Called(input)
	return args.Get(0).(*secretsmanager.GetSecretValueOutput), args.Error(1)
}
