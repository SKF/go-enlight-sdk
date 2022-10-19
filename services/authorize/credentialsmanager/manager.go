package credentialsmanager

import "context"

type DataStore struct {
	CA  []byte `json:"ca"`
	Key []byte `json:"key"`
	Crt []byte `json:"crt"`
}

type CredentialsManager interface {
	GetDataStore(ctx context.Context, secretsName string) (*DataStore, error)
}

type CredentialsManagerBuilder struct{}

func New() *CredentialsManagerBuilder {
	return &CredentialsManagerBuilder{}
}
