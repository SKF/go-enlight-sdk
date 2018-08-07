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

// WithTransportCredentials returns a DialOption which configures
// a connection level security credentials (e.g., TLS/SSL).
func WithTransportCredentials(serverName, clientCert, clientKey, caCert string) (opt grpc.DialOption, err error) {
	certificate, err := tls.LoadX509KeyPair(clientCert, clientKey)

	if err != nil {
		err = fmt.Errorf("Failed to load client certs, %+v", err)
		return
	}

	certPool := x509.NewCertPool()
	bs, err := ioutil.ReadFile(caCert)
	if err != nil {
		err = fmt.Errorf("Failed to read ca cert, %+v", err)
		return
	}

	ok := certPool.AppendCertsFromPEM(bs)
	if !ok {
		err = errors.New("Failed to append certs")
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
