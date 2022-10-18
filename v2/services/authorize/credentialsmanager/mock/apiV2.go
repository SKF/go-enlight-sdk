package mock

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/stretchr/testify/mock"
)

type APIV2Mock struct {
	mock.Mock
}

func CreateAPIV2() *APIV2Mock {
	return &APIV2Mock{
		mock.Mock{},
	}
}

func (a *APIV2Mock) GetSecretValue(ctx context.Context, params *secretsmanager.GetSecretValueInput, optFns ...func(*secretsmanager.Options)) (*secretsmanager.GetSecretValueOutput, error) {
	args := a.Called(ctx, params, optFns)
	return args.Get(0).(*secretsmanager.GetSecretValueOutput), args.Error(1)
}
