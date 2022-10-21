package credentialsmanager

import (
	"context"

	awsV2 "github.com/aws/aws-sdk-go-v2/aws"
	smV2 "github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	sessV1 "github.com/aws/aws-sdk-go/aws/session"
	smV1 "github.com/aws/aws-sdk-go/service/secretsmanager"
)

type DataStore struct {
	CA  []byte `json:"ca"`
	Key []byte `json:"key"`
	Crt []byte `json:"crt"`
}

type CredentialsManager struct {
	fetcher credentialsFetcher
}

type SMAPIV1 interface {
	GetSecretValue(input *smV1.GetSecretValueInput) (*smV1.GetSecretValueOutput, error)
}

type credentialsFetcher interface {
	GetDataStore(ctx context.Context, secretsName string) (*DataStore, error)
}

func New() *CredentialsManager {
	return &CredentialsManager{}
}

func (cm *CredentialsManager) UsingSDKV1Session(sess *sessV1.Session) *CredentialsManager {
	sm := smV1.New(sess)
	return cm.UsingSDKV1(sm)
}

func (cm *CredentialsManager) UsingSDKV1(sm SMAPIV1) *CredentialsManager {
	cm.fetcher = &credentialsManagerV1{sm: sm}
	return cm
}

func (cm *CredentialsManager) UsingSDKV2Config(cfg awsV2.Config) *CredentialsManager {
	sm := smV2.NewFromConfig(cfg)
	return cm.UsingSDKV2(sm)
}

func (cm *CredentialsManager) UsingSDKV2(sm SMAPIV2) *CredentialsManager {
	cm.fetcher = &credentialsManagerV2{sm: sm}
	return cm
}

func (cm *CredentialsManager) GetDataStore(ctx context.Context, secretsName string) (*DataStore, error) {
	return cm.fetcher.GetDataStore(ctx, secretsName)
}
