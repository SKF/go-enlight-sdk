package credentialsmanager

import (
	"context"
	"encoding/json"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
	"github.com/pkg/errors"

	"github.com/SKF/go-utility/v2/log"
)

type SMAPIV1 interface {
	GetSecretValue(input *secretsmanager.GetSecretValueInput) (*secretsmanager.GetSecretValueOutput, error)
}

type CredentialsMangerV1 struct {
	sm SMAPIV1
}

func CreateCredentialsManagerV1(sm SMAPIV1) CredentialsManager {
	return &CredentialsMangerV1{
		sm: sm,
	}
}

func (cm *CredentialsMangerV1) GetDataStore(ctx context.Context, secretsName string) (*DataStore, error) {
	input := &secretsmanager.GetSecretValueInput{
		SecretId:     aws.String(secretsName),
		VersionStage: aws.String("AWSCURRENT"),
	}

	result, err := cm.sm.GetSecretValue(input)
	if err != nil {
		log.WithTracing(ctx).WithError(err).
			WithField("secretsName", secretsName).
			Error("failed to get secrets")
		err = errors.Wrapf(err, "failed to get secret value from '%s'", secretsName)
		return nil, err
	}

	var out DataStore

	if err = json.Unmarshal([]byte(*result.SecretString), &out); err != nil {
		log.WithTracing(ctx).WithError(err).
			WithField("secretsName", secretsName).
			Error("failed to unmarshal secret")
		err = errors.Wrapf(err, "failed to unmarshal secret from '%s'", secretsName)

		return nil, err
	}

	return &out, err
}
