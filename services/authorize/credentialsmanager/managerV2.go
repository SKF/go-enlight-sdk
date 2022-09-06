package credentialsmanager

import (
	"context"
	"encoding/json"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/pkg/errors"

	"github.com/SKF/go-utility/v2/log"
)

type SMAPIV2 interface {
	GetSecretValue(ctx context.Context, params *secretsmanager.GetSecretValueInput, optFns ...func(*secretsmanager.Options)) (*secretsmanager.GetSecretValueOutput, error)
}

type CredentialsManagerV2 struct {
	sm SMAPIV2
}

func CreateCredentialManagerV2(sm SMAPIV2) CredentialsManager {
	return &CredentialsManagerV2{
		sm: sm,
	}
}

func (cm *CredentialsManagerV2) GetDataStore(ctx context.Context, secretsName string) (*DataStore, error) {
	input := &secretsmanager.GetSecretValueInput{
		SecretId:     aws.String(secretsName),
		VersionStage: aws.String("AWSCURRENT"),
	}

	result, err := cm.sm.GetSecretValue(ctx, input)
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
