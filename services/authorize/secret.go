package authorize

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/SKF/go-enlight-sdk/grpc"
	"github.com/SKF/go-utility/log"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
	"github.com/pkg/errors"
	googleGrpc "google.golang.org/grpc"
)

type dataStore struct {
	CA  []byte `json:"ca"`
	Key []byte `json:"key"`
	Crt []byte `json:"crt"`
}

func getSecret(ctx context.Context, sess *session.Session, secretsName string, out interface{}) (err error) {
	// credentials - default
	svc := secretsmanager.New(sess)
	input := &secretsmanager.GetSecretValueInput{
		SecretId:     aws.String(secretsName),
		VersionStage: aws.String("AWSCURRENT"),
	}

	result, err := svc.GetSecretValue(input)
	if err != nil {
		log.WithTracing(ctx).WithError(err).
			WithField("secretsName", secretsName).
			Error("failed to get secrets")
		err = errors.Wrapf(err, "failed to get secret value from '%s'", secretsName)
		return
	}

	if err = json.Unmarshal([]byte(*result.SecretString), out); err != nil {
		log.WithTracing(ctx).WithError(err).
			WithField("secretsName", secretsName).
			Error("failed to unmarshal secret")
		err = errors.Wrapf(err, "failed to unmarshal secret from '%s'", secretsName)
	}

	return err
}

func getCredentialOption(ctx context.Context, sess *session.Session, host, secretKeyName string) (googleGrpc.DialOption, error) {
	var clientCert dataStore
	if err := getSecret(ctx, sess, secretKeyName, &clientCert); err != nil {
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
