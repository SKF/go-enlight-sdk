package credentialsmanager

import (
	"context"
	"errors"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	api_mock "github.com/SKF/go-enlight-sdk/services/authorize/credentialsmanager/mock"
)

func Test_V2_GetDataStore_Success(t *testing.T) {
	t.Parallel()

	secret := `{"ca":"YWFhYQ==","key":"ZWVl","crt":"ZmZmZg=="}`
	expected := DataStore{
		CA:  []byte("aaaa"),
		Key: []byte("eee"),
		Crt: []byte("ffff"),
	}

	api := api_mock.CreateAPIV2()
	manager := New().UsingSDKV2(api)

	api.On("GetSecretValue", mock.Anything, mock.Anything, mock.Anything).Return(&secretsmanager.GetSecretValueOutput{
		SecretString: &secret,
	}, nil).Once()

	out, err := manager.GetDataStore(context.TODO(), "random-secret-123")
	assert.Equal(t, expected, *out)
	assert.NoError(t, err)

	api.AssertExpectations(t)
}

func Test_V2_GetDataStore_JsonError(t *testing.T) {
	t.Parallel()

	secret := `{"ca":"YWFhYQ==","key":"ZWVl","crt":"ZmZmZg=="`

	api := api_mock.CreateAPIV2()
	manager := New().UsingSDKV2(api)

	api.On("GetSecretValue", mock.Anything, mock.Anything, mock.Anything).Return(&secretsmanager.GetSecretValueOutput{
		SecretString: &secret,
	}, nil).Once()

	out, err := manager.GetDataStore(context.TODO(), "random-secret-123")
	assert.Nil(t, out)
	assert.Error(t, err)

	api.AssertExpectations(t)
}

func Test_V2_GetDataStore_GetSecretError(t *testing.T) {
	t.Parallel()

	var secret *secretsmanager.GetSecretValueOutput
	expectedErr := errors.New("failed to fetch secret")

	api := api_mock.CreateAPIV2()
	manager := New().UsingSDKV2(api)

	api.On("GetSecretValue", mock.Anything, mock.Anything, mock.Anything).Return(secret, expectedErr).Once()

	out, err := manager.GetDataStore(context.TODO(), "random-secret-123")
	assert.Nil(t, out)
	assert.Error(t, err)

	api.AssertExpectations(t)
}
