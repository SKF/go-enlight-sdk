package credentialsmanager

import "context"

type DataStore struct {
	CA  []byte `json:"ca"`
	Key []byte `json:"key"`
	Crt []byte `json:"crt"`
}

type CredentialsManager struct {
	fetcher credentialsFetcher
}

type credentialsFetcher interface {
	GetDataStore(ctx context.Context, secretsName string) (*DataStore, error)
}

func New() *CredentialsManager {
	return &CredentialsManager{}
}

func (cm *CredentialsManager) GetDataStore(ctx context.Context, secretsName string) (*DataStore, error) {
	return cm.fetcher.GetDataStore(ctx, secretsName)
}
