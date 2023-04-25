// Package grpc wraps "google.golang.org/grpc"
package grpc

import (
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	"io/ioutil"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

// WithBlock returns a DialOption which makes caller of Dial
// blocks until the underlying connection is up. Without this,
// Dial returns immediately and connecting the server happens
// in background.
func WithBlock() grpc.DialOption {
	return grpc.WithBlock()
}

// FailOnNonTempDialError returns a DialOption that specifies
// if gRPC fails on non-temporary dial errors. If f is true,
// and dialer returns a non-temporary error, gRPC will fail
// the connection to the network address and won't try to
// reconnect. The default value of FailOnNonTempDialError is
// false.
//
//This is an EXPERIMENTAL API.
func FailOnNonTempDialError(f bool) grpc.DialOption {
	return grpc.FailOnNonTempDialError(f)
}

// WithTransportCredentialsPEM returns a DialOption which configures
// a connection level security credentials (e.g., TLS/SSL).
func WithTransportCredentialsPEM(serverName string, certPEMBlock, keyPEMBlock, caPEMBlock []byte) (opt grpc.DialOption, err error) {
	certificate, err := tls.X509KeyPair(certPEMBlock, keyPEMBlock)
	if err != nil {
		err = fmt.Errorf("failed to load client certs, %+v", err)
		return
	}

	certPool := x509.NewCertPool()
	ok := certPool.AppendCertsFromPEM(caPEMBlock)
	if !ok {
		err = errors.New("failed to append certs")
		return
	}

	transportCreds := credentials.NewTLS(&tls.Config{
		ServerName:   serverName,
		Certificates: []tls.Certificate{certificate},
		RootCAs:      certPool,
	})

	opt = grpc.WithTransportCredentials(transportCreds)
	return
}

// WithTransportCredentials returns a DialOption which configures
// a connection level security credentials (e.g., TLS/SSL).
func WithTransportCredentials(serverName, certFile, keyFile, caFile string) (opt grpc.DialOption, err error) {
	certPEMBlock, err := ioutil.ReadFile(certFile)
	if err != nil {
		return nil, err
	}
	keyPEMBlock, err := ioutil.ReadFile(keyFile)
	if err != nil {
		return nil, err
	}
	caPEMBlock, err := ioutil.ReadFile(caFile)
	if err != nil {
		return nil, err
	}

	return WithTransportCredentialsPEM(serverName, certPEMBlock, keyPEMBlock, caPEMBlock)
}
