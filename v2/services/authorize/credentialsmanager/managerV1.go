package credentialsmanager

import (
	"context"
	"encoding/json"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
	"github.com/pkg/errors"

	"github.com/SKF/go-utility/v2/log"
)

type SMAPIV1 interface {
	GetSecretValue(input *secretsmanager.GetSecretValueInput) (*secretsmanager.GetSecretValueOutput, error)
}

type CredentialsManagerV1 struct {
	sm SMAPIV1
}

func (cm *CredentialsManager) UsingSDKV1Session(sess *session.Session) *CredentialsManager {
	sm := secretsmanager.New(sess)
	return cm.UsingSDKV1(sm)
}

func (cm *CredentialsManager) UsingSDKV1(sm SMAPIV1) *CredentialsManager {
	cm.fetcher = &CredentialsManagerV1{sm: sm}
	return cm
}

func (cm *CredentialsManagerV1) GetDataStore(ctx context.Context, secretsName string) (*DataStore, error) {
	input := &secretsmanager.GetSecretValueInput{
		SecretId:     aws.String(secretsName),
		VersionStage: aws.String("AWSCURRENT"),
	}

	logger := log.
		WithField("secretsName", secretsName).
		WithField("credentialsManager", "V1")

	result, err := cm.sm.GetSecretValue(input)
	if err != nil {
		logger.WithTracing(ctx).WithError(err).
			Error("failed to get secrets")
		err = errors.Wrapf(err, "failed to get secret value from '%s'", secretsName)
		return nil, err
	}

	var out DataStore

	if err = json.Unmarshal([]byte(*result.SecretString), &out); err != nil {
		logger.WithTracing(ctx).WithError(err).
			Error("failed to unmarshal secret")
		err = errors.Wrapf(err, "failed to unmarshal secret from '%s'", secretsName)

		return nil, err
	}

	return &out, err
}
