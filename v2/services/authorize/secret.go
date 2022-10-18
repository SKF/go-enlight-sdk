package authorize

import (
	"context"
	"fmt"

	googleGrpc "google.golang.org/grpc"

	"github.com/SKF/go-enlight-sdk/v2/grpc"
	"github.com/SKF/go-enlight-sdk/v2/services/authorize/credentialsmanager"
)

func getCredentialOption(ctx context.Context, host, secretKeyName string, cm credentialsmanager.CredentialsManager) (googleGrpc.DialOption, error) {
	clientCert, err := cm.GetDataStore(ctx, secretKeyName)
	if err != nil {
		panic(err)
	}

	return grpc.WithTransportCredentialsPEM(
		host,
		clientCert.Crt, clientCert.Key, clientCert.CA,
	)
}

func GetSecretKeyName(service, stage string) string {
	return fmt.Sprintf("authorize/%s/grpc/client/%s", stage, service)
}

func GetSecretKeyArn(accountId, region, service, stage string) string {
	return fmt.Sprintf("arn:aws:secretsmanager:%s:%s:secret:%s", region, accountId, GetSecretKeyName(service, stage))
}
