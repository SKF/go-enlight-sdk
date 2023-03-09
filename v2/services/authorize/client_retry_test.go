package authorize_test

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	_ "embed"
	"fmt"
	"net"
	"testing"

	"github.com/SKF/go-enlight-sdk/v2/services/authorize"
	"github.com/SKF/go-enlight-sdk/v2/services/authorize/credentialsmanager"
	authorizeproto "github.com/SKF/proto/v2/authorize"
	"github.com/SKF/proto/v2/common"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/status"
)

var (
	//go:embed certs/ca-cert.pem
	caCert []byte
	//go:embed certs/client-key.pem
	clientKey []byte
	//go:embed certs/client-cert.pem
	clientCert []byte
	//go:embed certs/server-key.pem
	serverKey []byte
	//go:embed certs/server-cert.pem
	serverCert []byte
)

type mockCredentialsFetcher struct{}

func (mock *mockCredentialsFetcher) GetDataStore(ctx context.Context, secretsName string) (*credentialsmanager.DataStore, error) {
	return &credentialsmanager.DataStore{
		CA:  caCert,
		Key: clientKey,
		Crt: clientCert,
	}, nil
}

func loadTLSCredentials() (credentials.TransportCredentials, error) {
	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(caCert) {
		return nil, fmt.Errorf("failed to add CA certificate")
	}

	serverCert, err := tls.X509KeyPair(serverCert, serverKey)
	if err != nil {
		return nil, err
	}

	config := &tls.Config{
		Certificates: []tls.Certificate{serverCert},
		ClientAuth:   tls.RequireAndVerifyClientCert,
		ClientCAs:    certPool,
		MinVersion:   tls.VersionTLS13,
	}

	return credentials.NewTLS(config), nil
}

type dummyAuthorizeServer struct {
	authorizeproto.UnimplementedAuthorizeServer
	timesFailed int
}

func (*dummyAuthorizeServer) LogClientState(context.Context, *authorizeproto.LogClientStateInput) (*common.Void, error) {
	return &common.Void{}, nil
}

func (srv *dummyAuthorizeServer) AddResource(context.Context, *authorizeproto.AddResourceInput) (*common.Void, error) {
	if srv.timesFailed++; srv.timesFailed <= 4 {
		return nil, status.Errorf(codes.Canceled, "too slow")
	}

	return &common.Void{}, nil
}

func newServer() *dummyAuthorizeServer {
	s := &dummyAuthorizeServer{}
	return s
}

func TestRetryPolicy(t *testing.T) {
	lis, err := net.Listen("tcp", "localhost:10000")
	require.NoError(t, err)

	defer lis.Close()

	tlsCredentials, err := loadTLSCredentials()
	require.NoError(t, err)

	var serverOpts []grpc.ServerOption
	serverOpts = append(serverOpts, grpc.Creds(tlsCredentials))

	grpcServer := grpc.NewServer(serverOpts...)
	authorizeproto.RegisterAuthorizeServer(grpcServer, newServer())

	//nolint:errcheck
	go grpcServer.Serve(lis)

	c := authorize.CreateClient()

	err = c.DialUsingCredentialsManager(context.Background(), &mockCredentialsFetcher{}, "localhost", "10000", "")
	require.NoError(t, err)

	err = c.AddResourceWithContext(context.Background(), common.Origin{
		Id:       "",
		Type:     "",
		Provider: "",
	})
	require.NoError(t, err)
}
