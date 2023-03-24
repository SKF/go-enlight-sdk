package authorize_test

import (
	"bytes"
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	_ "embed"
	"encoding/pem"
	"fmt"
	"math/big"
	"net"
	"testing"
	"time"

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
	//go:embed certs/rsa-key.pem
	rsaKey []byte
)

var ca = &x509.Certificate{
	SerialNumber:          big.NewInt(2019),
	Subject:               pkix.Name{},
	NotBefore:             time.Now(),
	NotAfter:              time.Now().AddDate(10, 0, 0),
	IsCA:                  true,
	BasicConstraintsValid: true,
}

func generateCA(privateKey *rsa.PrivateKey) ([]byte, error) {
	caBytes, err := x509.CreateCertificate(rand.Reader, ca, ca, &privateKey.PublicKey, privateKey)
	if err != nil {
		return nil, err
	}

	caPEM := new(bytes.Buffer)
	err = pem.Encode(caPEM, &pem.Block{
		Type:  "CERTIFICATE",
		Bytes: caBytes,
	})

	return caPEM.Bytes(), err
}

func generateDatastore(ca *x509.Certificate, privateKey *rsa.PrivateKey, caCertPEM []byte, validTime time.Duration) (credentialsmanager.DataStore, error) {
	ds := credentialsmanager.DataStore{}

	cert := &x509.Certificate{
		SerialNumber: big.NewInt(2019),
		Subject:      pkix.Name{},
		DNSNames:     []string{"localhost"},
		NotBefore:    time.Now(),
		NotAfter:     time.Now().Add(validTime),
		SubjectKeyId: []byte{1, 2, 3, 4, 6},
	}

	certBytes, err := x509.CreateCertificate(rand.Reader, cert, ca, &privateKey.PublicKey, privateKey)
	if err != nil {
		return ds, err
	}

	certPEM := new(bytes.Buffer)
	err = pem.Encode(certPEM, &pem.Block{
		Type:  "CERTIFICATE",
		Bytes: certBytes,
	})
	if err != nil {
		return ds, err
	}

	certPrivKeyPEM := new(bytes.Buffer)

	err = pem.Encode(certPrivKeyPEM, &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	})
	if err != nil {
		return ds, err
	}

	ds.Crt = certPEM.Bytes()
	ds.Key = certPrivKeyPEM.Bytes()
	ds.CA = caCertPEM

	return ds, nil
}

type mockCredentialsFetcher struct {
	ds        credentialsmanager.DataStore
	callCount int
}

func (mock *mockCredentialsFetcher) GetDataStore(ctx context.Context, secretsName string) (*credentialsmanager.DataStore, error) {
	mock.callCount += 1
	return &mock.ds, nil
}

func loadTLSCredentials(ds credentialsmanager.DataStore) (credentials.TransportCredentials, error) {
	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(ds.CA) {
		return nil, fmt.Errorf("failed to add CA certificate")
	}

	serverCert, err := tls.X509KeyPair(ds.Crt, ds.Key)
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

func (srv *dummyAuthorizeServer) GetResource(context.Context, *authorizeproto.GetResourceInput) (*authorizeproto.GetResourceOutput, error) {
	return &authorizeproto.GetResourceOutput{
		Resource: &common.Origin{
			Id:       "",
			Type:     "",
			Provider: "",
		},
	}, nil
}

func newServer() *dummyAuthorizeServer {
	s := &dummyAuthorizeServer{}
	return s
}

func parseRSAKey() (*rsa.PrivateKey, error) {
	pemBlock, _ := pem.Decode(rsaKey)
	k, err := x509.ParsePKCS8PrivateKey(pemBlock.Bytes)
	if err != nil {
		return nil, err
	}
	return k.(*rsa.PrivateKey), nil
}

type server struct {
	signal chan int
}

func launchServer(tlsCredentials credentials.TransportCredentials) (server, error) {
	server := server{
		signal: make(chan int),
	}

	var serverOpts []grpc.ServerOption
	serverOpts = append(serverOpts, grpc.Creds(tlsCredentials))

	grpcServer := grpc.NewServer(serverOpts...)
	authorizeproto.RegisterAuthorizeServer(grpcServer, newServer())

	lis, err := net.Listen("tcp", "localhost:10000")
	if err != nil {
		return server, err
	}

	go func() {
		//nolint:errcheck
		go grpcServer.Serve(lis)

		<-server.signal

		grpcServer.Stop()
		_ = lis.Close()
		server.signal <- 1
	}()

	return server, nil
}

func (s *server) Shutdown() {
	s.signal <- 1
	<-s.signal
	close(s.signal)
}

func TestDefaultDeadline(t *testing.T) {
	privateKey, err := parseRSAKey()
	require.NoError(t, err)

	caCertPEM, err := generateCA(privateKey)
	require.NoError(t, err)

	serverDataStore, err := generateDatastore(ca, privateKey, caCertPEM, 10*365*24*time.Hour)
	require.NoError(t, err)

	clientDataStore, err := generateDatastore(ca, privateKey, caCertPEM, 3*24*time.Hour)
	require.NoError(t, err)

	tlsCredentials, err := loadTLSCredentials(serverDataStore)
	require.NoError(t, err)

	server, err := launchServer(tlsCredentials)
	require.NoError(t, err)

	c := authorize.CreateClient()
	c.SetRequestTimeout(time.Millisecond)

	err = c.DialUsingCredentialsManager(context.Background(), &mockCredentialsFetcher{ds: clientDataStore}, "localhost", "10000", "")
	require.NoError(t, err)

	server.Shutdown()

	_, err = c.GetResourceWithContext(context.Background(), "", "")

	require.EqualError(t, err, "rpc error: code = DeadlineExceeded desc = context deadline exceeded")
}

func TestReconnect(t *testing.T) {
	privateKey, err := parseRSAKey()
	require.NoError(t, err)

	caCertPEM, err := generateCA(privateKey)
	require.NoError(t, err)

	serverDataStore, err := generateDatastore(ca, privateKey, caCertPEM, 10*365*24*time.Hour)
	require.NoError(t, err)

	clientDataStore, err := generateDatastore(ca, privateKey, caCertPEM, 24*time.Hour)
	require.NoError(t, err)

	tlsCredentials, err := loadTLSCredentials(serverDataStore)
	require.NoError(t, err)

	server, err := launchServer(tlsCredentials)
	require.NoError(t, err)

	c := authorize.CreateClient()

	err = c.DialUsingCredentialsManager(context.Background(), &mockCredentialsFetcher{ds: clientDataStore}, "localhost", "10000", "")
	require.NoError(t, err)

	server.Shutdown()

	server, err = launchServer(tlsCredentials)
	require.NoError(t, err)
	defer server.Shutdown()

	_, err = c.GetResourceWithContext(context.Background(), "", "")

	require.NoError(t, err)
}

func TestRetryPolicy(t *testing.T) {
	privateKey, err := parseRSAKey()
	require.NoError(t, err)

	caCertPEM, err := generateCA(privateKey)
	require.NoError(t, err)

	serverDataStore, err := generateDatastore(ca, privateKey, caCertPEM, 10*365*24*time.Hour)
	require.NoError(t, err)

	clientDataStore, err := generateDatastore(ca, privateKey, caCertPEM, 3*24*time.Hour)
	require.NoError(t, err)

	tlsCredentials, err := loadTLSCredentials(serverDataStore)
	require.NoError(t, err)

	server, err := launchServer(tlsCredentials)
	require.NoError(t, err)
	defer server.Shutdown()

	c := authorize.CreateClient()

	err = c.DialUsingCredentialsManager(context.Background(), &mockCredentialsFetcher{ds: clientDataStore}, "localhost", "10000", "")
	require.NoError(t, err)

	err = c.AddResourceWithContext(context.Background(), common.Origin{
		Id:       "",
		Type:     "",
		Provider: "",
	})
	require.NoError(t, err)
}

func TestClientHandshake_CertificateAboutToExpire(t *testing.T) {
	privateKey, err := parseRSAKey()
	require.NoError(t, err)

	caCertPEM, err := generateCA(privateKey)
	require.NoError(t, err)

	ds, err := generateDatastore(ca, privateKey, caCertPEM, 24*time.Hour-time.Second)
	require.NoError(t, err)

	cf := &mockCredentialsFetcher{ds: ds}

	ctx := context.Background()
	tls, err := authorize.NewAutoRefreshingTransportCredentials(ctx, cf, "secret", "localhost")
	require.NoError(t, err)

	require.Equal(t, 1, cf.callCount, "Certificates are loaded once during initialization")

	server, client := net.Pipe()
	server.Close()

	// Swap out certificates for fresh ones
	cf.ds, err = generateDatastore(ca, privateKey, caCertPEM, 3*24*time.Hour)
	require.NoError(t, err)

	for k := 0; k < 10; k++ {
		_, _, err = tls.ClientHandshake(ctx, "", client)
		require.Error(t, err, "io: read/write on closed pipe")
	}

	require.Equal(t, 2, cf.callCount, "Certificates are reloaded at first re-connect attempt only")
}
